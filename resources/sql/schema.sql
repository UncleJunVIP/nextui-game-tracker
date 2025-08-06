CREATE TABLE games
(
    id         INTEGER PRIMARY KEY,
    name       TEXT,
    system_tag TEXT,
    path       TEXT,
    image_path TEXT,
    created_at TEXT,
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
    id         INTEGER PRIMARY KEY,
    game_id    INTEGER,
    start_time TEXT,
    end_time   TEXT,
    playtime   INTEGER
);

CREATE UNIQUE INDEX play_session_id_index
    ON play_sessions (id);