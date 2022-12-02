package db

import (
	"DCMS/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func addConfigTx(t *testing.T) AddConfigTxResult {
	numberOfUrlFirst := 1
	numberOfUrlSecond := 1
	numberOfRequestUrl := 1
	store := NewStore(testDB)
	config := createRandomConfig(t)
	var urlSeconds []Urlsecond
	urlSeconds = make([]Urlsecond, numberOfUrlSecond)
	for i := 0; i < numberOfUrlSecond; i++ {
		urlSeconds[i] = createRandomUrlSecond(t, config)
	}
	arg := AddConfigTxParams{
		ID:           util.RandomInt(0, 10000),
		SyncType:     util.RandomSyncType(),
		IsLive:       util.RandomBoolean(),
		UrlHashFirst: util.RandomUrlHashGenerator(numberOfUrlFirst),
		UrlSecond:    urlSeconds,
		RequestUrl:   util.RandomStringList(100, numberOfRequestUrl),
	}

	result, err := store.AddConfigTx(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, result)

	//check the config
	require.Equal(t, arg.ID, result.Config.ID)
	require.Equal(t, arg.SyncType, result.Config.SyncType)
	require.Equal(t, arg.IsLive, result.Config.IsLive)
	require.NotZero(t, result.Config.ID)

	actualConfig, err2 := store.q.GetConfig(context.Background(), result.Config.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, actualConfig)
	require.Equal(t, arg.ID, actualConfig.ID)
	require.Equal(t, arg.SyncType, actualConfig.SyncType)
	require.Equal(t, arg.IsLive, actualConfig.IsLive)

	//check the urlFirst
	require.NotEmpty(t, result.UrlFirst)
	for i, urlFirst := range result.UrlFirst {
		require.Equal(t, arg.ID, urlFirst.UniqueID)
		require.Equal(t, arg.UrlHashFirst[i], urlFirst.UrlHash)
		actualUrlFirst, err3 := store.q.GetUrlFirst(context.Background(), urlFirst.ID)
		require.NoError(t, err3)
		require.Equal(t, urlFirst.ID, actualUrlFirst.ID)
		require.Equal(t, urlFirst.UniqueID, actualUrlFirst.UniqueID)
	}

	//check the UrlSecond
	require.NotEmpty(t, result.UrlSecond)
	for i, urlSecond := range result.UrlSecond {
		require.Equal(t, arg.ID, urlSecond.UniqueID)
		require.Equal(t, arg.UrlSecond[i].UrlHash, urlSecond.UrlHash)
		require.Equal(t, arg.ID, urlSecond.UniqueID)
		require.Equal(t, arg.UrlSecond[i].Regex, urlSecond.Regex)
		require.Equal(t, arg.UrlSecond[i].StartIndex, urlSecond.StartIndex)
		require.Equal(t, arg.UrlSecond[i].FinishIndex, urlSecond.FinishIndex)
		actualUrlSecond, err3 := store.q.GetUrlSecond(context.Background(), urlSecond.ID)
		require.NoError(t, err3)
		require.Equal(t, urlSecond.ID, actualUrlSecond.ID)
		require.Equal(t, urlSecond.UniqueID, actualUrlSecond.UniqueID)
		require.Equal(t, urlSecond.Regex, actualUrlSecond.Regex)
		require.Equal(t, urlSecond.StartIndex, actualUrlSecond.StartIndex)
		require.Equal(t, urlSecond.FinishIndex, actualUrlSecond.FinishIndex)
	}

	//check the RequestUrl
	require.NotEmpty(t, result.RequestUrl)
	for i, requestUrl := range result.RequestUrl {
		require.Equal(t, arg.ID, requestUrl.UniqueID)
		require.Equal(t, arg.RequestUrl[i], requestUrl.RequestUrl)
		actualRequestUrl, err3 := store.q.GetRequestUrl(context.Background(), requestUrl.ID)
		require.NoError(t, err3)
		require.Equal(t, requestUrl.ID, actualRequestUrl.ID)
		require.Equal(t, requestUrl.UniqueID, actualRequestUrl.UniqueID)
		require.Equal(t, requestUrl.RequestUrl, actualRequestUrl.RequestUrl)
	}
	return result
}

func TestStore_AddConfigTx(t *testing.T) {
	addConfigTx(t)
}

func TestStore_GetConfigTx(t *testing.T) {
	addConfigTxResult := addConfigTx(t)
	store := NewStore(testDB)
	result, err := store.GetConfigTx(context.Background(), GetConfigTxParams{addConfigTxResult.Config.ID})
	require.NoError(t, err)
	require.NotEmpty(t, result)

	//test config
	require.Equal(t, addConfigTxResult.Config.ID, result.Config.ID)
	require.Equal(t, addConfigTxResult.Config.IsLive, result.Config.IsLive)
	require.Equal(t, addConfigTxResult.Config.SyncType, result.Config.SyncType)

	//test UrlFirst
	for i, urlFirst := range addConfigTxResult.UrlFirst {
		require.Equal(t, urlFirst.ID, result.UrlFirst[i].ID)
		require.Equal(t, urlFirst.UrlHash, result.UrlFirst[i].UrlHash)
		require.Equal(t, urlFirst.UniqueID, result.UrlFirst[i].UniqueID)
	}

	//test UrlSecond
	for i, urlSecond := range addConfigTxResult.UrlSecond {
		require.Equal(t, urlSecond.ID, result.UrlSecond[i].ID)
		require.Equal(t, urlSecond.UrlHash, result.UrlSecond[i].UrlHash)
		require.Equal(t, urlSecond.UniqueID, result.UrlSecond[i].UniqueID)
		require.Equal(t, urlSecond.Regex, result.UrlSecond[i].Regex)
		require.Equal(t, urlSecond.StartIndex, result.UrlSecond[i].StartIndex)
		require.Equal(t, urlSecond.FinishIndex, result.UrlSecond[i].FinishIndex)
	}

	//test RequestUrl
	for i, requestUrl := range addConfigTxResult.RequestUrl {
		require.Equal(t, requestUrl.ID, result.RequestUrl[i].ID)
		require.Equal(t, requestUrl.RequestUrl, result.RequestUrl[i].RequestUrl)
		require.Equal(t, requestUrl.UniqueID, result.RequestUrl[i].UniqueID)
	}
}
