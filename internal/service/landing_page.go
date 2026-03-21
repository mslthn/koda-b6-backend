package service

import (
	"context"
	"main/internal/models"
	"main/internal/repository"
)

type LandingPageService struct {
	productRepo *repository.ProductRepository
	reviewRepo  *repository.ReviewRepository
}

func NewLandingPageService(pr *repository.ProductRepository, rr *repository.ReviewRepository) *LandingPageService{
	return &LandingPageService{
		productRepo: pr,
		reviewRepo: rr,
	}
}

func (s *LandingPageService) GetRecommendedProducts(ctx context.Context) ([]models.)