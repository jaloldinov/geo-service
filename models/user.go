package models

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	AccessToken string
	ID          int    `json:"id"`
	Username    string `json:"username"`
}

/*
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type CreateUserReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type CreateUserRespond struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type IdRequest struct {
	ID int `json:"id"`
}

//type LoginUserReq struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}
//
//type LoginUserRes struct {
//	AccessToken string
//	ID          int    `json:"id"`
//	Username    string `json:"username"`
//}
*/
