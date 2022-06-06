package lib

type BalanceKey string

const (
	BalancePost      BalanceKey = "POST"
	BalanceGet       BalanceKey = "GET"
	BalanceHistory   BalanceKey = "HISTORY_"
	BalanceThreshold BalanceKey = "THRESHOLD_"
)

func (balanceKey BalanceKey) ToKey() string {
	return string(balanceKey)
}

func (balanceKey BalanceKey) Description() string {
	switch balanceKey {
	case "POST":
		return "Http Request Post"
	case "GET":
		return "Http Request Get"
	case "HISTORY_":
		return "History Key"
	case "THRESHOLD_":
		return "Threshold Key"
	default:
		return ""
	}
}
