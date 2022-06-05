package lib

type BalanceKey string

const (
	BalancePost BalanceKey = "POST"
	BalanceGet  BalanceKey = "GET"
)

func (balanceKey BalanceKey) ToKey() string {
	return string(balanceKey)
}

func (balanceKey BalanceKey) Description() string {
	switch balanceKey {
	case "POST":
		return "POST"
	case "GET":
		return "GET"
	default:
		return ""
	}
}
