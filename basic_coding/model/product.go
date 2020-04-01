package model

type Product struct {
	Id    int
	Name  string
	Price int
}

type ProductInterface interface {
	Insert(Product)
	GetAll() Products
}

type Products struct {
	Item []Product
}

func (p *Products) Insert(data Product) {
	p.Item = append(p.Item, data)
}

func (p *Products) GetAll() []Product {
	return p.Item
}
