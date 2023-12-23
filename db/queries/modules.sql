-- name: GetModule :one
SELECT * FROM modules
WHERE id = $1 LIMIT 1;

-- name: CreateModule :one
INSERT INTO modules (domain, category, purpose) VALUES($1,$2,$3) RETURNING *;

-- name: ListModules :many
SELECT * FROM modules LIMIT $1 OFFSET $2;
