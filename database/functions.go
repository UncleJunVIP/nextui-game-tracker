package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	nextuigametracker "nextui-game-tracker"
	"nextui-game-tracker/utils"
	"os"
	"path/filepath"
	"sync"
)

var initOnce sync.Once

var db *sql.DB
var queries *Queries

func InitializeDB(isDev bool) {
	initOnce.Do(func() {
		initDatabase(isDev)
	})
}

func initDatabase(isDev bool) {
	ctx := context.Background()

	var err error
	dbPath := utils.GetGameTrackerDBPath(isDev)

	dbDir := filepath.Dir(dbPath)
	if dbDir != "." && dbDir != "" {
		err := os.MkdirAll(dbDir, 0755)
		if err != nil {
			log.Fatal("Unable to open database file", err)
		}
	}

	db, queries, err = openDatabase(isDev)
	if err != nil {
		log.Fatal("Unable to open database file", err)
	}

	schemaExists, err := tableExists(db, "games")
	if !schemaExists {
		log.Println("Schema does not exist, initializing...")
		if _, err := db.ExecContext(ctx, nextuigametracker.DDL); err != nil {
			log.Fatal("Unable to init schema", err)
		} else {
			log.Println("Schema initialized")
		}
	}

	queries = New(db)
}

func StartSession(gamePath string) error {
	_, err := stopExistingSessions(queries, true)
	if err != nil {
		return fmt.Errorf("unable to close existing sessions %w", err)
	}

	id, err := queries.FetchIDByPath(context.Background(), sql.NullString{String: gamePath, Valid: true})

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

	err = queries.StartSession(context.Background(), sessionParams)

	return err
}

func ResumeSession() (int64, error) {
	_, err := stopExistingSessions(queries, true)
	if err != nil {
		return -1, fmt.Errorf("unable to close existing sessions %w", err)
	}

	lastSession, err := queries.FindLastSession(context.Background())

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return -1, nil
	} else if err != nil {
		return -1, err
	}

	sessionParams := StartSessionParams{
		GameID:    sql.NullInt64{Int64: lastSession.GameID.Int64, Valid: true},
		StartTime: sql.NullString{String: utils.Now(), Valid: true},
	}

	err = queries.StartSession(context.Background(), sessionParams)
	if err != nil {
		return -1, err
	}

	return lastSession.GameID.Int64, nil

}

func StopSession() ([]PlaySession, error) {
	closedSessions, err := stopExistingSessions(queries, false)
	if err != nil {
		return nil, fmt.Errorf("unable to stop existing session(s) %w", err)
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

func stopExistingSessions(queries *Queries, forceClosed bool) ([]PlaySession, error) {
	csp := StopSessionParams{
		EndTime:      sql.NullString{String: utils.Now(), Valid: true},
		ForceStopped: sql.NullInt64{Int64: utils.BoolToInt64(forceClosed), Valid: true},
	}

	return queries.StopSession(context.Background(), csp)
}

func openDatabase(isDev bool) (*sql.DB, *Queries, error) {
	var err error
	dbPath := utils.GetGameTrackerDBPath(isDev)

	dbDir := filepath.Dir(dbPath)
	if dbDir != "." && dbDir != "" {
		err := os.MkdirAll(dbDir, 0755)
		if err != nil {
			return nil, nil, err
		}
	}

	dbc, err := sql.Open("sqlite", "file:"+dbPath)
	if err != nil {
		return nil, nil, err
	}

	_, err = dbc.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, nil, err
	}

	queries := New(dbc)

	return dbc, queries, nil
}

func CloseDatabase() {
	_ = db.Close()
}

func tableExists(db *sql.DB, tableName string) (bool, error) {
	query := `SELECT name FROM sqlite_master WHERE type='table' AND name=?`
	var name string
	err := db.QueryRow(query, tableName).Scan(&name)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return err == nil, err
}
