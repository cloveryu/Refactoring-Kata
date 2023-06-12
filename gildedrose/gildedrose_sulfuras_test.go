package gildedrose_test

import (
	"clover.yu/refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sulfuras_quality_and_sellIn_not_change_when_update(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 3, 66},
	}

	gildedrose.PassOneDay(items)

	assert.Equal(t, 3, items[0].SellIn)
	assert.Equal(t, 66, items[0].Quality)
}
