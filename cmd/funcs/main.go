package main

import (
	"fmt"
	"github.com/shohinsherov/stats/v2/pkg/stats"
	"github.com/shohinsherov/bank/v2/pkg/types"
)

func main() {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	} 
	second := map[types.Category]types.Money{
		
		"food": 20,
	}

	result := stats.PeriodsDynamic(first, second)

	fmt.Println(result)
}