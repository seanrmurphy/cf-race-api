-- name: GetRaceInfo :many
SELECT * FROM race_info
LIMIT 10; 

-- name: GetRaceInfoById :one
SELECT * FROM race_info
WHERE id = ? LIMIT 1;

-- name: GetRaceResults :many
SELECT * FROM race_results
WHERE race_id = ? LIMIT 100;

-- name: AddRace :one
INSERT INTO race_info (
  location, name, event_date, run_types, created_at
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: AddRaceResult :one
INSERT INTO race_results (
  first_name, last_name, race_number, run_type, race_id, start_time, end_time
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetFilteredRaceResults :many
SELECT * FROM race_results
LIMIT 100;
