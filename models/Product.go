package models

type Product struct {
	Id          int
	Name        string
	Description string
	Quantity    int
}

type Products []*Product
