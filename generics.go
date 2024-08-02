package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	reducer := func(val float64, t Transaction) float64 {
		if t.From == name {
			return val - t.Sum
		}

		if t.To == name {
			return val + t.Sum
		}

		return val
	}

	return Reduce(transactions, reducer, 0.0)
}

func Reduce[A, B any](transactions []A, fn func(B, A) B, initialValue B) B {
	res := initialValue

	for _, t := range transactions {
		res = fn(res, t)
	}

	return res
}
