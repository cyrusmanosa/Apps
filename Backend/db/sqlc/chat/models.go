// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Record struct {
	TargetID  int32              `json:"target_id"`
	RoleType  string             `json:"role_type"`
	MediaType string             `json:"media_type"`
	Message   pgtype.Text        `json:"message"`
	Media     pgtype.Text        `json:"media"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}
