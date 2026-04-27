package util

const (	
	INR = "INR"
	USD = "USD"
)


func IsSupportedCurrency(currency string) bool{
	switch currency{
	case INR,USD:
		return true
	}
	return false
}