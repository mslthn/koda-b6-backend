package repository

import (
	"context"
	"main/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/pelletier/go-toml/query"
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

	row := r.db.QueryRow(context.Background(), query, email, code)

	var
}