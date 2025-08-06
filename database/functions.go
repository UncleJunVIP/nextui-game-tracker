package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	"nextui-game-tracker/utils"
	"os"
	"path/filepath"
)

var devMode bool

func EnableDevMode() {
	devMode = true
}

func StartSession(gamePath string) error {
	db, queries := openDatabase()
	defer closeDatabase(db)

	_, err := closeExistingSessions(queries, true)
	if err != nil {
		return fmt.Errorf("unable to close existing sessions %w", err)
	}

	ctx := context.Background()

	id, err := queries.FetchIDByPath(ctx, sql.NullString{String: gamePath, Valid: true})

	if id == 0 {
		id, err = NewGame(gamePath, queries)
		if err != nil {
			return fmt.Errorf("unable to create new game %w", err)
		}
	}

	sessionParams := StartSessionParams{
		GameID:    sql.NullInt64{Int64: id, Valid: true},
		StartTime: sql.NullString{String: utils.Now(), Valid: true},
	}

	err = queries.StartSession(ctx, sessionParams)

	return err
}

func EndSession() ([]PlaySession, error) {
	db, queries := openDatabase()
	defer closeDatabase(db)

	closedSessions, err := closeExistingSessions(queries, false)
	if err != nil {
		return nil, fmt.Errorf("unable to close existing sessions %w", err)
	}

	return closedSessions, nil
}

func NewGame(path string, queries *Queries) (int64, error) {
	name := utils.ParseRomName(path)

	gameParams := NewGameParams{
		Name:      sql.NullString{String: name, Valid: true},
		Path:      sql.NullString{String: path, Valid: true},
		UpdatedAt: sql.NullString{String: utils.Now(), Valid: true},
	}

	return queries.NewGame(context.Background(), gameParams)
}

func closeExistingSessions(queries *Queries, forceClosed bool) ([]PlaySession, error) {
	csp := CloseSessionParams{
		EndTime:     sql.NullString{String: utils.Now(), Valid: true},
		ForceClosed: sql.NullInt64{Int64: utils.BoolToInt64(forceClosed), Valid: true},
	}

	return queries.CloseSession(context.Background(), csp)
}

func openDatabase() (*sql.DB, *Queries) {
	var err error
	dbPath := utils.GetGameTrackerDBPath(devMode)

	dbDir := filepath.Dir(dbPath)
	if dbDir != "." && dbDir != "" {
		err := os.MkdirAll(dbDir, 0755)
		if err != nil {
			log.Fatal("Unable to open database file", err)
		}
	}

	dbc, err := sql.Open("sqlite", "file:"+dbPath)
	if err != nil {
		log.Fatal("Unable to open database file", err)
	}

	queries := New(dbc)

	return dbc, queries
}

func closeDatabase(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
