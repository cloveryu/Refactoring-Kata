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
	increaseQuality(item)

	decreaseSellIn(item)

	if isExpired(item) {
		increaseQuality(item)
	}
}

type BackstagePassesHandler struct{}

func (handler *BackstagePassesHandler) PassOneDay(item *Item) {
	increaseQuality(item)

	if item.SellIn < 11 {
		increaseQuality(item)
	}
	if item.SellIn < 6 {
		increaseQuality(item)
	}

	decreaseSellIn(item)

	if isExpired(item) {
		item.Quality = 0
	}
}

type SulfurasHandler struct{}

func (handler *SulfurasHandler) PassOneDay(_ *Item) {}

type NormalItemHandler struct{}

func (handler *NormalItemHandler) PassOneDay(item *Item) {
	decreaseQuality(item)

	decreaseSellIn(item)

	if isExpired(item) {
		decreaseQuality(item)
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

func increaseQuality(item *Item) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
}

func decreaseQuality(item *Item) {
	if item.Quality > 0 {
		item.Quality = item.Quality - 1
	}
}

func decreaseSellIn(item *Item) {
	item.SellIn = item.SellIn - 1
}

func isExpired(item *Item) bool {
	return item.SellIn < 0
}
