package gildedrose_test

import (
	"clover.yu/refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, -1, items[0].SellIn)
	assert.Equal(t, 0, items[0].Quality)
}
