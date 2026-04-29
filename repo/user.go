package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/noyandey/go-ddd-starter/domain"
	"github.com/noyandey/go-ddd-starter/internal/user"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
			email,
			password,
			first_name,
			last_name,
			username
		)
		VALUES (
			:email,
			:password,
			:first_name,
			:last_name,
			:username
		) RETURNING id
	`

	var userId int
	rows, err := r.db.NamedQuery(query, user)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&userId)
	}

	user.ID = userId
	return &user, nil
}

func (r *userRepo) Find(email, password string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT
			id,
			email,
			password,
			first_name,
			last_name,
			username
		FROM users
		WHERE email = $1 AND password = $2`

	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if sql.ErrNoRows == err {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
