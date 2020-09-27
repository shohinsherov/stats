package main

import (
	"fmt"
	"github.com/shohinsherov/stats/v2/pkg/stats"
	"github.com/shohinsherov/bank/v2/pkg/types"
)

func main() {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 20},
		{ID: 1, Category: "food", Amount: 10},
		{ID: 1, Category: "food", Amount: 10},
		{ID: 1, Category: "auto", Amount: 20},
		{ID: 1, Category: "auto", Amount: 20},
	}
	
	expected := map[types.Category]types.Money {
		"auto": 10,
		"food": 10,
	}

	result := stats.CategoriesAvg(payments)

	fmt.Println(expected)
	fmt.Println("-------------------")
	fmt.Println(result)
}