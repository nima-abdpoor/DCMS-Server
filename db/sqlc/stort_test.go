package db

import (
	"DCMS/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func addConfigTx(t *testing.T) AddConfigTxResult {
	numberOfUrlFirst := 4
	numberOfUrlSecond := 2
	numberOfRegexes := 5
	numberOfRequestUrl := 3
	store := NewStore(testDB)
	config := createRandomConfig(t)
	var urlSeconds []UrlSecondTx
	urlSeconds = make([]UrlSecondTx, numberOfUrlSecond)
	var regexes []Regex
	for i := 0; i < numberOfUrlSecond; i++ {
		second := createRandomUrlSecond(t, config)
		for j := 0; j < numberOfRegexes; j++ {
			regexes = append(regexes, createRandomRegex(t, second))
		}
		urlSeconds[i] = UrlSecondTx{
			ID:       second.ID,
			UniqueID: second.UniqueID,
			UrlHash:  second.UrlHash,
			Regex:    regexes,
		}
	}
	arg := AddConfigTxParams{
		ID:                     util.RandomInt(0, 10000),
		NetworkType:            util.RandomNetworkType(),
		IsLive:                 util.RandomBoolean(),
		SaveResponse:           util.RandomBoolean(),
		SaveRequest:            util.RandomBoolean(),
		SaveError:              util.RandomBoolean(),
		SaveSuccess:            util.RandomBoolean(),
		RepeatIntervalTimeUnit: util.RandomTimeUnit(),
		RepeatInterval:         util.RandomInt(1, 10),
		RequiresBatteryNotLow:  util.RandomBoolean(),
		RequiresStorageNotLow:  util.RandomBoolean(),
		RequiresCharging:       util.RandomBoolean(),
		RequiresDeviceIdl:      util.RandomBoolean(),
		UrlHashFirst:           util.RandomUrlHashGenerator(numberOfUrlFirst),
		UrlSecond:              urlSeconds,
		RequestUrl:             util.RandomStringList(100, numberOfRequestUrl),
	}

	result, err := store.AddConfigTx(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, result)

	//check the config
	require.Equal(t, arg.ID, result.Config.ID)
	require.Equal(t, arg.NetworkType, result.Config.NetworkType)
	require.Equal(t, arg.IsLive, result.Config.IsLive)
	require.Equal(t, arg.SaveResponse, result.Config.SaveResponse)
	require.Equal(t, arg.SaveResponse, result.Config.SaveResponse)
	require.Equal(t, arg.SaveError, result.Config.SaveError)
	require.Equal(t, arg.SaveSuccess, result.Config.SaveSuccess)
	require.Equal(t, arg.RepeatIntervalTimeUnit, result.Config.RepeatIntervalTimeUnit)
	require.Equal(t, arg.RepeatInterval, result.Config.RepeatInterval)
	require.Equal(t, arg.RequiresStorageNotLow, result.Config.RequiresStorageNotLow)
	require.Equal(t, arg.RequiresBatteryNotLow, result.Config.RequiresBatteryNotLow)
	require.Equal(t, arg.RequiresCharging, result.Config.RequiresCharging)
	require.Equal(t, arg.RequiresDeviceIdl, result.Config.RequiresDeviceIdl)
	require.NotZero(t, result.Config.ID)

	actualConfig, err2 := store.q.GetConfig(context.Background(), result.Config.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, actualConfig)
	require.Equal(t, arg.ID, actualConfig.ID)
	require.Equal(t, arg.NetworkType, actualConfig.NetworkType)
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
		actualUrlSecond, err3 := store.q.GetUrlSecond(context.Background(), urlSecond.ID)
		actualRegex, err4 := store.q.GetRegexByUrlId(context.Background(), urlSecond.ID)
		require.NoError(t, err3)
		require.NoError(t, err4)
		require.Equal(t, urlSecond.ID, actualUrlSecond.ID)
		require.Equal(t, urlSecond.UniqueID, actualUrlSecond.UniqueID)
		//check the Regex
		for j, regex := range actualRegex {
			require.Equal(t, regex.UrlID, urlSecond.ID)
			require.Equal(t, arg.UrlSecond[i].Regex[j].Regex, regex.Regex)
			require.Equal(t, arg.UrlSecond[i].Regex[j].StartIndex, regex.StartIndex)
			require.Equal(t, arg.UrlSecond[i].Regex[j].FinishIndex, regex.FinishIndex)
			require.NotZero(t, regex.ID)
		}
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
	require.Equal(t, addConfigTxResult.Config.NetworkType, result.Config.NetworkType)
	require.Equal(t, addConfigTxResult.Config.SaveRequest, result.Config.SaveRequest)
	require.Equal(t, addConfigTxResult.Config.SaveResponse, result.Config.SaveResponse)
	require.Equal(t, addConfigTxResult.Config.SaveError, result.Config.SaveError)
	require.Equal(t, addConfigTxResult.Config.SaveSuccess, result.Config.SaveSuccess)
	require.Equal(t, addConfigTxResult.Config.RepeatInterval, result.Config.RepeatInterval)
	require.Equal(t, addConfigTxResult.Config.RepeatIntervalTimeUnit, result.Config.RepeatIntervalTimeUnit)
	require.Equal(t, addConfigTxResult.Config.RequiresBatteryNotLow, result.Config.RequiresBatteryNotLow)
	require.Equal(t, addConfigTxResult.Config.RequiresCharging, result.Config.RequiresCharging)
	require.Equal(t, addConfigTxResult.Config.RequiresStorageNotLow, result.Config.RequiresStorageNotLow)
	require.Equal(t, addConfigTxResult.Config.RequiresDeviceIdl, result.Config.RequiresDeviceIdl)

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
		for i2, regex := range urlSecond.Regex {
			require.Equal(t, regex.Regex, result.UrlSecond[i].Regex[i2].Regex)
			require.Equal(t, regex.StartIndex, result.UrlSecond[i].Regex[i2].StartIndex)
			require.Equal(t, regex.FinishIndex, result.UrlSecond[i].Regex[i2].FinishIndex)
		}
	}

	//test RequestUrl
	for i, requestUrl := range addConfigTxResult.RequestUrl {
		require.Equal(t, requestUrl.ID, result.RequestUrl[i].ID)
		require.Equal(t, requestUrl.RequestUrl, result.RequestUrl[i].RequestUrl)
		require.Equal(t, requestUrl.UniqueID, result.RequestUrl[i].UniqueID)
	}
}
