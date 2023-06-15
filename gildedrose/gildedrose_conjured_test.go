package gildedrose_test

import (
	"clover.yu/refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_conjured_quality_decrease_by_1_and_sellIn_decrease_by_2_when_update_given_in_sell_and_quality_greater_than_1(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 3, 8},
	}

	gildedrose.PassOneDay(items)

	assert.Equal(t, 2, items[0].SellIn)
	assert.Equal(t, 6, items[0].Quality)
}

func Test_conjured_quality_is_0_and_sellIn_decrease_by_1_when_update_given_in_sell_and_quality_is_1(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 3, 1},
	}

	gildedrose.PassOneDay(items)

	assert.Equal(t, 2, items[0].SellIn)
	assert.Equal(t, 0, items[0].Quality)
}

func Test_conjured_quality_decrease_by_4_and_sellIn_decrease_by_1_when_update_given_out_sell_and_quality_greater_than_3(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 0, 8},
	}

	gildedrose.PassOneDay(items)

	assert.Equal(t, -1, items[0].SellIn)
	assert.Equal(t, 4, items[0].Quality)
}

func Test_conjured_quality_is_0_and_sellIn_decrease_by_1_when_update_given_out_sell_and_quality_not_greater_than_3(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 0, 3},
	}

	gildedrose.PassOneDay(items)

	assert.Equal(t, -1, items[0].SellIn)
	assert.Equal(t, 0, items[0].Quality)
}
