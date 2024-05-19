// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package storage

import (
	"context"
	"time"
)

const createProfile = `-- name: CreateProfile :exec
INSERT INTO profiles (username, password, type) 
VALUES ( ?, ?, ? )
`

type CreateProfileParams struct {
	Username string
	Password string
	Type     int64
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) error {
	_, err := q.db.ExecContext(ctx, createProfile, arg.Username, arg.Password, arg.Type)
	return err
}

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (id, user_id, created_at, expires_at, device, device_name, client_name, client_version)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id, user_id, created_at, expires_at, device, device_name, client_name, client_version
`

type CreateSessionParams struct {
	ID            string
	UserID        int64
	CreatedAt     time.Time
	ExpiresAt     time.Time
	Device        string
	DeviceName    string
	ClientName    string
	ClientVersion string
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.UserID,
		arg.CreatedAt,
		arg.ExpiresAt,
		arg.Device,
		arg.DeviceName,
		arg.ClientName,
		arg.ClientVersion,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.Device,
		&i.DeviceName,
		&i.ClientName,
		&i.ClientVersion,
	)
	return i, err
}

const getAdminUser = `-- name: GetAdminUser :one
SELECT id, username, password, type FROM profiles
WHERE type = 0
`

func (q *Queries) GetAdminUser(ctx context.Context) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getAdminUser)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Type,
	)
	return i, err
}

const getProfiles = `-- name: GetProfiles :many
SELECT id, username FROM profiles
`

type GetProfilesRow struct {
	ID       int64
	Username string
}

func (q *Queries) GetProfiles(ctx context.Context) ([]GetProfilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getProfiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProfilesRow
	for rows.Next() {
		var i GetProfilesRow
		if err := rows.Scan(&i.ID, &i.Username); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserWithPassword = `-- name: GetUserWithPassword :one
SELECT id, username, password, type FROM profiles 
WHERE username = ? and password = ?
`

type GetUserWithPasswordParams struct {
	Username string
	Password string
}

func (q *Queries) GetUserWithPassword(ctx context.Context, arg GetUserWithPasswordParams) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getUserWithPassword, arg.Username, arg.Password)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Type,
	)
	return i, err
}

const isFinishedSetup = `-- name: IsFinishedSetup :one
SELECT count(*) FROM profiles
`

func (q *Queries) IsFinishedSetup(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, isFinishedSetup)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateAdminUser = `-- name: UpdateAdminUser :exec
UPDATE profiles 
SET username = ?, password = ?
WHERE type = 0
`

type UpdateAdminUserParams struct {
	Username string
	Password string
}

func (q *Queries) UpdateAdminUser(ctx context.Context, arg UpdateAdminUserParams) error {
	_, err := q.db.ExecContext(ctx, updateAdminUser, arg.Username, arg.Password)
	return err
}
