package main

import "log"

type Item interface {
	GetFullPrice() int
}

func main() {
	boxWItems := Box{
		price: 100,
		items: []Item{
			&Box{
				price: 50,
				items: []Item{
					&Product{price: 10},
					&Product{price: 2},
					&Box{
						price: 30,
						items: []Item{
							&Product{
								price: 1,
							},
						},
					},
				}},
		}}
	price := boxWItems.GetFullPrice()
	log.Println(price)
}
