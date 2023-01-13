package password

import (
	"DCMS/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	password := util.RandomString(6)
	hashedPassword, err := hashedPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	err = checkPassword(password, hashedPassword)
	require.NoError(t, err)

	err = checkPassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := util.RandomString(6)
	err = checkPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
