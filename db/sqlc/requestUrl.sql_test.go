package db

import (
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomRequestUrl(t *testing.T, config Config) Requesturl {
	arg := CreateRequestUrlParams{
		UniqueID:   config.ID,
		RequestUrl: util.RandomString(100),
	}
	requestUrl, err := testQueries.CreateRequestUrl(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, requestUrl)
	require.Equal(t, arg.UniqueID, config.ID)
	require.Equal(t, arg.RequestUrl, requestUrl.RequestUrl)
	require.NotZero(t, requestUrl.ID)
	return requestUrl
}

func TestQueries_CreateRequestUrl(t *testing.T) {
	createRandomRequestUrl(t, createRandomConfig(t))
}

func TestQueries_DeleteRequestUrl(t *testing.T) {
	config := createRandomConfig(t)
	requestUrl := createRandomRequestUrl(t, config)
	err := testQueries.DeleteRequestUrl(context.Background(), requestUrl.ID)
	require.NoError(t, err)
	requestUrlSecond, err2 := testQueries.GetRequestUrl(context.Background(), requestUrl.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, requestUrlSecond)
}

func TestQueries_GetRequestUrl(t *testing.T) {
	config := createRandomConfig(t)
	requestUrl := createRandomRequestUrl(t, config)
	actualRequestUrl, err := testQueries.GetRequestUrl(context.Background(), requestUrl.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actualRequestUrl)
	require.Equal(t, requestUrl.ID, actualRequestUrl.ID)
	require.Equal(t, requestUrl.RequestUrl, actualRequestUrl.RequestUrl)
}

func TestQueries_ListRequestUrls(t *testing.T) {
	for i := 0; i < 10; i++ {
		config := createRandomConfig(t)
		createRandomRequestUrl(t, config)
	}
	arg := ListRequestUrlsParams{
		Limit:  5,
		Offset: 5,
	}

	requestUrls, err := testQueries.ListRequestUrls(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, requestUrls, 5)
	for _, url := range requestUrls {
		require.NotEmpty(t, url)
	}
}

func TestQueries_UpdateRequestUrl(t *testing.T) {
	config := createRandomConfig(t)
	requestUrl := createRandomRequestUrl(t, config)
	arg := UpdateRequestUrlParams{
		UniqueID:   config.ID,
		RequestUrl: util.RandomString(100),
	}
	updatedRequestUrl, err := testQueries.UpdateRequestUrl(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRequestUrl)
	require.Equal(t, requestUrl.ID, updatedRequestUrl.ID)
	require.Equal(t, arg.UniqueID, updatedRequestUrl.UniqueID)
	require.Equal(t, arg.RequestUrl, updatedRequestUrl.RequestUrl)
}
