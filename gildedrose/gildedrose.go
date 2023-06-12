package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func PassOneDay(items []*Item) {
	for _, item := range items {
		item.passByDay()
	}
}

func (item *Item) passByDay() {
	item.updateSellInDays()
	item.updateQuality()

	if item.isExpired() {
		item.updateQualityAfterExpiration()
	}
}

func (item *Item) updateQuality() {
	if !item.isAgedBrie() && !item.isBackstagePasses() {
		if item.Quality > 0 {
			if !item.isSulfuras() {
				item.Quality = item.Quality - 1
			}
		}
	} else {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
			if item.isBackstagePasses() {
				if item.SellIn < 10 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
				if item.SellIn < 5 {
					if item.Quality < 50 {
						item.Quality = item.Quality + 1
					}
				}
			}
		}
	}
}

func (item *Item) updateQualityAfterExpiration() {
	if !item.isAgedBrie() {
		if !item.isBackstagePasses() {
			if item.Quality > 0 {
				if !item.isSulfuras() {
					item.Quality = item.Quality - 1
				}
			}
		} else {
			item.Quality = item.Quality - item.Quality
		}
	} else {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
		}
	}
}

func (item *Item) isExpired() bool {
	return item.SellIn < 0
}

func (item *Item) updateSellInDays() {
	if !item.isSulfuras() {
		item.SellIn = item.SellIn - 1
	}
}

func (item *Item) isBackstagePasses() bool {
	return item.Name == "Backstage passes to a TAFKAL80ETC concert"
}

func (item *Item) isSulfuras() bool {
	return item.Name == "Sulfuras, Hand of Ragnaros"
}

func (item *Item) isAgedBrie() bool {
	return item.Name == "Aged Brie"
}
