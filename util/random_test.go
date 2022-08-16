package util

import "testing"

func TestRandomBalance(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Logf("Random Balance: %d\n", RandomMoney())
	}
}

func TestRandomOwner(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Logf("Random Owner: %s\n", RandomOwner())
	}
}

func TestRandomCurrency(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Logf("Random Owner: %s\n", RandomCurrency())
	}
}

func TestRandomString(t *testing.T) {
	t.Log(RandomString(32))
}
