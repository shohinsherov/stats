package stats

import (
	"github.com/shohinsherov/bank/v2/pkg/types"
)

// Avg рассчитает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	n := 0
	for _, pay := range payments {
		if pay.Status != types.StatusFail {
			n++
			sum += pay.Amount
		}
	}
	return sum / types.Money(n)
}

// TotalInCategory находим сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, pay := range payments {
		if pay.Category == category && pay.Status != types.StatusFail {
			sum += pay.Amount
		}
	}
	return sum
}
