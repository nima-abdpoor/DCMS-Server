package api

import (
	"DCMS/db/sqlc"
)

func MapToUrlSecondDb(urlSeconds []UrlSecond, id int64) []db.Urlsecond {
	var result []db.Urlsecond
	for _, second := range urlSeconds {
		result = append(result, db.Urlsecond{
			UniqueID:    id,
			UrlHash:     second.UrlHash,
			Regex:       second.Regex,
			StartIndex:  second.StartIndex,
			FinishIndex: second.FinishIndex,
		})
	}
	return result
}
