package utils

const (
	USD = "USD"
	EUR = "EUR"
	THB = "THB"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, THB:
		return true
	}
	return false
}
