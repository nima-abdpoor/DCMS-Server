package api

import (
	"DCMS/db/postgresql/sqlc"
)

func MapToUrlSecondDb(urlSeconds []UrlSecond, id int64) []db.UrlSecondTx {
	var result []db.UrlSecondTx
	for i, second := range urlSeconds {
		result = append(result, db.UrlSecondTx{
			UniqueID: id,
			UrlHash:  second.UrlHash,
		})
		for _, regex := range second.Regex {
			result[i].Regex = append(result[i].Regex, db.Regex{
				Regex:       regex.Regex,
				StartIndex:  regex.StartIndex,
				FinishIndex: regex.FinishIndex,
			})
		}
	}
	return result
}
