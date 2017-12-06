package GildedRose

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RegularItem(t *testing.T) {
	item := Item{"+5 Dexterity Vest", 2, 8}

	testUpdateItem(t, &item, 1, 7)
	testUpdateItem(t, &item, 0, 6)
	testUpdateItem(t, &item, -1, 4)
	testUpdateItem(t, &item, -2, 2)
	testUpdateItem(t, &item, -3, 0)
	testUpdateItem(t, &item,-4, 0)
}

func Test_AgedBrie(t *testing.T) {
	item := Item{"Aged Brie", 2, 0}

	testUpdateItem(t, &item, 1, 1)
	testUpdateItem(t, &item, 0, 2)
	testUpdateItem(t, &item, -1, 4)
	testUpdateItem(t, &item, -2, 6)
}

func Test_AgedBrie50Cap(t *testing.T) {
	item := Item{"Aged Brie", 10, 48}

	testUpdateItem(t, &item, 9, 49)
	testUpdateItem(t, &item, 8, 50)
	testUpdateItem(t, &item, 7, 50)
}

func Test_Sulfuras(t *testing.T) {
	item := Item{"Sulfuras, Hand of Ragnaros", 0, 80}

	testUpdateItem(t, &item, 0, 80)
	testUpdateItem(t, &item, 0, 80)
}

func Test_BackstagePasses(t *testing.T) {
	item := Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20}

	testUpdateItem(t, &item, 14, 21)
	testUpdateItem(t, &item, 13, 22)
	testUpdateItem(t, &item, 12, 23)
	testUpdateItem(t, &item, 11, 24)
	testUpdateItem(t, &item, 10, 25)
	testUpdateItem(t, &item, 9, 27)
	testUpdateItem(t, &item, 8, 29)
	testUpdateItem(t, &item, 7, 31)
	testUpdateItem(t, &item, 6, 33)
	testUpdateItem(t, &item, 5, 35)
	testUpdateItem(t, &item, 4, 38)
	testUpdateItem(t, &item, 3, 41)
	testUpdateItem(t, &item, 2, 44)
	testUpdateItem(t, &item, 1, 47)
	testUpdateItem(t, &item, 0, 50)
	testUpdateItem(t, &item, -1, 0)
	testUpdateItem(t, &item, -2, 0)
}

func Test_Conjured(t *testing.T) {
	item := Item{"Conjured Mana Cake", 2, 10}
	t.Skip(item)

	testUpdateItem(t, &item, 1, 8)
	testUpdateItem(t, &item, 0, 6)
	testUpdateItem(t, &item, -1, 2)
	testUpdateItem(t, &item, -2, 0)
}


func testUpdateItem(t *testing.T, item *Item, expectedSellIn int, expectedQuality int) {
	UpdateItem(item)
	assert.Equal(t, expectedSellIn, item.sellIn)
	assert.Equal(t, expectedQuality, item.quality)
}
