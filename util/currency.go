package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSupportedCurrency(currencry string) bool {
	switch currencry {
	case USD, EUR, CAD:
		return true
	}
	return false
}
