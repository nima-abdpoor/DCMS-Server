package db

import (
	"context"
	"database/sql"
	"fmt"
)

type store struct {
	q  *Queries
	db *sql.DB
}

func newStore(db *sql.DB) *store {
	return &store{
		q:  New(db),
		db: db,
	}
}

// execute a function within a database transaction.
func (store *store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type AddConfigTxParams struct {
	ID           int64       `json:"id"`
	SyncType     string      `json:"sync_type"`
	IsLive       bool        `json:"is_live"`
	UrlHashFirst []string    `json:"urlHashFirst"`
	UrlSecond    []Urlsecond `json:"urlSecond"`
	RequestUrl   []string    `json:"requestUrl"`
}

type AddConfigTxResult struct {
	Config     Config       `json:"config"`
	UrlFirst   []Urlfirst   `json:"urlFirst"`
	UrlSecond  []Urlsecond  `json:"urlSecond"`
	RequestUrl []Requesturl `json:"requestUrl"`
}

func (store store) AddConfigTx(ctx context.Context, arg AddConfigTxParams) (AddConfigTxResult, error) {
	var result AddConfigTxResult
	result.UrlFirst = make([]Urlfirst, len(arg.UrlHashFirst))
	result.UrlSecond = make([]Urlsecond, len(arg.UrlSecond))
	result.RequestUrl = make([]Requesturl, len(arg.RequestUrl))
	err := store.execTx(ctx, func(queries *Queries) error {
		var err error
		result.Config, err = queries.CreateConfig(ctx, CreateConfigParams{
			ID:       arg.ID,
			SyncType: arg.SyncType,
			IsLive:   arg.IsLive,
		})
		if err != nil {
			return err
		}

		for i, s := range arg.UrlHashFirst {
			result.UrlFirst[i], err = queries.CreateUrlFirst(ctx, CreateUrlFirstParams{
				UniqueID: arg.ID,
				UrlHash:  s,
			})
			if err != nil {
				return err
			}
		}

		for i, urlSecond := range arg.UrlSecond {
			result.UrlSecond[i], err = queries.CreateUrlSecond(ctx, CreateUrlSecondParams{
				UniqueID:    arg.ID,
				UrlHash:     urlSecond.UrlHash,
				Regex:       urlSecond.Regex,
				StartIndex:  urlSecond.StartIndex,
				FinishIndex: urlSecond.FinishIndex,
			})
			if err != nil {
				return err
			}
		}

		for i, s := range arg.RequestUrl {
			result.RequestUrl[i], err = queries.CreateRequestUrl(context.Background(), CreateRequestUrlParams{
				UniqueID:   arg.ID,
				RequestUrl: s,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	return result, err
}
