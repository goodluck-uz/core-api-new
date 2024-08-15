package models

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	About      string `json:"about"`
	Avatar     string `json:"avatar"`
	Banner     string `json:"banner"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CountryID  int    `json:"country_id"`
	Role       string `json:"role"`
	CategoryID int    `json:"category_id"`
}

type CreateUserRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	About      string `json:"about"`
	Avatar     string `json:"avatar"`
	Banner     string `json:"banner"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CountryID  int    `json:"country_id"`
	Role       string `json:"role"`
	CategoryID int    `json:"category_id"`
}

type UpdateUserRequest struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	About      string `json:"about"`
	Avatar     string `json:"avatar"`
	Banner     string `json:"banner"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CountryID  int    `json:"country_id"`
	Role       string `json:"role"`
	CategoryID int    `json:"category_id"`
}

type GetListUserRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type ListUserResponse struct {
	Users []*User `json:"users"`
	Total int     `json:"total"`
}
