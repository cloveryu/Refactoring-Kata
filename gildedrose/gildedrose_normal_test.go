package gildedrose_test

import (
	"clover.yu/refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_quality_and_sellIn_decrease_by_1_when_update_given_in_sell_and_quality_greater_than_0(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 3, 8},
	}

	gildedrose.PassOneDay(items)

	assert.Equal(t, 2, items[0].SellIn)
	assert.Equal(t, 7, items[0].Quality)
}

func Test_quality_is_0_and_sellIn_decrease_by_1_when_update_given_in_sell_and_quality_is_0(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 3, 0},
	}

	gildedrose.PassOneDay(items)

	assert.Equal(t, 2, items[0].SellIn)
	assert.Equal(t, 0, items[0].Quality)
}

func Test_quality_decrease_by_2_and_sellIn_decrease_by_1_when_update_given_out_sell_and_quality_greater_than_1(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 5},
	}

	gildedrose.PassOneDay(items)

	assert.Equal(t, -1, items[0].SellIn)
	assert.Equal(t, 3, items[0].Quality)
}

func Test_quality_is_0_and_sellIn_decrease_by_1_when_update_given_out_sell_and_quality_not_greater_than_1(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 1},
	}

	gildedrose.PassOneDay(items)

	assert.Equal(t, -1, items[0].SellIn)
	assert.Equal(t, 0, items[0].Quality)
}
