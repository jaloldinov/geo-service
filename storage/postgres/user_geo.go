package postgres

import (
	"context"
	"fmt"
	"geo/models"
	"geo/pkg/helper"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userGeoRepo struct {
	db *pgxpool.Pool
}

func NewUserGeoRepo(db *pgxpool.Pool) *userGeoRepo {
	return &userGeoRepo{
		db: db,
	}
}

func (b *userGeoRepo) CreateUserGeo(c context.Context, req *models.CreateUserGeoRequest) (*models.CreateUserGeoRespond, error) {

	query := `INSERT INTO users(
					user_id,
					started_at,
					finished_at,
					location
					) VALUES ($1, $2, $3, $4)`
	_, err := b.db.Exec(context.Background(), query,
		req.UserID,
		req.StartedAt,
		req.FinishedAt,
		req.Locations,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &models.CreateUserGeoRespond{Message: "created"}, nil
}

func (b *userGeoRepo) GetUserGeo(c context.Context, req *models.GetUserGeoRequest) (*models.UserGeoRespond, error) {
	var (
		respond models.UserGeoRespond
		filter  string
		params  = make(map[string]interface{})
	)

	// Filter with created at range
	if req.StartedAt != "" {
		filter += " AND started_at >= :started_at"
		params["started_at"] = req.StartedAt
	}

	if req.FinishedAt != "" {
		filter += " AND finished_at <= :finished_at"
		params["finished_at"] = req.FinishedAt
	}

	query := `
	SELECT 
		user_id,
		json_agg(location)
	FROM user_geo  
	WHERE true ` + filter + `
	GROUP BY user_id`

	q, arr := helper.ReplaceQueryParams(query, params)
	err := b.db.QueryRow(context.Background(), q, arr...).Scan(
		&respond.UserID,
		&respond.Locations,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	//respond.FirstConnectedLocation = []string{}
	//respond.Status = false
	//respond.RealTimeLocation = []string{}
	//respond.Distance = ""
	if err != nil {
		return nil, fmt.Errorf("error while getting rows: %w", err)
	}

	return &respond, nil
}

/*
func (b *userGeoRepo) GetUserByEmail(c context.Context, req *models.LoginUserReq) (resp *models.User, err error) {
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
*/
