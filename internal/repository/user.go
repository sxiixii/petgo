/*
Package repository реализует репозитории для разных сущностей
*/
package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	DB *pgx.Conn
}

type User struct {
	Email string `db:"email"`
	Name  string `db:"name"`
}

func New(dbConn *pgx.Conn) *UserRepository {
	return &UserRepository{
		dbConn,
	}
}

func (u *UserRepository) Get(ctx context.Context, userID string) ([]User, error) {
	rows, err := u.DB.Query(ctx, "SELECT body ->> 'email' as email, body ->> '__name' as name FROM head.\"users\" where id = $1", userID)
	if err != nil {
		return []User{}, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	user, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		return []User{}, fmt.Errorf("collect rows failed: %w", err)
	}

	return user, nil
}
