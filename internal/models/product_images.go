package models

type ProductImage struct {
	ID   int    `json:"image_id"`
	ProductID int    `json:"product_id"`
	Image_url string `json:"image_url"`
}

type CreateProductImageRequest struct {
	ProductID int    `json:"product_id" binding:"required"`
	Image_url string `json:"image_url" binding:"required"`
}

type UpdateProductImageRequest struct {
	ProductID *int    `json:"product_id"`
	Image_url *string `json:"image_url"`
}
