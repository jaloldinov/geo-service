package postgres

import (
	"context"
	"fmt"
	"geo/models"
	"github.com/jackc/pgx"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}
func (b *userRepo) CreateUser(c context.Context, req *models.CreateUserReq) (*models.CreateUserRes, error) {
	user := models.CreateUserRes{}
	var lastInsertId int
	query := `INSERT INTO users(
					username, 
					password, 
					email
					) VALUES ($1, $2, $3) returning id`
	_, err := b.db.Exec(context.Background(), query,
		req.Username,
		req.Password,
		req.Email,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	user.ID = lastInsertId
	user.Email = req.Email
	user.Username = req.Username

	return &user, nil
}

/*
	func (b *userRepo) CreateUser(c context.Context, req *models.CreateUserReq) (*models.CreateUserRespond, error) {
		user := models.CreateUserRespond{}
		var lastInsertId int
		query := `INSERT INTO users(
						first_name,
						last_name,
						username,
						password
						) VALUES ($1, $2, $3, $4) returning id`
		_, err := b.db.Exec(context.Background(), query,
			req.FirstName,
			req.LastName,
			req.Username,
			req.Password,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}

		user.ID = lastInsertId
		user.Username = req.Username

		return &user, nil
	}

	func (b *userRepo) GetUserByID(c context.Context, req *models.IdRequest) (resp *models.User, err error) {
		u := models.User{}
		query := `SELECT
					id,
					first_name,
					last_name,
					username,
					password
					FROM users WHERE id = $1`

		err = b.db.QueryRow(context.Background(), query, req.ID).Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Username,
			&u.Password,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, fmt.Errorf("user not found")
			}
			return nil, fmt.Errorf("failed to get user: %w", err)
		}

		return &u, nil
	}
*/
func (b *userRepo) GetUserByEmail(c context.Context, req *models.LoginUserReq) (resp *models.User, err error) {
	u := models.User{}
	query := `SELECT
				id,
				email,
				username,
				password
				FROM users WHERE email = $1`

	err = b.db.QueryRow(context.Background(), query, req.Email).Scan(
		&u.ID,
		&u.Email,
		&u.Username,
		&u.Password,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &u, nil
}
