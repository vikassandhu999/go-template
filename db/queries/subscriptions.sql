-- name: CreateSubscription :one
INSERT INTO subscriptions (
    account_id,
    plan_id,
    start_date,
    end_date,
    trial_start_date,
    trial_end_date,
    status,
    cost
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING *;

-- name: GetSubscriptionById :one
SELECT * FROM subscriptions
WHERE id = $1;

-- name: ListSubscriptionsForAccount :many
SELECT * FROM subscriptions
WHERE account_id = $1;

-- name: UpdateSubscriptionStatus :one
UPDATE subscriptions
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSubscription :one
DELETE FROM subscriptions
WHERE id = $1
RETURNING *;
