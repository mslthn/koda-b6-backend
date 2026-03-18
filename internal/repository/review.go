package repository

import (
	"context"
	"main/internal/models"

	"github.com/jackc/pgx/v5"
)

type ReviewRepository struct{
	db *pgx.Conn
}

func NewReviewRepository(db *pgx.Conn) *ReviewRepository{
	return &ReviewRepository{
		db: db,
	}
}

func (r *ReviewRepository) Create(ctx context.Context, rvw models.Review) error{
	query := `INSERT INTO review (product_id, user_id, review_description, rating) VALUES ($1, $2, #3, $4)`
	_, err := r.db.Exec(ctx, query, rvw.ProductID, rvw.UserID, rvw.ReviewDesc, rvw.Rating)
	return err
}

func (r *ReviewRepository) GetAll(ctx context.Context) ([]models.Review, error){
	query := `SELECT product_id, user_id, review_description, rating FROM review`
	row, err := r.db.Query(ctx, query)
	if err != nil{
		return nil, err
	}
	
	reviews, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Review])
	if err!= nil{
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepository) GetById(ctx context.Context, id int) (*models.Review, error){
	query := `SELECT product_id, user_id, review_description, rating FROM review WHERE id=$1`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil{
		return  nil, err
	}
	review, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Review])
	if err != nil{
		return nil, err
	}

	return &review, nil
}

func (r *ReviewRepository) Update(ctx context.Context, id int, rvw models.Review) error{
	query := `UPDATE review SET product_id=$1, user_id=$2, review_description=$3, rating=$4 WHERE id=$5`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *ReviewRepository) Delete(ctx context.Context, id int) error{
	query := `DELETE FROM review WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}