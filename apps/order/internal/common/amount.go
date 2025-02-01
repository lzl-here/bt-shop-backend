package common

import "github.com/shopspring/decimal"

func AmountEqual(a, b string) bool {
	am, _ := decimal.NewFromString(a)
	bm, _ := decimal.NewFromString(b)
	return am.Equals(bm)
}

func AmountPlus(a, b string) string {
	am, _ := decimal.NewFromString(a)
	bm, _ := decimal.NewFromString(b)
	return am.Add(bm).String()
}

func AmountMinus(a, b string) string {
	am, _ := decimal.NewFromString(a)
	bm, _ := decimal.NewFromString(b)
	return am.Sub(bm).String()
}
