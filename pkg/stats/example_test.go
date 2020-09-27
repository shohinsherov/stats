package stats

import (
	"fmt"

	"github.com/shohinsherov/bank/v2/pkg/types"
)

func ExampleAvg() {
	card := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "car",
			Status:   types.StatusFail,
		},
		{
			ID:       2,
			Amount:   10,
			Category: "car",
			Status:   types.StatusFail,
		},
		{
			ID:       3,
			Amount:   10,
			Category: "car",
			Status:   types.StatusOk,
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
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   10,
			Category: "car",
			Status:   types.StatusFail,
		},
		{
			ID:       3,
			Amount:   10,
			Category: "photo",
			Status:   types.StatusInProgress,
		},
	}

	fmt.Println(TotalInCategory(card, "car"))

	// Output: 10
}
