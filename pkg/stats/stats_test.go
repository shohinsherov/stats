package stats

import (
	"reflect"
	"testing"

	"github.com/shohinsherov/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 10},
		{ID: 1, Category: "food", Amount: 10},
		{ID: 1, Category: "food", Amount: 10},
		{ID: 1, Category: "auto", Amount: 10},
		{ID: 1, Category: "auto", Amount: 10},
	}

	expected := map[types.Category]types.Money{
		"auto": 10,
		"food": 10,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expeced: %v, actual: %v", expected, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{

		"food": 20,
	}

	result := PeriodsDynamic(first, second)

	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expeced: %v, actual: %v", expected, result)
	}
}
