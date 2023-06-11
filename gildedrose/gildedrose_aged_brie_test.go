package gildedrose_test

import (
	"clover.yu/refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_aged_brie_quality_increased_by_1_and_sellIn_decrease_by_1_when_update_given_in_sell_and_quality_less_than_50(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 3, 8},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, 2, items[0].SellIn)
	assert.Equal(t, 9, items[0].Quality)
}

func Test_aged_brie_quality_is_50_and_sellIn_decrease_by_1_when_update_given_in_sell_and_quality_is_50(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 3, 50},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, 2, items[0].SellIn)
	assert.Equal(t, 50, items[0].Quality)
}

func Test_aged_brie_quality_increased_by_2_and_sellIn_decrease_by_1_when_update_given_out_sell_and_quality_less_than_49(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 0, 8},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, -1, items[0].SellIn)
	assert.Equal(t, 10, items[0].Quality)
}

func Test_aged_brie_quality_is_50_and_sellIn_decrease_by_1_when_update_given_out_sell_and_quality_not_less_than_49(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 0, 49},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, -1, items[0].SellIn)
	assert.Equal(t, 50, items[0].Quality)
}
