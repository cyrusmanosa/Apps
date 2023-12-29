// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: payment.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payment (
    pay_id,
    fullname,
    payment_type,
    amount,
    product
) VALUES (
    $1,$2,$3,$4,$5
) RETURNING pay_id, fullname, payment_type, amount, product, pay_at
`

type CreatePaymentParams struct {
	PayID       uuid.UUID `json:"pay_id"`
	Fullname    string    `json:"fullname"`
	PaymentType string    `json:"payment_type"`
	Amount      int32     `json:"amount"`
	Product     string    `json:"product"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRow(ctx, createPayment,
		arg.PayID,
		arg.Fullname,
		arg.PaymentType,
		arg.Amount,
		arg.Product,
	)
	var i Payment
	err := row.Scan(
		&i.PayID,
		&i.Fullname,
		&i.PaymentType,
		&i.Amount,
		&i.Product,
		&i.PayAt,
	)
	return i, err
}

const getPayment = `-- name: GetPayment :one
SELECT pay_id, fullname, payment_type, amount, product, pay_at FROM payment
WHERE pay_id = $1
`

func (q *Queries) GetPayment(ctx context.Context, payID uuid.UUID) (Payment, error) {
	row := q.db.QueryRow(ctx, getPayment, payID)
	var i Payment
	err := row.Scan(
		&i.PayID,
		&i.Fullname,
		&i.PaymentType,
		&i.Amount,
		&i.Product,
		&i.PayAt,
	)
	return i, err
}
