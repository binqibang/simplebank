package util

const (
	USD = "USD"
	EUR = "EUR"
	JPY = "JPY"
	CNY = "CNY"
	GBP = "GBP"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, JPY, CNY, GBP:
		return true
	}
	return false
}
