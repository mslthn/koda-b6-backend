package repository

import (
	"context"
	"errors"
	"fmt"
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

func (r *UserRepository) Create(user models.User) error {
	query := `INSERT INTO "users" (fullname, email, password, role_id) VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(context.Background(), query, user.Fullname, user.Email, user.Password, user.RoleID)

	return err
}

func (r *UserRepository) GetUser() ([]models.User, error) {
	query := `SELECT u.id, u.fullname, u.email, u.password, COALESCE(ur.role, '') as role, COALESCE(u.address, '') as address, COALESCE(u.phone, '') as phone, COALESCE(u.picture, '') as picture FROM "users" u
				LEFT JOIN "user_role" ur ON u.role_id = ur.id`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New("Get all users query errors : " + err.Error())
	}

	user, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, errors.New("Get all users collect row error : " + err.Error())
	}

	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `SELECT u.id, u.fullname, u.email, u.password, COALESCE(ur.role, '') as role, COALESCE(u.address, '') as address, COALESCE(u.phone, '') as phone, COALESCE(u.picture, '') as picture 
				FROM "users" u 
				LEFT JOIN "user_role" ur ON u.role_id = ur.id 
				WHERE u.email=$1`

	rows, err := r.db.Query(context.Background(), query, email)
	if err != nil {
		fmt.Printf("DEBUG GetByEmail: Query error: %v\n", err)
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[models.User])
	if err != nil {
		fmt.Printf("DEBUG GetByEmail: CollectOneRow error: %v\n", err)
		return nil, err
	}

	fmt.Printf("DEBUG GetByEmail: User found - %+v\n", user)
	return user, nil
}

func (r *UserRepository) GetById(id int) (*models.User, error) {
	fmt.Printf("Mencari User dengan ID: %d\n", id)
	query := `SELECT u.id, u.fullname, u.email, u.password, COALESCE(ur.role, '') as role, COALESCE(u.address, '') as address, COALESCE(u.phone, '') as phone, COALESCE(u.picture, '') as picture 
				FROM "users" u 
				LEFT JOIN "user_role" ur ON u.role_id = ur.id 
				WHERE u.id=$1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[models.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user models.User) error {
	query := `UPDATE "users" SET fullname=$1, email=$2, password=$3 WHERE id=$4`

	_, err := r.db.Exec(
		context.Background(),
		query,
		user.Fullname,
		user.Email,
		user.Password,
		user.ID,
	)
	return err
}

func (r *UserRepository) UpdatePasswordByEmail(email string, password string) error {
	query := `UPDATE "users" SET password=$1 WHERE email=$2`

	_, err := r.db.Exec(context.Background(), query, password, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM "users" WHERE id=$1`

	_, err := r.db.Exec(context.Background(), query, id)

	return err
}