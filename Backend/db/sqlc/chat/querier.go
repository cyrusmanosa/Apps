// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type ChatQuerier interface {
	CreateRecord(ctx context.Context, arg CreateRecordParams) (Record, error)
	DeleteRecord(ctx context.Context, arg DeleteRecordParams) error
	GetRecord(ctx context.Context, targetID int32) ([]Record, error)
	GetTargetID(ctx context.Context) ([]int32, error)
	UpdateRecord(ctx context.Context, arg UpdateRecordParams) (Record, error)
}

var _ ChatQuerier = (*Queries)(nil)
