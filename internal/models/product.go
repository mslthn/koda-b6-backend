package models

type Product struct {
	IDProduct   int    `json:"product_id"`
	Name        string `json:"name"`
	Desc        string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	IsFlashsale bool   `json:"isFlashsale"`
}

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Desc        string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required, numeric, min=0"`
	Quantity    int    `json:"quantity" binding:"required, min=0"`
	IsFlashsale bool   `json:"isFlashsale"`
}

type UpdateProductRequest struct {
	Name        *string `json:"name"`
	Desc        *string `json:"description"`
	Price       *int    `json:"price" binding:"omitempty, numeric, min=0"`
	Quantity    *int    `json:"quantity" binding:"omitempty, min=0"`
	IsFlashsale *bool   `json:"isFlashsale"`
}