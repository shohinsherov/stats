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

//  CategoriesAvg считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	catAvg := map[types.Category]types.Money{}
	for _, payment := range payments {
		catAvg[payment.Category] = payment.Amount
	}
	for key := range catAvg {
		sum := types.Money(0)
		n := 0
		for _, payment := range payments {
			if payment.Category == key {
				sum += payment.Amount
				n++
			}

		}
		catAvg[key] = sum / types.Money(n)
	}

	return catAvg
}

func PeriodsDynamic(first map[types.Category]types.Money, 
					second map[types.Category]types.Money,) map[types.Category]types.Money {

	result := map[types.Category]types.Money{}

	for key, value := range second {
		result[key] = value
	}

	for key, value := range first {
		result[key] -= value
	}

	return result
   }
   