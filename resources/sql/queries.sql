-- name: ListGames :many
SELECT *
FROM games
ORDER BY name;

-- name: FetchIDByPath :one
SELECT id
FROM games
WHERE path = ?;

-- name: NewGame :one
INSERT INTO games (name, path, updated_at)
VALUES (?, ?, ?)
RETURNING id;

-- name: StartSession :exec
INSERT INTO play_sessions (game_id, start_time)
VALUES (?, ?);

-- name: CloseSession :many
UPDATE play_sessions
SET end_time     = @end_time,
    force_closed = @force_closed,
    invalid      = CASE
                       WHEN @end_time < start_time THEN 1
                       ELSE 0
        END
WHERE end_time IS NULL
RETURNING *;
