// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: socialmedia.sql

package db

import (
	"context"

)

const createsocialmedia = `-- name: Createsocialmedia :one
INSERT INTO socialmedia (
    user_id,
    target_id,
    image,
    contact,
    location,
    appointment,
    sns
) VALUES (
    $1,$2,$3,$4,$5,$6,$7
) RETURNING user_id, target_id, image, contact, location, appointment, sns
`

type CreatesocialmediaParams struct {
	UserID      int32       `json:"user_id"`
	TargetID    int32       `json:"target_id"`
	Image       bool `json:"image"`
	Contact     bool `json:"contact"`
	Location    bool `json:"location"`
	Appointment bool `json:"appointment"`
	Sns         bool `json:"sns"`
}

func (q *Queries) Createsocialmedia(ctx context.Context, arg CreatesocialmediaParams) (Socialmedium, error) {
	row := q.db.QueryRow(ctx, createsocialmedia,
		arg.UserID,
		arg.TargetID,
		arg.Image,
		arg.Contact,
		arg.Location,
		arg.Appointment,
		arg.Sns,
	)
	var i Socialmedium
	err := row.Scan(
		&i.UserID,
		&i.TargetID,
		&i.Image,
		&i.Contact,
		&i.Location,
		&i.Appointment,
		&i.Sns,
	)
	return i, err
}

const deletesocialmedia = `-- name: Deletesocialmedia :exec
DELETE FROM socialmedia 
WHERE user_id = $1
`

func (q *Queries) Deletesocialmedia(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deletesocialmedia, userID)
	return err
}


const getsocialmedia = `-- name: Getsocialmedia :one
SELECT user_id, target_id, image, contact, location, appointment, sns FROM socialmedia
WHERE user_id = $1 AND target_id = $2 
LIMIT 1
`

type GetsocialmediaParams struct {
	UserID   int32 `json:"user_id"`
	TargetID int32 `json:"target_id"`
}

func (q *Queries) Getsocialmedia(ctx context.Context, arg GetsocialmediaParams) (Socialmedium, error) {
	row := q.db.QueryRow(ctx, getsocialmedia, arg.UserID, arg.TargetID)
	var i Socialmedium
	err := row.Scan(
		&i.UserID,
		&i.TargetID,
		&i.Image,
		&i.Contact,
		&i.Location,
		&i.Appointment,
		&i.Sns,
	)
	return i, err
}


const listsocialmedia = `-- name: Listsocialmedia :many
SELECT user_id, target_id, image, contact, location, appointment, sns FROM socialmedia
ORDER BY user_id
`

func (q *Queries) Listsocialmedia(ctx context.Context) ([]Socialmedium, error) {
	rows, err := q.db.Query(ctx, listsocialmedia)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Socialmedium{}
	for rows.Next() {
		var i Socialmedium
		if err := rows.Scan(
			&i.UserID,
			&i.TargetID,
			&i.Image,
			&i.Contact,
			&i.Location,
			&i.Appointment,
			&i.Sns,
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

const updatesocialmedia = `-- name: Updatesocialmedia :one
UPDATE socialmedia
SET target_id = $2,
    image = $3,
    contact = $4,
    location = $5,
    appointment = $6,
    sns  = $7
WHERE user_id = $1
RETURNING user_id, target_id, image, contact, location, appointment, sns
`

type UpdatesocialmediaParams struct {
	UserID      int32       `json:"user_id"`
	TargetID    int32       `json:"target_id"`
	Image       bool `json:"image"`
	Contact     bool `json:"contact"`
	Location    bool `json:"location"`
	Appointment bool `json:"appointment"`
	Sns         bool `json:"sns"`
}

func (q *Queries) Updatesocialmedia(ctx context.Context, arg UpdatesocialmediaParams) (Socialmedium, error) {
	row := q.db.QueryRow(ctx, updatesocialmedia,
		arg.UserID,
		arg.TargetID,
		arg.Image,
		arg.Contact,
		arg.Location,
		arg.Appointment,
		arg.Sns,
	)
	var i Socialmedium
	err := row.Scan(
		&i.UserID,
		&i.TargetID,
		&i.Image,
		&i.Contact,
		&i.Location,
		&i.Appointment,
		&i.Sns,
	)
	return i, err
}
