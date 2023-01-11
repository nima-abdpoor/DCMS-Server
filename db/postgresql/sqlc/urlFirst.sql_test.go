package db

import (
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUrlFirst(t *testing.T, config Config) Urlfirst {
	arg := CreateUrlFirstParams{
		UniqueID: config.ID,
		UrlHash:  util.RandomUrlHashGenerator(1)[0],
	}
	urlFirst, err := testQueries.CreateUrlFirst(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, urlFirst)
	require.Equal(t, arg.UniqueID, config.ID)
	require.Equal(t, arg.UrlHash, urlFirst.UrlHash)
	require.NotZero(t, urlFirst.ID)
	return urlFirst
}

func TestQueries_CreateUrlFirst(t *testing.T) {
	createRandomUrlFirst(t, createRandomConfig(t))
}

func TestQueries_DeleteUrlFirst(t *testing.T) {
	config := createRandomConfig(t)
	urlFirst := createRandomUrlFirst(t, config)
	err := testQueries.DeleteUrlFirst(context.Background(), urlFirst.ID)
	require.NoError(t, err)
	urlFirstSecond, err2 := testQueries.GetUrlFirst(context.Background(), urlFirst.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, urlFirstSecond)
}

func TestQueries_GetUrlFirst(t *testing.T) {
	config := createRandomConfig(t)
	urlFirst := createRandomUrlFirst(t, config)
	actualUrlFirst, err := testQueries.GetUrlFirst(context.Background(), urlFirst.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actualUrlFirst)
	require.Equal(t, urlFirst.ID, actualUrlFirst.ID)
	require.Equal(t, urlFirst.UrlHash, actualUrlFirst.UrlHash)
}

func TestQueries_ListUrlFirsts(t *testing.T) {
	for i := 0; i < 10; i++ {
		config := createRandomConfig(t)
		createRandomUrlFirst(t, config)
	}
	arg := ListUrlFirstsParams{
		Limit:  5,
		Offset: 5,
	}

	urlFirsts, err := testQueries.ListUrlFirsts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, urlFirsts, 5)
	for _, url := range urlFirsts {
		require.NotEmpty(t, url)
	}
}

func TestQueries_UpdateUrlFirst(t *testing.T) {
	config := createRandomConfig(t)
	urlFirst := createRandomUrlFirst(t, config)
	arg := UpdateUrlFirstParams{
		UniqueID: urlFirst.UniqueID,
		UrlHash:  util.RandomUrlHashGenerator(1)[0],
	}
	updatedUrlFirst, err := testQueries.UpdateUrlFirst(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUrlFirst)
	require.Equal(t, urlFirst.ID, updatedUrlFirst.ID)
	require.Equal(t, arg.UniqueID, updatedUrlFirst.UniqueID)
	require.Equal(t, arg.UrlHash, updatedUrlFirst.UrlHash)
}

func TestQueries_GetUrlFirstByUniqueId(t *testing.T) {
	config := createRandomConfig(t)
	var urlFirsts []Urlfirst
	numberOfUrlFirst := 10
	for i := 0; i < numberOfUrlFirst; i++ {
		urlFirsts = append(urlFirsts, createRandomUrlFirst(t, config))
	}
	require.Equal(t, numberOfUrlFirst, len(urlFirsts))
	require.NotEmpty(t, urlFirsts)
	actualUrlFirsts, err := testQueries.GetUrlFirstByUniqueId(context.Background(), config.ID)
	require.NoError(t, err)
	for i, url := range urlFirsts {
		require.Equal(t, config.ID, url.UniqueID)
		require.Equal(t, config.ID, actualUrlFirsts[i].UniqueID)
		require.Equal(t, urlFirsts[i].UrlHash, actualUrlFirsts[i].UrlHash)
	}
}
