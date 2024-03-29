// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: fixinformation.sql

package db

import (
	"context"

)

const createUserFixInformation = `-- name: CreateUserFixInformation :one
INSERT INTO fixinformation (
    first_name,
    last_name,
    email,
    hashed_password,
    birth,
    country,
    gender,
    blood,
    age,
    constellation,
    certification
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11
) RETURNING user_id, first_name, last_name, email, hashed_password, birth, country, gender, blood, age, constellation, certification, created_at, password_changed_at, role
`

type CreateUserFixInformationParams struct {
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	Email          string      `json:"email"`
	HashedPassword string `json:"hashed_password"`
	Birth          string      `json:"birth"`
	Country        string      `json:"country"`
	Gender         string      `json:"gender"`
	Blood          string      `json:"blood"`
	Age            int32       `json:"age"`
	Constellation  string      `json:"constellation"`
	Certification  bool `json:"certification"`
}

func (q *Queries) CreateUserFixInformation(ctx context.Context, arg CreateUserFixInformationParams) (Fixinformation, error) {
	row := q.db.QueryRow(ctx, createUserFixInformation,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.HashedPassword,
		arg.Birth,
		arg.Country,
		arg.Gender,
		arg.Blood,
		arg.Age,
		arg.Constellation,
		arg.Certification,
	)
	var i Fixinformation
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.HashedPassword,
		&i.Birth,
		&i.Country,
		&i.Gender,
		&i.Blood,
		&i.Age,
		&i.Constellation,
		&i.Certification,
		&i.CreatedAt,
		&i.PasswordChangedAt,
		&i.Role,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM fixinformation 
WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteUser, userID)
	return err
}

const fixSearchAccompany = `-- name: FixSearchAccompany :many
SELECT user_id, first_name, last_name, email, hashed_password, birth, country, gender, blood, age, constellation, certification, created_at, password_changed_at, role FROM fixinformation
WHERE user_id = $1
    AND (age >= $2 OR $2 IS NULL)
    AND (age < $3 OR $3 IS NULL)
`

type FixSearchAccompanyParams struct {
	UserID int32 `json:"user_id"`
	Age    *int32 `json:"age"`
	Age_2  *int32 `json:"age_2"`
}

func (q *Queries) FixSearchAccompany(ctx context.Context, arg FixSearchAccompanyParams) ([]Fixinformation, error) {
	rows, err := q.db.Query(ctx, fixSearchAccompany, arg.UserID, arg.Age, arg.Age_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Fixinformation{}
	for rows.Next() {
		var i Fixinformation
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.HashedPassword,
			&i.Birth,
			&i.Country,
			&i.Gender,
			&i.Blood,
			&i.Age,
			&i.Constellation,
			&i.Certification,
			&i.CreatedAt,
			&i.PasswordChangedAt,
			&i.Role,
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

const fixSearchHobby = `-- name: FixSearchHobby :many
SELECT user_id, first_name, last_name, email, hashed_password, birth, country, gender, blood, age, constellation, certification, created_at, password_changed_at, role FROM fixinformation
WHERE user_id != $1
    AND (age >= $2 OR $2 IS NULL)
    AND (age < $3 OR $3 IS NULL)
    AND (gender = $4 OR $4 IS NULL)
`

type FixSearchHobbyParams struct {
	UserID int32  `json:"user_id"`
	Age    *int32  `json:"age"`
	Age_2  *int32  `json:"age_2"`
	Gender *string `json:"gender"`
}

func (q *Queries) FixSearchHobby(ctx context.Context, arg FixSearchHobbyParams) ([]Fixinformation, error) {
	rows, err := q.db.Query(ctx, fixSearchHobby,
		arg.UserID,
		arg.Age,
		arg.Age_2,
		arg.Gender,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Fixinformation{}
	for rows.Next() {
		var i Fixinformation
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.HashedPassword,
			&i.Birth,
			&i.Country,
			&i.Gender,
			&i.Blood,
			&i.Age,
			&i.Constellation,
			&i.Certification,
			&i.CreatedAt,
			&i.PasswordChangedAt,
			&i.Role,
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

const fixSearchLover = `-- name: FixSearchLover :many
SELECT user_id, first_name, last_name, email, hashed_password, birth, country, gender, blood, age, constellation, certification, created_at, password_changed_at, role FROM fixinformation
WHERE user_id != $1  
    AND (age >= $2 OR $2 IS NULL)
    AND (age < $3 OR $3 IS NULL)
    AND (gender = $4 OR $4 IS NULL)
`

type FixSearchLoverParams struct {
	UserID int32  `json:"user_id"`
	Age    int32  `json:"age"`
	Age_2  int32  `json:"age_2"`
	Gender string `json:"gender"`
}

func (q *Queries) FixSearchLover(ctx context.Context, arg FixSearchLoverParams) ([]Fixinformation, error) {
	rows, err := q.db.Query(ctx, fixSearchLover,
		arg.UserID,
		arg.Age,
		arg.Age_2,
		arg.Gender,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Fixinformation{}
	for rows.Next() {
		var i Fixinformation
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.HashedPassword,
			&i.Birth,
			&i.Country,
			&i.Gender,
			&i.Blood,
			&i.Age,
			&i.Constellation,
			&i.Certification,
			&i.CreatedAt,
			&i.PasswordChangedAt,
			&i.Role,
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

const getUserFixInformation = `-- name: GetUserFixInformation :one
SELECT user_id, first_name, last_name, email, hashed_password, birth, country, gender, blood, age, constellation, certification, created_at, password_changed_at, role FROM fixinformation
WHERE user_id = $1
`

func (q *Queries) GetUserFixInformation(ctx context.Context, userID int32) (Fixinformation, error) {
	row := q.db.QueryRow(ctx, getUserFixInformation, userID)
	var i Fixinformation
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.HashedPassword,
		&i.Birth,
		&i.Country,
		&i.Gender,
		&i.Blood,
		&i.Age,
		&i.Constellation,
		&i.Certification,
		&i.CreatedAt,
		&i.PasswordChangedAt,
		&i.Role,
	)
	return i, err
}

const listFixInformation = `-- name: ListFixInformation :many
SELECT user_id, first_name, last_name, email, hashed_password, birth, country, gender, blood, age, constellation, certification, created_at, password_changed_at, role FROM fixinformation
ORDER BY user_id
`

func (q *Queries) ListFixInformation(ctx context.Context) ([]Fixinformation, error) {
	rows, err := q.db.Query(ctx, listFixInformation)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Fixinformation{}
	for rows.Next() {
		var i Fixinformation
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.HashedPassword,
			&i.Birth,
			&i.Country,
			&i.Gender,
			&i.Blood,
			&i.Age,
			&i.Constellation,
			&i.Certification,
			&i.CreatedAt,
			&i.PasswordChangedAt,
			&i.Role,
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

const loginAtEmail = `-- name: LoginAtEmail :one
SELECT user_id, first_name, last_name, email, hashed_password, birth, country, gender, blood, age, constellation, certification, created_at, password_changed_at, role FROM fixinformation
WHERE email = $1
`

func (q *Queries) LoginAtEmail(ctx context.Context, email string) (Fixinformation, error) {
	row := q.db.QueryRow(ctx, loginAtEmail, email)
	var i Fixinformation
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.HashedPassword,
		&i.Birth,
		&i.Country,
		&i.Gender,
		&i.Blood,
		&i.Age,
		&i.Constellation,
		&i.Certification,
		&i.CreatedAt,
		&i.PasswordChangedAt,
		&i.Role,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE fixinformation
SET hashed_password = $2
WHERE user_id = $1
RETURNING user_id, first_name, last_name, email, hashed_password, birth, country, gender, blood, age, constellation, certification, created_at, password_changed_at, role
`

type UpdatePasswordParams struct {
	UserID         int32       `json:"user_id"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (Fixinformation, error) {
	row := q.db.QueryRow(ctx, updatePassword, arg.UserID, arg.HashedPassword)
	var i Fixinformation
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.HashedPassword,
		&i.Birth,
		&i.Country,
		&i.Gender,
		&i.Blood,
		&i.Age,
		&i.Constellation,
		&i.Certification,
		&i.CreatedAt,
		&i.PasswordChangedAt,
		&i.Role,
	)
	return i, err
}
