-- name: AccountCreated :one
INSERT INTO accounts (
    domain, name, email, password
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1;

-- name: GetAccountByDomain :one
SELECT * FROM accounts
WHERE domain = $1;

-- name: UpdateAccount :one
UPDATE accounts
SET
  name = $2,
  email = $3,
  password = $4
WHERE id = $1
RETURNING *;
