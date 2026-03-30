package repository

import (
	"context"
	"errors"
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

func (r *UserRepository) Create(user models.User) error{
	query := `INSERT INTO "user" (fullname, email, password) VALUE ($1, $2, $3)`

	_, err := r.db.Exec(context.Background(), query, user.Fullname, user.Email, user.Password)

	return err
}

func (r *UserRepository) GetUser() ([]models.User, error){
	query := `SELECT id, fullname, email, password, role FROM "user"
	JOIN "user_role" ON "user"."id" = "user_role"."id"`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil{
		return  nil, errors.New("Get all users query errors : "+ err.Error())
	}

	user, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, errors.New("Get all users collect row error : "+ err.Error())
	}
	

	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error){
	query := `SELECT id, fullname, email, password FROM "user" WHERE email=$1`

	rows, err := r.db.Query(context.Background(), query, email)
	if err != nil{
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[models.User])
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetById(id int) (*models.User, error){
	query := `SELECT id, fullname, email, password FROM "user" WHERE id=$1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil{
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[models.User])
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user models.User) error{
	query := `UPDATE user SET fullname=$1, email=$2, password=$3 WHERE id=$4`

	_, err := r.db.Exec(
		context.Background(),
		query,
		user.Fullname,
		user.Email,
		user.Password,
		user.ID,
	)
	return  err
}

func (r *UserRepository) UpdatePasswordByEmail(email string, password string) error{
	query := `UPDATE user SET password=$1 WHERE email=$2`

	_, err := r.db.Exec(context.Background(), query, password, email)
	if err != nil{
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error{
	query := `DELETE FROM user WHERE id=$1`

	_, err := r.db.Exec(context.Background(), query, id)

	return  err
}