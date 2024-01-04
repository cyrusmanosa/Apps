// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: record.sql

package db

import (
	"context"
	"fmt"
	"time"

)

func (q *Queries) CreateChatTable(ctx context.Context, tablename string) error {

	createTable := fmt.Sprintf(`CREATE TABLE %s (
		"target_id" INT NOT NULL,
		"msg_type" VARCHAR NOT NULL,
		"message" VARCHAR,
		"images" VARCHAR,
		"created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
	  );`,tablename)
	
		_, err := q.db.Exec(ctx, createTable)
		return err
	}



type DeleteRecordParams struct {
	TargetID  int32              `json:"target_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) DeleteRecord(ctx context.Context, arg DeleteRecordParams, tablename string) error {

deleteRecord := fmt.Sprintf(`-- name: DeleteRecord :exec
DELETE FROM %s
WHERE target_id = $1
AND created_at = $2
`,tablename)

	_, err := q.db.Exec(ctx, deleteRecord, arg.TargetID, arg.CreatedAt)
	return err
}



func (q *Queries) Getrecord(ctx context.Context, targetID int32,tablename string) ([]Record, error) {

	getrecord := fmt.Sprintf(`-- name: Getrecord :many
	SELECT target_id, msg_type, message, images, created_at FROM %s
	WHERE target_id = $1
	ORDER BY created_at
	`,tablename)


	rows, err := q.db.Query(ctx, getrecord, targetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Record{}
	for rows.Next() {
		var i Record
		if err := rows.Scan(
			&i.TargetID,
			&i.MsgType,
			&i.Message,
			&i.Images,
			&i.CreatedAt,
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

type CreateRecordParams struct {
	TargetID int32       `json:"target_id"`
	MsgType  string      `json:"msg_type"`
	Message  string `json:"message"`
	Images   string `json:"images"`
}

func (q *Queries) CreateRecord(ctx context.Context, arg CreateRecordParams,tablename string) (Record, error) {

	CreateRecord := fmt.Sprintf(`-- name: CreateRecord :one
	INSERT INTO %s (
		target_id,
		msg_type,
		message,
		images
	) VALUES (
		$1,$2,$3,$4
	) RETURNING target_id, msg_type, message, images, created_at
	`,tablename)

	row := q.db.QueryRow(ctx, CreateRecord,
		arg.TargetID,
		arg.MsgType,
		arg.Message,
		arg.Images,
	)
	var i Record
	err := row.Scan(
		&i.TargetID,
		&i.MsgType,
		&i.Message,
		&i.Images,
		&i.CreatedAt,
	)
	return i, err
}

type UpdateRecordParams struct {
	TargetID  int32              `json:"target_id"`
	MsgType   string             `json:"msg_type"`
	Message   string        `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) UpdateRecord(ctx context.Context, arg UpdateRecordParams,tablename string) (Record, error) {
	 updateRecord := fmt.Sprintf(`-- name: UpdateRecord :one
	UPDATE %s
	SET message = $4
	WHERE target_id = $1
	AND msg_type = $2
	AND created_at = $3
	RETURNING target_id, msg_type, message, images, created_at
	`,tablename)
	row := q.db.QueryRow(ctx, updateRecord,
		arg.TargetID,
		arg.MsgType,
		arg.CreatedAt,
		arg.Message,
	)
	var i Record
	err := row.Scan(
		&i.TargetID,
		&i.MsgType,
		&i.Message,
		&i.Images,
		&i.CreatedAt,
	)
	return i, err
}
