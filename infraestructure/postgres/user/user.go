package user

import (
	"context"
	"database/sql"
	"ecommerce/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	psqlInsert = "INSERT INTO users (id, email, password, details, created_at) VALUES ($1, $2, $3, $4, $5)"
	psqlGetAll = "SELECT id, email, password, details, created_at, updated_at FROM users"
)

type User struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) User {
	return User{db}
}

func (u User) Create(m *model.User) error {
	_, err := u.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Email,
		m.Password,
		m.CreatedAt,
		m.IsAdmin,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	query := psqlGetAll + " WHERE email = $1"
	row := u.db.QueryRow(
		context.Background(),
		query,
		email,
	)

	return u.scanRow(row)
}

func (u User) GetAll() (model.Users, error) {
	rows, err := u.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := model.Users{}
	for rows.Next() {
		user, err := u.scanRow(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (u User) scanRow(s pgx.Row) (model.User, error) {
	user := model.User{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.IsAdmin,
		&user.Details,
		&user.CreatedAt,
	)

	if err != nil {
		return user, err
	}

	user.UpdatedAt = updatedAtNull.Int64

	return user, nil

}
