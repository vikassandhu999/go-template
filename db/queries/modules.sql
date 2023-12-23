-- name: GetModule :one
SELECT * FROM modules
WHERE id = $1 LIMIT 1;
