// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
)

type Config struct {
	ID       string `json:"id"`
	IsLive   bool   `json:"is_live"`
	SyncType string `json:"sync_type"`
}

type Requesturl struct {
	ID         string         `json:"id"`
	RequestUrl sql.NullString `json:"request_url"`
}

type Urlfirst struct {
	ID      string         `json:"id"`
	UrlHash sql.NullString `json:"url_hash"`
}

type Urlsecond struct {
	ID          string         `json:"id"`
	UrlHash     sql.NullString `json:"url_hash"`
	Regex       sql.NullString `json:"regex"`
	StartIndex  sql.NullInt32  `json:"start_index"`
	FinishIndex sql.NullInt32  `json:"finish_index"`
}
