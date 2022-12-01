package db

import (
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func createRandomUrlSecond(t *testing.T, config Config) Urlsecond {
	arg := CreateUrlSecondParams{
		UniqueID:    config.ID,
		UrlHash:     strconv.Itoa(int(util.RandomUrlHashGenerator())),
		Regex:       util.RandomString(10),
		StartIndex:  int32(util.RandomInt(0, 40)),
		FinishIndex: int32(util.RandomInt(0, 40)),
	}
	urlSecond, err := testQueries.CreateUrlSecond(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, urlSecond)
	require.Equal(t, arg.UniqueID, config.ID)
	require.Equal(t, arg.UrlHash, urlSecond.UrlHash)
	require.Equal(t, arg.Regex, urlSecond.Regex)
	require.Equal(t, arg.StartIndex, urlSecond.StartIndex)
	require.Equal(t, arg.FinishIndex, urlSecond.FinishIndex)
	require.NotZero(t, urlSecond.ID)
	return urlSecond
}

func TestQueries_CreateUrlSecond(t *testing.T) {
	config := createRandomConfig(t)
	createRandomUrlSecond(t, config)
	createRandomUrlSecond(t, config)
	createRandomUrlSecond(t, config)
}

func TestQueries_DeleteUrlSecond(t *testing.T) {
	config := createRandomConfig(t)
	urlSecond := createRandomUrlSecond(t, config)
	err := testQueries.DeleteUrlSecond(context.Background(), urlSecond.ID)
	require.NoError(t, err)
	urlSecondActual, err2 := testQueries.GetUrlSecond(context.Background(), urlSecond.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, urlSecondActual)
}

func TestQueries_GetUrlSecond(t *testing.T) {
	config := createRandomConfig(t)
	urlSecond := createRandomUrlSecond(t, config)
	actualUrlSecond, err := testQueries.GetUrlSecond(context.Background(), urlSecond.UniqueID)
	require.NoError(t, err)
	require.NotEmpty(t, actualUrlSecond)
	require.Equal(t, urlSecond.ID, actualUrlSecond.ID)
	require.Equal(t, urlSecond.UrlHash, actualUrlSecond.UrlHash)
	require.Equal(t, urlSecond.Regex, actualUrlSecond.Regex)
	require.Equal(t, urlSecond.StartIndex, actualUrlSecond.StartIndex)
	require.Equal(t, urlSecond.FinishIndex, actualUrlSecond.FinishIndex)
}

func TestQueries_ListUrlSeconds(t *testing.T) {
	for i := 0; i < 10; i++ {
		config := createRandomConfig(t)
		createRandomUrlSecond(t, config)
	}
	arg := ListUrlSecondsParams{
		Limit:  5,
		Offset: 5,
	}

	urlSeconds, err := testQueries.ListUrlSeconds(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, urlSeconds, 5)
	for _, url := range urlSeconds {
		require.NotEmpty(t, url)
	}
}

func TestQueries_UpdateUrlSecond(t *testing.T) {
	config := createRandomConfig(t)
	urlSecond := createRandomUrlSecond(t, config)
	arg := UpdateUrlSecondParams{
		UniqueID:    urlSecond.UniqueID,
		UrlHash:     strconv.Itoa(int(util.RandomUrlHashGenerator())),
		Regex:       util.RandomString(10),
		StartIndex:  int32(util.RandomInt(0, 40)),
		FinishIndex: int32(util.RandomInt(0, 40)),
	}
	updatedUrlSecond, err := testQueries.UpdateUrlSecond(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUrlSecond)
	require.Equal(t, urlSecond.ID, updatedUrlSecond.ID)
	require.Equal(t, arg.UniqueID, updatedUrlSecond.UniqueID)
	require.Equal(t, arg.UrlHash, updatedUrlSecond.UrlHash)
	require.Equal(t, arg.StartIndex, updatedUrlSecond.StartIndex)
	require.Equal(t, arg.FinishIndex, updatedUrlSecond.FinishIndex)
	require.Equal(t, arg.Regex, updatedUrlSecond.Regex)
}
