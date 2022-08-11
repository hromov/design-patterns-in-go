package main

type Product struct {
	price int
}

func (p *Product) GetFullPrice() int {
	return p.price
}
