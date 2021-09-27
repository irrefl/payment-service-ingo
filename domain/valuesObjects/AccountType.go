package valuesObjects

import "errors"

type AccountType int

const (
	Normal AccountType = iota
	Premium
	Gold
)

func GetAccountType(a AccountType) (string, error) {
	if a == 0 {
		return "Normal", nil
	}
	return "data", errors.New("empty string")
}
