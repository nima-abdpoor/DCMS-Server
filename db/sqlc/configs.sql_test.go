package db

import (
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomConfig(t *testing.T) Config {
	arg := CreateConfigParams{
		ID:                     util.RandomInt(1, 1000000),
		NetworkType:            util.RandomNetworkType(),
		IsLive:                 util.RandomBoolean(),
		SaveRequest:            util.RandomBoolean(),
		SaveResponse:           util.RandomBoolean(),
		SaveError:              util.RandomBoolean(),
		SaveSuccess:            util.RandomBoolean(),
		RepeatIntervalTimeUnit: util.RandomTimeUnit(),
		RepeatInterval:         util.RandomInt(1, 10),
		RequiresBatteryNotLow:  util.RandomBoolean(),
		RequiresStorageNotLow:  util.RandomBoolean(),
		RequiresCharging:       util.RandomBoolean(),
		RequiresDeviceIdl:      util.RandomBoolean(),
	}
	config, err := testQueries.CreateConfig(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, config)
	require.Equal(t, arg.ID, config.ID)
	require.Equal(t, arg.NetworkType, config.NetworkType)
	require.Equal(t, arg.IsLive, config.IsLive)
	require.Equal(t, arg.SaveRequest, config.SaveRequest)
	require.Equal(t, arg.SaveResponse, config.SaveResponse)
	require.Equal(t, arg.SaveError, config.SaveError)
	require.Equal(t, arg.SaveSuccess, config.SaveSuccess)
	require.Equal(t, arg.RepeatIntervalTimeUnit, config.RepeatIntervalTimeUnit)
	require.Equal(t, arg.RepeatInterval, config.RepeatInterval)
	require.Equal(t, arg.RequiresStorageNotLow, config.RequiresStorageNotLow)
	require.Equal(t, arg.RequiresBatteryNotLow, config.RequiresBatteryNotLow)
	require.Equal(t, arg.RequiresCharging, config.RequiresCharging)
	require.Equal(t, arg.RequiresDeviceIdl, config.RequiresDeviceIdl)
	return config
}

func TestQueries_CreateConfig(t *testing.T) {
	createRandomConfig(t)
}

func TestQueries_DeleteConfig(t *testing.T) {
	config := createRandomConfig(t)

	err := testQueries.DeleteConfig(context.Background(), config.ID)
	require.NoError(t, err)

	actualConfig, err2 := testQueries.GetConfig(context.Background(), config.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, actualConfig)
}

func TestQueries_GetConfig(t *testing.T) {
	config := createRandomConfig(t)
	actualConfig, err := testQueries.GetConfig(context.Background(), config.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actualConfig)
	require.Equal(t, config.ID, actualConfig.ID)
	require.Equal(t, config.NetworkType, actualConfig.NetworkType)
	require.Equal(t, config.IsLive, actualConfig.IsLive)
	require.Equal(t, config.SaveRequest, actualConfig.SaveRequest)
	require.Equal(t, config.SaveResponse, actualConfig.SaveResponse)
	require.Equal(t, config.SaveError, actualConfig.SaveError)
	require.Equal(t, config.SaveSuccess, actualConfig.SaveSuccess)
	require.Equal(t, config.RepeatIntervalTimeUnit, actualConfig.RepeatIntervalTimeUnit)
	require.Equal(t, config.RepeatInterval, actualConfig.RepeatInterval)
	require.Equal(t, config.RequiresStorageNotLow, actualConfig.RequiresStorageNotLow)
	require.Equal(t, config.RequiresBatteryNotLow, actualConfig.RequiresBatteryNotLow)
	require.Equal(t, config.RequiresCharging, actualConfig.RequiresCharging)
	require.Equal(t, config.RequiresDeviceIdl, actualConfig.RequiresDeviceIdl)
}

func TestQueries_ListConfigs(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomConfig(t)
	}
	arg := ListConfigsParams{
		Limit:  5,
		Offset: 5,
	}

	configs, err := testQueries.ListConfigs(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, configs, 5)
	for _, account := range configs {
		require.NotEmpty(t, account)
	}
}

func TestQueries_UpdateConfig(t *testing.T) {
	config := createRandomConfig(t)
	arg := UpdateConfigParams{
		ID:                     config.ID,
		IsLive:                 config.IsLive,
		NetworkType:            config.NetworkType,
		SaveResponse:           config.SaveResponse,
		SaveRequest:            config.SaveRequest,
		SaveError:              config.SaveError,
		SaveSuccess:            config.SaveSuccess,
		RepeatInterval:         config.RepeatInterval,
		RepeatIntervalTimeUnit: config.RepeatIntervalTimeUnit,
		RequiresBatteryNotLow:  config.RequiresBatteryNotLow,
		RequiresStorageNotLow:  config.RequiresStorageNotLow,
		RequiresCharging:       config.RequiresCharging,
		RequiresDeviceIdl:      config.RequiresDeviceIdl,
	}
	updatedConfig, err := testQueries.UpdateConfig(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedConfig)
	require.Equal(t, config.ID, updatedConfig.ID)
	require.Equal(t, arg.ID, updatedConfig.ID)
	require.Equal(t, arg.IsLive, updatedConfig.IsLive)
	require.Equal(t, arg.NetworkType, updatedConfig.NetworkType)
	require.Equal(t, arg.SaveRequest, updatedConfig.SaveRequest)
	require.Equal(t, arg.SaveResponse, updatedConfig.SaveResponse)
	require.Equal(t, arg.SaveError, updatedConfig.SaveError)
	require.Equal(t, arg.SaveSuccess, updatedConfig.SaveSuccess)
	require.Equal(t, arg.RepeatInterval, updatedConfig.RepeatInterval)
	require.Equal(t, arg.RepeatIntervalTimeUnit, updatedConfig.RepeatIntervalTimeUnit)
	require.Equal(t, arg.RequiresBatteryNotLow, updatedConfig.RequiresBatteryNotLow)
	require.Equal(t, arg.RequiresCharging, updatedConfig.RequiresCharging)
	require.Equal(t, arg.RequiresStorageNotLow, updatedConfig.RequiresStorageNotLow)
	require.Equal(t, arg.RequiresDeviceIdl, updatedConfig.RequiresDeviceIdl)
}
