package stats

import (
	"reflect"
	"github.com/shohinsherov/bank/v2/pkg/types"
	"testing"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 10},
		{ID: 1, Category: "food", Amount: 10},
		{ID: 1, Category: "food", Amount: 10},
		{ID: 1, Category: "auto", Amount: 10},
		{ID: 1, Category: "auto", Amount: 10},
	}
	
	expected := map[types.Category]types.Money {
		"auto": 10,
		"food": 10,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expeced: %v, actual: %v",expected, result)
	}


}