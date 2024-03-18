// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: subscriptions.sql

package model

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSubscription = `-- name: CreateSubscription :one
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
) RETURNING id, account_id, plan_id, start_date, end_date, trial_start_date, trial_end_date, status, cost
`

type CreateSubscriptionParams struct {
	AccountID      pgtype.Int4
	PlanID         pgtype.Int4
	StartDate      pgtype.Timestamptz
	EndDate        pgtype.Timestamptz
	TrialStartDate pgtype.Timestamptz
	TrialEndDate   pgtype.Timestamptz
	Status         pgtype.Int2
	Cost           pgtype.Numeric
}

func (q *Queries) CreateSubscription(ctx context.Context, arg CreateSubscriptionParams) (Subscription, error) {
	row := q.db.QueryRow(ctx, createSubscription,
		arg.AccountID,
		arg.PlanID,
		arg.StartDate,
		arg.EndDate,
		arg.TrialStartDate,
		arg.TrialEndDate,
		arg.Status,
		arg.Cost,
	)
	var i Subscription
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PlanID,
		&i.StartDate,
		&i.EndDate,
		&i.TrialStartDate,
		&i.TrialEndDate,
		&i.Status,
		&i.Cost,
	)
	return i, err
}

const deleteSubscription = `-- name: DeleteSubscription :one
DELETE FROM subscriptions
WHERE id = $1
RETURNING id, account_id, plan_id, start_date, end_date, trial_start_date, trial_end_date, status, cost
`

func (q *Queries) DeleteSubscription(ctx context.Context, id int32) (Subscription, error) {
	row := q.db.QueryRow(ctx, deleteSubscription, id)
	var i Subscription
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PlanID,
		&i.StartDate,
		&i.EndDate,
		&i.TrialStartDate,
		&i.TrialEndDate,
		&i.Status,
		&i.Cost,
	)
	return i, err
}

const getSubscriptionById = `-- name: GetSubscriptionById :one
SELECT id, account_id, plan_id, start_date, end_date, trial_start_date, trial_end_date, status, cost FROM subscriptions
WHERE id = $1
`

func (q *Queries) GetSubscriptionById(ctx context.Context, id int32) (Subscription, error) {
	row := q.db.QueryRow(ctx, getSubscriptionById, id)
	var i Subscription
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PlanID,
		&i.StartDate,
		&i.EndDate,
		&i.TrialStartDate,
		&i.TrialEndDate,
		&i.Status,
		&i.Cost,
	)
	return i, err
}

const listSubscriptionsForAccount = `-- name: ListSubscriptionsForAccount :many
SELECT id, account_id, plan_id, start_date, end_date, trial_start_date, trial_end_date, status, cost FROM subscriptions
WHERE account_id = $1
`

func (q *Queries) ListSubscriptionsForAccount(ctx context.Context, accountID pgtype.Int4) ([]Subscription, error) {
	rows, err := q.db.Query(ctx, listSubscriptionsForAccount, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Subscription
	for rows.Next() {
		var i Subscription
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.PlanID,
			&i.StartDate,
			&i.EndDate,
			&i.TrialStartDate,
			&i.TrialEndDate,
			&i.Status,
			&i.Cost,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSubscriptionStatus = `-- name: UpdateSubscriptionStatus :one
UPDATE subscriptions
SET status = $2
WHERE id = $1
RETURNING id, account_id, plan_id, start_date, end_date, trial_start_date, trial_end_date, status, cost
`

type UpdateSubscriptionStatusParams struct {
	ID     int32
	Status pgtype.Int2
}

func (q *Queries) UpdateSubscriptionStatus(ctx context.Context, arg UpdateSubscriptionStatusParams) (Subscription, error) {
	row := q.db.QueryRow(ctx, updateSubscriptionStatus, arg.ID, arg.Status)
	var i Subscription
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PlanID,
		&i.StartDate,
		&i.EndDate,
		&i.TrialStartDate,
		&i.TrialEndDate,
		&i.Status,
		&i.Cost,
	)
	return i, err
}