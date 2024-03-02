-- name: GetPages :one
SELECT * FROM pages
WHERE id = $1 LIMIT 1;
