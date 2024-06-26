// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package storage

import (
	"database/sql"
	"time"
)

type ContentLibrary struct {
	ID             int64
	CreatedAt      time.Time
	FilePath       string
	MediaLibraryID int64
	Extension      string
	Name           string
}

type ContentMetadatum struct {
	ID            int64
	CreatedAt     time.Time
	ContentID     int64
	Title         string
	Description   string
	PosterUrl     string
	ReleaseDate   time.Time
	SeasonNumber  sql.NullInt64
	EpisodeNumber sql.NullInt64
	Type          string
}

type MediaLibrary struct {
	ID          int64
	CreatedAt   time.Time
	Name        string
	Description string
	DevicePath  string
	MediaType   string
	OwnerID     int64
}

type Profile struct {
	ID       int64
	Username string
	Password string
	Type     int64
}

type Session struct {
	ID            string
	UserID        int64
	CreatedAt     time.Time
	ExpiresAt     time.Time
	Device        string
	DeviceName    string
	ClientName    string
	ClientVersion string
}
