package repository

import (
	"context"
	"main/internal/models"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct{
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository{
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(ctx context.Context, p models.Product) error{
	query := `INSERT INTO "products" (name, description, price, quantity, isFlashSale) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, p.Name, p.Desc, p.Price, p.Quantity, p.IsFlashsale)
	return  err
}

func (r *ProductRepository) GetAll(ctx context.Context) ([]models.Product, error){
	query := `SELECT id, name, desc, price, quantity, is_flash_sale FROM products`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}
	
	return products, nil
}

func (r *ProductRepository) FindById(ctx context.Context, id int) (*models.Product, error){
	query := `SELECT id, name, desc, price, quantity, is_flash_sale FROM products WHERE id=$1`

	rows, err := r.db.Query(ctx, query, id)
	if err != nil{
		return  nil, err
	}

	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Update(ctx context.Context, id int, p models.Product) error{
	query := `UPDATE products SET name=$1, desc=$2, price=$3, quantity=$4, is_flash_sale=$5 WHERE id=%6`
	
	_, err := r.db.Exec(ctx, query, p.Name, p.Desc, p.Price, p.Quantity, p.IsFlashsale, id)
	return  err
}

func (r *ProductRepository) Delete(ctx context.Context, id int) error{
	query := `DELETE FROM products WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *ProductRepository) GetRecommendedProducts(ctx context.Context) ([]models.ProductLandingPage, error){
	query := `
			SELECT p.id, p.name, p.description, p.price
			count(rev.id) as num_review
			FROM products p
			join review rev
			on p.id = rev.product_id
			group by p.id
			order by num_review desc
			limit 4
			`
	rows, err := r.db.Query(ctx, query)
	if err != nil{
		return nil, err
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductLandingPage])
}