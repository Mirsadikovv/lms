package models

type Student struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	External_id string `json:"external_id"`
	Phone       string `json:"phone"`
	Mail        string `json:"mail"`
	Pasword     string `json:"pasword"`
	IsActive    bool   `json:"isactive"`
}

type GetStudent struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	External_id string `json:"external_id"`
	Phone       string `json:"phone"`
	Mail        string `json:"mail"`
}

type GetAllStudentsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllStudentsResponse struct {
	Students []GetStudent `json:"students"`
	Count    int64        `json:"count"`
}

type Activity struct {
	Id       string `json:"id"`
	IsActive bool   `json:"activity"`
}
