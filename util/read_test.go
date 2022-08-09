package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadTxt(t *testing.T) {
	names, err := ReadTxt("names.txt")
	require.NoError(t, err)
	for i := 0; i < 10; i++ {
		t.Log(names[i])
	}
}
