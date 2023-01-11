// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

type Config struct {
	ID                     int64  `json:"id"`
	IsLive                 bool   `json:"is_live"`
	SaveRequest            bool   `json:"save_request"`
	SaveResponse           bool   `json:"save_response"`
	SaveError              bool   `json:"save_error"`
	SaveSuccess            bool   `json:"save_success"`
	NetworkType            string `json:"network_type"`
	RepeatInterval         int64  `json:"repeat_interval"`
	RepeatIntervalTimeUnit string `json:"repeat_interval_time_unit"`
	RequiresBatteryNotLow  bool   `json:"requires_battery_not_low"`
	RequiresStorageNotLow  bool   `json:"requires_storage_not_low"`
	RequiresCharging       bool   `json:"requires_charging"`
	RequiresDeviceIdl      bool   `json:"requires_device_idl"`
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