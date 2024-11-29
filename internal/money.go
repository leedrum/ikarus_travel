package internal

import (
	"fmt"
)

func FormatMoney(amount float64) string {
	if amount == 0 {
		return "0"
	}

	thousand := "."
	decimal := ","

	intAmount := int(amount)
	floatAmount := amount - float64(intAmount)
	amountStr := fmt.Sprintf("%d", intAmount)

	if intAmount < 1000 {
		if floatAmount > 0 {
			return fmt.Sprintf("%d%s%f", intAmount, decimal, floatAmount)
		} else {
			return fmt.Sprintf("%d", intAmount)
		}
	}

	amountLen := len(amountStr)
	for i := amountLen - 3; i > 0; i -= 3 {
		amountStr = amountStr[:i] + thousand + amountStr[i:]
	}
	if floatAmount > 0 {
		amountStr = fmt.Sprintf("%s%s%f", amountStr, decimal, floatAmount)
	}
	return amountStr
}
