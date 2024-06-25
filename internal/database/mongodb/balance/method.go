package balance

import "context"

func (hand *BalanceHandler) Add(balance ...*Balance) error {
	interfaces := make([]interface{}, len(balance))
	for i, b := range balance {
		interfaces[i] = b
	}
	_, err := hand.balance.InsertMany(
		context.Background(),
		interfaces,
	)
	return err
}
