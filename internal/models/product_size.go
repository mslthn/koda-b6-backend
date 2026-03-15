package models

type ProductSize struct {
	ID    int    `json:"variant_id"`
	ProductID int    `json:"product_id"`
	SizeName  string `json:"variant_name"`
	AddPrice  int    `json:"additional_price"`
}

type CreateProductSizeRequest struct {
	ProductID int    `json:"product_id" binding:"required"`
	SizeName  string `json:"variant_name" binding:"required"`
	AddPrice  int    `json:"additional_price"`
}

type UpdateProductSizeRequest struct {
	ProductID *int    `json:"product_id"`
	SizeName  *string `json:"variant_name"`
	AddPrice  *int    `json:"additional_price"`
}
