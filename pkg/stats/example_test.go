package stats

import (
	"fmt"

	"github.com/shohinsherov/bank/pkg/types"
)

func ExampleAvg() {
	card := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "car",
		},
		{
			ID:       2,
			Amount:   10,
			Category: "car",
		},
		{
			ID:       3,
			Amount:   10,
			Category: "car",
		},
	}

	result := Avg(card)
	fmt.Println(result)

	// Output: 10
}

func ExampleTotalInCategory() {
	card := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "car",
		},
		{
			ID:       2,
			Amount:   10,
			Category: "car",
		},
		{
			ID:       3,
			Amount:   10,
			Category: "photo",
		},
	}

	fmt.Println(TotalInCategory(card, "car"))

	// Output: 20
}
