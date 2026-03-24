package service

import (
	"context"
	"errors"
	"main/internal/models"
	"main/internal/repository"
)

type ProductImageService struct{
	repo *repository.ProductImageRepository
}

func NewProductImageService(repo *repository.ProductImageRepository) *ProductImageService{
	return &ProductImageService{repo:repo}
}

func (s *ProductImageService) Create(ctx context.Context, img models.ProductImage) error{
	return s.repo.Create(ctx, img)
}

func (s *ProductImageService) GetAll(ctx context.Context) ([]models.ProductImage, error){
	return  s.repo.FindAll(ctx)
}

func (s *ProductImageService) GetById(ctx context.Context, id int) (*models.ProductImage, error){
	return  s.repo.FindByID(ctx, id)
}

func (s *ProductImageService) Update(ctx context.Context, id int, req models.UpdateProductImageRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("Image not found!")
	}

	if req.ProductID != nil {
		existing.ProductID = *req.ProductID
	}
	if req.Image_url != nil {
		existing.Image_url = *req.Image_url
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *ProductImageService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

