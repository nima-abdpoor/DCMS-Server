package db

import (
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomRegex(t *testing.T, urlSecond Urlsecond) Regex {
	arg := CreateRegexParams{
		UrlID:       urlSecond.ID,
		Regex:       util.RandomString(10),
		StartIndex:  int32(util.RandomInt(0, 40)),
		FinishIndex: int32(util.RandomInt(0, 40)),
	}
	regex, err := testQueries.CreateRegex(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, urlSecond)
	require.Equal(t, arg.UrlID, urlSecond.ID)
	require.Equal(t, arg.Regex, regex.Regex)
	require.Equal(t, arg.StartIndex, regex.StartIndex)
	require.Equal(t, arg.FinishIndex, regex.FinishIndex)
	require.NotZero(t, urlSecond.ID)
	return regex
}

func TestQueries_CreateRegex(t *testing.T) {
	createRandomRegex(t, createRandomUrlSecond(t, createRandomConfig(t)))
}

func TestQueries_DeleteRegex(t *testing.T) {
	config := createRandomConfig(t)
	urlSecond := createRandomUrlSecond(t, config)
	regex := createRandomRegex(t, urlSecond)
	err := testQueries.DeleteRegex(context.Background(), regex.ID)
	require.NoError(t, err)
	regexActual, err2 := testQueries.GetRegex(context.Background(), regex.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, regexActual)
}

func TestQueries_GetRegex(t *testing.T) {
	config := createRandomConfig(t)
	urlSecond := createRandomUrlSecond(t, config)
	regex := createRandomRegex(t, urlSecond)
	actualRegex, err := testQueries.GetRegex(context.Background(), regex.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actualRegex)
	require.Equal(t, regex.ID, actualRegex.ID)
	require.Equal(t, regex.Regex, actualRegex.Regex)
	require.Equal(t, regex.StartIndex, actualRegex.StartIndex)
	require.Equal(t, regex.FinishIndex, actualRegex.FinishIndex)
}

func TestQueries_GetRegexByUrlId(t *testing.T) {
	config := createRandomConfig(t)
	urlSecond := createRandomUrlSecond(t, config)
	var regexes []Regex
	numberOfRegexes := 10
	for i := 0; i < numberOfRegexes; i++ {
		regexes = append(regexes, createRandomRegex(t, urlSecond))
	}
	require.Equal(t, numberOfRegexes, len(regexes))
	require.NotEmpty(t, regexes)
	actualRegex, err := testQueries.GetRegexByUrlId(context.Background(), urlSecond.ID)
	require.NoError(t, err)
	for i, regex := range regexes {
		require.Equal(t, urlSecond.ID, regex.UrlID)
		require.Equal(t, urlSecond.ID, actualRegex[i].UrlID)
		require.Equal(t, regexes[i].Regex, actualRegex[i].Regex)
		require.Equal(t, regexes[i].StartIndex, actualRegex[i].StartIndex)
		require.Equal(t, regexes[i].FinishIndex, actualRegex[i].FinishIndex)
	}
}

func TestQueries_ListRegexes(t *testing.T) {
	for i := 0; i < 10; i++ {
		config := createRandomConfig(t)
		urlSecond := createRandomUrlSecond(t, config)
		createRandomRegex(t, urlSecond)
	}
	arg := ListRegexesParams{
		Limit:  5,
		Offset: 5,
	}

	regexes, err := testQueries.ListRegexes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, regexes, 5)
	for _, url := range regexes {
		require.NotEmpty(t, url)
	}
}

func TestQueries_UpdateRegex(t *testing.T) {
	config := createRandomConfig(t)
	urlSecond := createRandomUrlSecond(t, config)
	regex := createRandomRegex(t, urlSecond)
	arg := UpdateRegexParams{
		UrlID:       regex.UrlID,
		Regex:       util.RandomString(10),
		StartIndex:  int32(util.RandomInt(0, 40)),
		FinishIndex: int32(util.RandomInt(0, 40)),
	}
	updateRegex, err := testQueries.UpdateRegex(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updateRegex)
	require.Equal(t, urlSecond.ID, updateRegex.UrlID)
	require.Equal(t, arg.UrlID, updateRegex.UrlID)
	require.Equal(t, arg.StartIndex, updateRegex.StartIndex)
	require.Equal(t, arg.FinishIndex, updateRegex.FinishIndex)
	require.Equal(t, arg.Regex, updateRegex.Regex)
}
