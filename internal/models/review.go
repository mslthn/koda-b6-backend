package models

type Review struct {
	ID         int     `json:"id"`
	ProductID  int     `json:"product_id"`
	UserID     int     `json:"user_id"`
	ReviewDesc string  `json:"review_description"`
	Rating     float64 `json:"rating"`
}

type CreateReview struct {
	ProductID  int     `json:"product_id"`
	UserID     int     `json:"user_id"`
	ReviewDesc string  `json:"review_description"`
	Rating     float64 `json:"rating" binding:"min=1, max=5"`
}

type UpdateReview struct {
	ProductID  *int     `json:"product_id"`
	UserID     *int     `json:"user_id"`
	ReviewDesc *string  `json:"review_description"`
	Rating     *float64 `json:"rating" binding:"min=1, max=5"`
}

