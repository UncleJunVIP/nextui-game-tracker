CREATE TABLE games
(
    id         INTEGER PRIMARY KEY,
    name       TEXT,
    path       TEXT,
    updated_at TEXT
);

CREATE UNIQUE INDEX games_id_index
    ON games (id);

CREATE INDEX games_name_index
    ON games (name);

CREATE INDEX games_path_index
    ON games (path);

CREATE TABLE play_sessions
(
    id           INTEGER PRIMARY KEY,
    game_id      INTEGER,
    start_time   TEXT,
    end_time     TEXT,
    force_closed INTEGER DEFAULT 0,
    invalid      INTEGER DEFAULT 0
);

CREATE UNIQUE INDEX play_session_id_index
    ON play_sessions (id);