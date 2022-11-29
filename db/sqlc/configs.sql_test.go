package db

import (
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func createRandomConfig(t *testing.T) Config {
	arg := CreateConfigParams{
		ID:       strconv.Itoa(int(util.RandomInt(1, 1000000))),
		SyncType: util.RandomSyncType(),
		IsLive:   util.RandomBoolean(),
	}
	config, err := testQueries.CreateConfig(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, config)
	require.Equal(t, arg.ID, config.ID)
	require.Equal(t, arg.SyncType, config.SyncType)
	require.Equal(t, arg.IsLive, config.IsLive)
	require.NotZero(t, config.ID)
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
	require.Equal(t, config.SyncType, actualConfig.SyncType)
	require.Equal(t, config.IsLive, actualConfig.IsLive)
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
		ID:       config.ID,
		IsLive:   util.RandomBoolean(),
		SyncType: util.RandomSyncType(),
	}
	updatedConfig, err := testQueries.UpdateConfig(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedConfig)
	require.Equal(t, config.ID, updatedConfig.ID)
	require.Equal(t, arg.ID, updatedConfig.ID)
	require.Equal(t, arg.IsLive, updatedConfig.IsLive)
	require.Equal(t, arg.SyncType, updatedConfig.SyncType)
}
