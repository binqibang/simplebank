package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("..")
	if err != nil {
		t.Error(err)
	}
	require.NotEmpty(t, config)
	t.Log(config)
}
