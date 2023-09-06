package models

type Product struct {
	Id          int    `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

//type APIProduct struct {
//	Name        string
//	Description string
//	Quantity    int
//}

type Products []*Product
