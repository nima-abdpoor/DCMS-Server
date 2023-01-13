package db

import (
	"DCMS/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomCustomer(t *testing.T) Customer {
	hashedPassword, err := util.HashedPassword(util.RandomString(int(util.RandomInt(5, 15))))
	require.NoError(t, err)
	arg := CreateCustomerParams{
		Username:    util.RandomString(int(util.RandomInt(5, 15))),
		Password:    hashedPassword,
		Info:        util.RandomString(int(util.RandomInt(5, 15))),
		Email:       util.RandomString(int(util.RandomInt(5, 15))),
		PackageName: util.RandomString(int(util.RandomInt(5, 15))),
		SdkUuid:     util.RandomString(int(util.RandomInt(5, 15))),
		SecretKey:   util.RandomString(int(util.RandomInt(5, 15))),
	}
	customer, err := testQueries.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.Equal(t, arg.Email, customer.Email)
	require.Equal(t, arg.Info, customer.Info)
	require.Equal(t, arg.SdkUuid, customer.SdkUuid)
	require.Equal(t, arg.Username, customer.Username)
	require.Equal(t, arg.PackageName, customer.PackageName)
	require.Equal(t, arg.SecretKey, customer.SecretKey)
	require.Equal(t, arg.Password, customer.Password)
	return customer
}
func TestQueries_CreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestQueries_DeleteCustomer(t *testing.T) {
	customer := createRandomCustomer(t)
	err := testQueries.DeleteCustomer(context.Background(), customer.ID)
	require.NoError(t, err)
	actualCustomer, err2 := testQueries.GetCustomer(context.Background(), customer.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, actualCustomer)
}

func TestQueries_GetCustomer(t *testing.T) {
	customer := createRandomCustomer(t)
	actualRegex, err := testQueries.GetCustomer(context.Background(), customer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actualRegex)
	require.Equal(t, customer.ID, actualRegex.ID)
	require.Equal(t, customer.Email, customer.Email)
	require.Equal(t, customer.Info, customer.Info)
	require.Equal(t, customer.SdkUuid, customer.SdkUuid)
	require.Equal(t, customer.Username, customer.Username)
	require.Equal(t, customer.PackageName, customer.PackageName)
	require.Equal(t, customer.SecretKey, customer.SecretKey)
	require.Equal(t, customer.Password, customer.Password)
}

func TestQueries_GetCustomerByUserName(t *testing.T) {
	customer := createRandomCustomer(t)
	actualRegex, err := testQueries.GetCustomerByUsername(context.Background(), customer.Username)
	require.NoError(t, err)
	require.NotEmpty(t, actualRegex)
	require.Equal(t, customer.ID, actualRegex.ID)
	require.Equal(t, customer.Email, customer.Email)
	require.Equal(t, customer.Info, customer.Info)
	require.Equal(t, customer.SdkUuid, customer.SdkUuid)
	require.Equal(t, customer.Username, customer.Username)
	require.Equal(t, customer.PackageName, customer.PackageName)
	require.Equal(t, customer.SecretKey, customer.SecretKey)
	require.Equal(t, customer.Password, customer.Password)
}

func TestQueries_UpdateCustomer(t *testing.T) {
	customer := createRandomCustomer(t)
	hashedPassword, err := util.HashedPassword(util.RandomString(int(util.RandomInt(5, 15))))
	require.NoError(t, err)
	arg := UpdateCustomerParams{
		ID:          customer.ID,
		Username:    util.RandomString(int(util.RandomInt(5, 15))),
		Password:    hashedPassword,
		Info:        util.RandomString(int(util.RandomInt(5, 15))),
		Email:       util.RandomString(int(util.RandomInt(5, 15))),
		PackageName: util.RandomString(int(util.RandomInt(5, 15))),
		SdkUuid:     util.RandomString(int(util.RandomInt(5, 15))),
		SecretKey:   util.RandomString(int(util.RandomInt(5, 15))),
	}
	updatedCustomer, err := testQueries.UpdateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedCustomer)
	require.Equal(t, arg.ID, updatedCustomer.ID)
	require.Equal(t, arg.Email, updatedCustomer.Email)
	require.Equal(t, arg.Info, updatedCustomer.Info)
	require.Equal(t, arg.SdkUuid, updatedCustomer.SdkUuid)
	require.Equal(t, arg.Username, updatedCustomer.Username)
	require.Equal(t, arg.PackageName, updatedCustomer.PackageName)
	require.Equal(t, arg.SecretKey, updatedCustomer.SecretKey)
	require.Equal(t, arg.Password, updatedCustomer.Password)
}
