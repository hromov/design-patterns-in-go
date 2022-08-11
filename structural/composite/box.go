package main

type Box struct {
	items []Item
	price int
}

func (b *Box) GetFullPrice() int {
	sum := b.price
	for _, item := range b.items {
		sum += item.GetFullPrice()
	}
	return sum
}
