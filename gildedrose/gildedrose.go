package gildedrose

var itemHandlers = map[string]ItemHandler{
	"Aged Brie":                                 &AgedBrieHandler{},
	"Sulfuras, Hand of Ragnaros":                &SulfurasHandler{},
	"Backstage passes to a TAFKAL80ETC concert": &BackstagePassesHandler{},
}

type Item struct {
	Name            string
	SellIn, Quality int
}

type ItemHandler interface {
	PassOneDay(item *Item)
}

type AgedBrieHandler struct{}

func (handler *AgedBrieHandler) PassOneDay(item *Item) {
	updateQuality(item)

	updateSellIn(item)

	if isExpired(item) {
		updateQualityAfterExpiration(item)
	}
}

type BackstagePassesHandler struct{}

func (handler *BackstagePassesHandler) PassOneDay(item *Item) {
	updateQuality(item)

	updateSellIn(item)

	if isExpired(item) {
		updateQualityAfterExpiration(item)
	}
}

type SulfurasHandler struct{}

func (handler *SulfurasHandler) PassOneDay(item *Item) {
	updateQuality(item)

	updateSellIn(item)

	if isExpired(item) {
		updateQualityAfterExpiration(item)
	}
}

type NormalItemHandler struct{}

func (handler *NormalItemHandler) PassOneDay(item *Item) {
	updateQuality(item)

	updateSellIn(item)

	if isExpired(item) {
		updateQualityAfterExpiration(item)
	}
}

func PassOneDay(items []*Item) {
	for _, item := range items {
		if handler, ok := itemHandlers[item.Name]; ok {
			handler.PassOneDay(item)
		} else {
			(&NormalItemHandler{}).PassOneDay(item)
		}
	}
}

func updateQualityAfterExpiration(item *Item) {
	if !isAgedBrie(item) {
		if !isBackstagePasses(item) {
			if item.Quality > 0 {
				if !isSulfuras(item) {
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

func updateSellIn(item *Item) {
	if !isSulfuras(item) {
		item.SellIn = item.SellIn - 1
	}
}

func updateQuality(item *Item) {
	if !isAgedBrie(item) && !isBackstagePasses(item) {
		if item.Quality > 0 {
			if !isSulfuras(item) {
				item.Quality = item.Quality - 1
			}
		}
	} else {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
			if isBackstagePasses(item) {
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
}

func isBackstagePasses(item *Item) bool {
	return item.Name == "Backstage passes to a TAFKAL80ETC concert"
}

func isAgedBrie(item *Item) bool {
	return item.Name == "Aged Brie"
}

func isSulfuras(item *Item) bool {
	return item.Name == "Sulfuras, Hand of Ragnaros"
}

func isExpired(item *Item) bool {
	return item.SellIn < 0
}
