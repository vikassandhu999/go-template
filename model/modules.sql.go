// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: modules.sql

package model

import (
	"context"
)

const getPages = `-- name: GetPages :one
SELECT id, kind, title, slug, working_content_id, published_content_id FROM pages
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPages(ctx context.Context, id int64) (Page, error) {
	row := q.db.QueryRow(ctx, getPages, id)
	var i Page
	err := row.Scan(
		&i.ID,
		&i.Kind,
		&i.Title,
		&i.Slug,
		&i.WorkingContentID,
		&i.PublishedContentID,
	)
	return i, err
}
