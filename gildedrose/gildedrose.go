package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
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
					if item.SellIn < 11 {
						if item.Quality < 50 {
							item.Quality = item.Quality + 1
						}
					}
					if item.SellIn < 6 {
						if item.Quality < 50 {
							item.Quality = item.Quality + 1
						}
					}
				}
			}
		}

		if !item.isSulfuras() {
			item.SellIn = item.SellIn - 1
		}

		if item.SellIn < 0 {
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
	}
}

func (i *Item) isBackstagePasses() bool {
	return i.Name == "Backstage passes to a TAFKAL80ETC concert"
}

func (i *Item) isSulfuras() bool {
	return i.Name == "Sulfuras, Hand of Ragnaros"
}

func (i *Item) isAgedBrie() bool {
	return i.Name == "Aged Brie"
}
