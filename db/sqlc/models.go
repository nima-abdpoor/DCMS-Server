// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import ()

type Config struct {
	ID       int64  `json:"id"`
	IsLive   bool   `json:"is_live"`
	SyncType string `json:"sync_type"`
}

type Regex struct {
	ID          int64  `json:"id"`
	UrlID       int64  `json:"url_id"`
	Regex       string `json:"regex"`
	StartIndex  int32  `json:"start_index"`
	FinishIndex int32  `json:"finish_index"`
}

type Requesturl struct {
	ID         int64  `json:"id"`
	UniqueID   int64  `json:"unique_id"`
	RequestUrl string `json:"request_url"`
}

type Urlfirst struct {
	ID       int64  `json:"id"`
	UniqueID int64  `json:"unique_id"`
	UrlHash  string `json:"url_hash"`
}

type Urlsecond struct {
	ID       int64  `json:"id"`
	UniqueID int64  `json:"unique_id"`
	UrlHash  string `json:"url_hash"`
}
