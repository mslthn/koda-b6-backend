package repository

import (
	"context"
	"main/internal/models"

	"github.com/jackc/pgx/v5"
)

type AuthRepository struct{
	db *pgx.Conn
}

func RequestForgotPassword(db *pgx.Conn) *AuthRepository{
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateForgotPassword(data models.ForgotPassword) error {
	query := `INSERT INTO forgot_password (email, otp_code, created_at, expired_at) VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(
		context.Background(),
		query,
		data.Email,
		data.OTPCode,
		data.CreatedAt,
		data.ExpiredAt,
	)
	return err
}

func (r *AuthRepository) GetDataByEmailCode(email string, code string) (*models.ForgotPassword, error){
	query := `SELECT id, email, otp_code, created_at, expired_at FROM forgot_password WHERE email=$1 AND otp_code=$2`

	rows, err := r.db.Query(context.Background(), query, email, code)
	if err != nil {
		return  nil, err
	}

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil{
		return nil, err
	}
	return  &data, nil
}

func (r *AuthRepository) DeleteDataByCode(code string) error{
	query := `DELETE FROM forgot_password WHERE otp_code=$1`

	_, err := r.db.Exec(context.Background(), query, code)

	return  err
}