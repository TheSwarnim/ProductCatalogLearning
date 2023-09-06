package main

import "go_microservice_learning_1/models"

func main() {

}

var productList = []*models.Product{
	{
		Id:          1,
		Name:        "Toy Car",
		Description: "Classic red toy car",
		Quantity:    50,
	},
	{
		Id:          2,
		Name:        "Book",
		Description: "Hardcover novel",
		Quantity:    120,
	},
	{
		Id:          3,
		Name:        "Smartphone",
		Description: "Latest model smartphone",
		Quantity:    20,
	},
	{
		Id:          4,
		Name:        "Laptop",
		Description: "Powerful work laptop",
		Quantity:    15,
	},
	{
		Id:          5,
		Name:        "Board Game",
		Description: "Fun family board game",
		Quantity:    30,
	},
}
