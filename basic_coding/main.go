package main

import (
	"fmt"

	"github.com/rudiarta/kulina_test/basic_coding/model"
)

func main() {
	var products model.Products
	products.Insert(model.Product{
		Id:    1,
		Name:  "indomie",
		Price: 5000,
	})
	products.Insert(model.Product{
		Id:    2,
		Name:  "indomie goreng",
		Price: 50001,
	})
	products.Insert(model.Product{
		Id:    12,
		Name:  "indomie goreng",
		Price: 50001,
	})

	fmt.Println(products.GetAll())
	var sum int = 0
	for _, value := range products.GetAll() {
		sum = sum + value.Price
	}
	fmt.Println("SUM: ", sum)
}
