package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	balance := 0.0

	for _, t := range transactions {
		if t.To == name {
			balance += t.Sum
		}

		if t.From == name {
			balance -= t.Sum
		}
	}

	return balance
}
