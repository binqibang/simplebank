package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ0123456789"
)

func init() {
	// Generate different random seed by time.
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max.
func RandomInt(min, max int64) int64 {
	// Int63n() returns a non-negative int64 pseudo-random
	// number in [0,n).
	return min + rand.Int63n(max-min+1)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomString generates a specified length random string.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(Alphabet)
	for i := 0; i < n; i++ {
		ch := Alphabet[rand.Intn(k)]
		sb.WriteByte(ch)
	}
	return sb.String()
}

func RandomOwner() string {
	names, err := ReadTxt("D:\\GoProjects\\simplebank\\util\\names.txt")
	if err != nil {
		panic(err)
	}
	k := len(names)
	idx := rand.Intn(k)
	name := strings.Replace(names[idx], "\n", "", -1)
	return name
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "JPY", "CNY", "CAD", "GBP"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
