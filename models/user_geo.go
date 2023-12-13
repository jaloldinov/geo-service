package models

type CreateUserGeoRequest struct {
	UserID     int      `json:"user_id"`
	StartedAt  string   `json:"started_at"`
	FinishedAt string   `json:"finished_at"`
	Locations  Location `json:"locations"`
}

type CreateUserGeoRespond struct {
	Message string `json:"message"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GetUserGeoRequest struct {
	UserID     int    `json:"user_id"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
}

type UserGeoRespond struct {
	UserID                 int        `json:"user_id"`
	FirstConnectedLocation Location   `json:"first_connected_location"`
	Status                 bool       `json:"status"`
	RealTimeLocation       Location   `json:"real_time_location"`
	Distance               string     `json:"distance"`
	Locations              []Location `json:"locations"`
}
