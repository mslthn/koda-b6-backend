package models

type Product struct {
	IDProduct   int    `json:"product_id"`
	Name        string `json:"name"`
	Desc        string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"qusntity"`
	IsFlashsale bool   `json:"isFlashsale"`
}

type AddProduct struct {
	Name        string `json:"name"`
	Desc        string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"qusntity"`
	IsFlashsale bool   `json:"isFlashsale"`
}

type EditProduct struct {
	Name        string `json:"name"`
	Desc        string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"qusntity"`
	IsFlashsale bool   `json:"isFlashsale"`
}