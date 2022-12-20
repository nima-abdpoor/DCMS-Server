package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	q  *Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		q:  New(db),
		db: db,
	}
}

// execute a function within a database transaction.
func (store *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
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
	ID           int64
	SyncType     string
	IsLive       bool
	SaveRequest  bool
	SaveResponse bool
	SaveError    bool
	SaveSuccess  bool
	UrlHashFirst []string
	UrlSecond    []UrlSecondTx
	RequestUrl   []string
}

type AddConfigTxResult struct {
	Config     Config
	UrlFirst   []Urlfirst
	UrlSecond  []UrlSecondTx
	RequestUrl []Requesturl
}

func (store Store) AddConfigTx(ctx context.Context, arg AddConfigTxParams) (AddConfigTxResult, error) {
	var result AddConfigTxResult
	result.UrlFirst = make([]Urlfirst, len(arg.UrlHashFirst))
	result.UrlSecond = make([]UrlSecondTx, len(arg.UrlSecond))
	result.RequestUrl = make([]Requesturl, len(arg.RequestUrl))
	err := store.execTx(ctx, func(queries *Queries) error {
		var err error
		result.Config, err = queries.CreateConfig(ctx, CreateConfigParams{
			ID:           arg.ID,
			SyncType:     arg.SyncType,
			IsLive:       arg.IsLive,
			SaveRequest:  arg.SaveRequest,
			SaveResponse: arg.SaveResponse,
			SaveError:    arg.SaveError,
			SaveSuccess:  arg.SaveSuccess,
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
			result.UrlSecond[i].Regex = make([]Regex, len(arg.UrlSecond[i].Regex))
			urlSeconds, err := queries.CreateUrlSecond(ctx, CreateUrlSecondParams{
				UniqueID: arg.ID,
				UrlHash:  urlSecond.UrlHash,
			})
			if err != nil {
				return err
			}
			for j, regex := range arg.UrlSecond[i].Regex {
				result.UrlSecond[i].Regex[j], err = queries.CreateRegex(ctx, CreateRegexParams{
					UrlID:       urlSeconds.ID,
					Regex:       regex.Regex,
					StartIndex:  regex.StartIndex,
					FinishIndex: regex.FinishIndex,
				})
				if err != nil {
					return err
				}
			}

			result.UrlSecond[i] = UrlSecondTx{
				ID:       urlSeconds.ID,
				UniqueID: urlSeconds.UniqueID,
				UrlHash:  urlSeconds.UrlHash,
				Regex:    arg.UrlSecond[i].Regex,
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

type GetConfigTxParams struct {
	ID int64 `json:"id"`
}

type GetConfigTxResult struct {
	Config     Config        `json:"config"`
	UrlFirst   []Urlfirst    `json:"urlFirst"`
	UrlSecond  []UrlSecondTx `json:"urlSecond"`
	RequestUrl []Requesturl  `json:"requestUrl"`
}

type UrlSecondTx struct {
	ID       int64
	UniqueID int64
	UrlHash  string
	Regex    []Regex
}

func (store Store) GetConfigTx(ctx context.Context, arg GetConfigTxParams) (GetConfigTxResult, error) {
	var result GetConfigTxResult
	err := store.execTx(ctx, func(queries *Queries) error {
		var err error
		result.Config, err = queries.GetConfig(ctx, arg.ID)
		if err != nil {
			return err
		}

		result.UrlFirst, err = queries.GetUrlFirstByUniqueId(ctx, arg.ID)
		if err != nil {
			return err
		}

		urlSeconds, err := queries.GetUrlSecondByUniqueId(ctx, arg.ID)
		if err != nil {
			return err
		}

		result.UrlSecond = make([]UrlSecondTx, len(urlSeconds))
		for i, url := range urlSeconds {
			regexes, err := queries.GetRegexByUrlId(ctx, url.ID)
			if err != nil {
				return err
			}
			result.UrlSecond[i] = UrlSecondTx{
				ID:       url.ID,
				UniqueID: url.UniqueID,
				UrlHash:  url.UrlHash,
				Regex:    regexes,
			}
		}

		result.RequestUrl, err = queries.GetRequestUrlByUniqueId(context.Background(), arg.ID)
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
