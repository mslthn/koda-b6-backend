package models

type ProductVariant struct {
	IDV   int    `json:"variant_id"`
	ProductID   int    `json:"product_id"`
	VariantName string `json:"variant_name"`
	AddPrice    int    `json:"additional_price"`
}

type CreateProductVarianRequest struct{
	ProductID   int    `json:"product_id" binding:"required"`
	VariantName string `json:"variant_name" binding:"required"`
	AddPrice    int    `json:"additional_price"`
}

type UpdateProductVariantRequest struct{
	ProductID   *int    `json:"product_id"`
	VariantName *string `json:"variant_name"`
	AddPrice    *int    `json:"additional_price"`
}