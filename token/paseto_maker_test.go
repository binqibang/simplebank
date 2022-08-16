package token

import (
	"github.com/stretchr/testify/require"
	"simplebank/util"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {
	key := util.RandomString(32)
	maker, err := NewPasetoMaker(key)
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	t.Logf("%s", username)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	t.Logf("%s", token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	t.Logf("%+v", payload)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	t.Logf("%v", token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
