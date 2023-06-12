package gildedrose_test

import (
	"clover.yu/refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_backstage_pass_quality_increased_by_1_and_sellIn_decrease_by_1_when_update_given_sellIn_more_than_10_and_quality_less_than_50(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 11, 8},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, 10, items[0].SellIn)
	assert.Equal(t, 9, items[0].Quality)
}

func Test_backstage_pass_quality_is_50_and_sellIn_decrease_by_1_when_update_given_sellIn_more_than_10_and_quality_is_50(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 13, 50},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, 12, items[0].SellIn)
	assert.Equal(t, 50, items[0].Quality)
}

func Test_backstage_pass_quality_increased_by_2_and_sellIn_decrease_by_1_when_update_given_sellIn_between_5_and_10_and_quality_less_than_49(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 8, 8},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, 7, items[0].SellIn)
	assert.Equal(t, 10, items[0].Quality)
}

func Test_backstage_pass_quality_is_50_and_sellIn_decrease_by_1_when_update_given_sellIn_between_5_and_10_and_quality_not_less_than_49(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 6, 49},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, 5, items[0].SellIn)
	assert.Equal(t, 50, items[0].Quality)
}

func Test_backstage_pass_quality_increased_by_3_and_sellIn_decrease_by_1_when_update_given_sellIn_within_5_and_quality_less_than_48(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 3, 44},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, 2, items[0].SellIn)
	assert.Equal(t, 47, items[0].Quality)
}

func Test_backstage_pass_quality_is_50_and_sellIn_decrease_by_1_when_update_given_sellIn_within_5_and_quality_not_less_than_48(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 2, 49},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, 1, items[0].SellIn)
	assert.Equal(t, 50, items[0].Quality)
}

func Test_backstage_pass_quality_is_0_and_sellIn_decrease_by_1_when_update_given_out_sell(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 0, 24},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, -1, items[0].SellIn)
	assert.Equal(t, 0, items[0].Quality)
}
