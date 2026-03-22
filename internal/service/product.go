package service

import (
	"context"
	"errors"
	"main/internal/models"
	"main/internal/repository"
)

type ProductService struct{
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService{
	return  &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAll(ctx context.Context) ([]models.Product, error){
	return s.repo.GetAll(ctx)
}

func (s *ProductService) GetById(ctx context.Context ,id int) (*models.Product, error){
	return s.repo.FindById(ctx ,id)
}

func (s *ProductService) AddProduct(ctx context.Context ,add models.CreateProductRequest) error{
	if add.Price <= 0{
		return errors.New("Price shouldn't be 0 or less")
	}
	product := models.Product{
		Name: add.Name,
		Desc: add.Desc,
		Price: add.Price,
		Quantity: add.Quantity,
	}
	return s.repo.Create(ctx, product)
}

func (s *ProductService) Update(ctx context.Context, id int, req models.UpdateProductRequest) error{
	prod, err := s.repo.FindById(ctx, id)

	if err != nil {
		return errors.New("Couldn't find product")
	}
	if req.Name != nil{
		prod.Name = *req.Name
	}
	if req.Desc != nil{
		prod.Desc = *req.Desc
	}
	if req.Price != nil{
		if *req.Price > 0{
			prod.Price = *req.Price
		}
	}
	if req.Quantity != nil{
		if *req.Quantity > 0{
			prod.Quantity = *req.Quantity
		}
	}
	if req.IsFlashsale != nil{
		prod.IsFlashsale = *req.IsFlashsale
	}

	return s.repo.Update(ctx, id, *prod)
}

func (s *ProductService) Delete(ctx context.Context, id int) error{
	return s.repo.Delete(ctx, id)
}