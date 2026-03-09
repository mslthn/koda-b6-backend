package repository

import (
	"context"
	"main/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, role, fullname, email, password, address, phone, profile_picture FROM users`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id, &u.Role, &u.Fullname, &u.Email, &u.Password, &u.Address, &u.Phone, &u.ProfilePicture)
		if err != nil {
			return nil, err
		}
		users = append(users, u)

	}
	return users, nil
}

func (r *UserRepository) GetById(ctx context.Context, Id int) (*models.User, error) {
	query := `SELECT id, role, fullname, email, password, address, phone, profile_picture FROM users WHERE id = $1`

	var u models.User

	err := r.db.QueryRow(ctx, query, Id).Scan(&u.Id, &u.Role, &u.Fullname, &u.Email, &u.Password, &u.Address, &u.Phone, &u.ProfilePicture)

	if err != nil {
		return nil, err
	}
	return &u, nil

}

func (r *UserRepository) GetByEmail(ctx context.Context, Email string) (*models.User, error) {
	query := `SELECT id, role, fullname, email, password, address, phone, profile_picture FROM users WHERE email = $1`

	var u models.User

	err := r.db.QueryRow(ctx, query, Email).Scan(&u.Id, &u.Role, &u.Fullname, &u.Email, &u.Password, &u.Address, &u.Phone, &u.ProfilePicture)

	if err != nil {
		return nil, err
	}
	return &u, nil

}

func (r *UserRepository) CreateUser(ctx context.Context, )