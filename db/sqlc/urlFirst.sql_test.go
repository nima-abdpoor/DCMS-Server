package db

import (
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func createRandomUrlFirst(t *testing.T, config Config) Urlfirst {
	arg := CreateUrlFirstParams{
		ID:      config.ID,
		UrlHash: strconv.Itoa(int(util.RandomUrlHashGenerator())),
	}
	urlFirst, err := testQueries.CreateUrlFirst(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, urlFirst)
	require.Equal(t, arg.ID, urlFirst.ID)
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
		ID:      urlFirst.ID,
		UrlHash: strconv.Itoa(int(util.RandomUrlHashGenerator())),
	}
	updatedUrlFirst, err := testQueries.UpdateUrlFirst(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUrlFirst)
	require.Equal(t, urlFirst.ID, updatedUrlFirst.ID)
	require.Equal(t, arg.ID, updatedUrlFirst.ID)
	require.Equal(t, arg.UrlHash, updatedUrlFirst.UrlHash)
}
