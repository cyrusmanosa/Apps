// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: accompanyrequest.sql

package db

import (
	"context"

)

const createAccompanyRequest = `-- name: CreateAccompanyRequest :one
INSERT INTO accompanyrequest (
    user_id,
    era,
    city,
    gender,
    speaklanguage,
    find_type,
    find_target,
    sociability,
    certification
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9
) RETURNING user_id, era, city, gender, speaklanguage, find_type, find_target, sociability, certification, info_changed_at
`

type CreateAccompanyRequestParams struct {
	UserID        int32       `json:"user_id"`
	Era           int32 `json:"era"`
	City          string      `json:"city"`
	Gender        string `json:"gender"`
	Speaklanguage string      `json:"speaklanguage"`
	FindType      string      `json:"find_type"`
	FindTarget    string      `json:"find_target"`
	Sociability   string      `json:"sociability"`
	Certification bool `json:"certification"`
}

func (q *Queries) CreateAccompanyRequest(ctx context.Context, arg CreateAccompanyRequestParams) (Accompanyrequest, error) {
	row := q.db.QueryRow(ctx, createAccompanyRequest,
		arg.UserID,
		arg.Era,
		arg.City,
		arg.Gender,
		arg.Speaklanguage,
		arg.FindType,
		arg.FindTarget,
		arg.Sociability,
		arg.Certification,
	)
	var i Accompanyrequest
	err := row.Scan(
		&i.UserID,
		&i.Era,
		&i.City,
		&i.Gender,
		&i.Speaklanguage,
		&i.FindType,
		&i.FindTarget,
		&i.Sociability,
		&i.Certification,
		&i.InfoChangedAt,
	)
	return i, err
}

const deleteUserAccompany = `-- name: DeleteUserAccompany :exec
DELETE FROM accompanyrequest 
WHERE user_id = $1
`

func (q *Queries) DeleteUserAccompany(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteUserAccompany, userID)
	return err
}

const getUserAccompany = `-- name: GetUserAccompany :one
SELECT user_id, era, city, gender, speaklanguage, find_type, find_target, sociability, certification, info_changed_at FROM accompanyrequest
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUserAccompany(ctx context.Context, userID int32) (Accompanyrequest, error) {
	row := q.db.QueryRow(ctx, getUserAccompany, userID)
	var i Accompanyrequest
	err := row.Scan(
		&i.UserID,
		&i.Era,
		&i.City,
		&i.Gender,
		&i.Speaklanguage,
		&i.FindType,
		&i.FindTarget,
		&i.Sociability,
		&i.Certification,
		&i.InfoChangedAt,
	)
	return i, err
}

const listUserAccompany = `-- name: ListUserAccompany :many
SELECT user_id, era, city, gender, speaklanguage, find_type, find_target, sociability, certification, info_changed_at FROM accompanyrequest
ORDER BY user_id
`

func (q *Queries) ListUserAccompany(ctx context.Context) ([]Accompanyrequest, error) {
	rows, err := q.db.Query(ctx, listUserAccompany)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Accompanyrequest{}
	for rows.Next() {
		var i Accompanyrequest
		if err := rows.Scan(
			&i.UserID,
			&i.Era,
			&i.City,
			&i.Gender,
			&i.Speaklanguage,
			&i.FindType,
			&i.FindTarget,
			&i.Sociability,
			&i.Certification,
			&i.InfoChangedAt,
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

const updateUserAccompany = `-- name: UpdateUserAccompany :one
UPDATE accompanyrequest
SET era = $2,
    city = $3,
    gender = $4,
    speaklanguage = $5,
    find_type = $6,
    find_target = $7,
    sociability = $8,
    certification = $9
WHERE user_id = $1
RETURNING user_id, era, city, gender, speaklanguage, find_type, find_target, sociability, certification, info_changed_at
`

type UpdateUserAccompanyParams struct {
	UserID        int32       `json:"user_id"`
	Era           int32 `json:"era"`
	City          string      `json:"city"`
	Gender        string `json:"gender"`
	Speaklanguage string      `json:"speaklanguage"`
	FindType      string      `json:"find_type"`
	FindTarget    string      `json:"find_target"`
	Sociability   string      `json:"sociability"`
	Certification bool `json:"certification"`
}

func (q *Queries) UpdateUserAccompany(ctx context.Context, arg UpdateUserAccompanyParams) (Accompanyrequest, error) {
	row := q.db.QueryRow(ctx, updateUserAccompany,
		arg.UserID,
		arg.Era,
		arg.City,
		arg.Gender,
		arg.Speaklanguage,
		arg.FindType,
		arg.FindTarget,
		arg.Sociability,
		arg.Certification,
	)
	var i Accompanyrequest
	err := row.Scan(
		&i.UserID,
		&i.Era,
		&i.City,
		&i.Gender,
		&i.Speaklanguage,
		&i.FindType,
		&i.FindTarget,
		&i.Sociability,
		&i.Certification,
		&i.InfoChangedAt,
	)
	return i, err
}
