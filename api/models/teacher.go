package models

type Teacher struct {
	Id            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Subject_id    string `json:"subject_id"`
	Start_working string `json:"start_working"`
	Phone         string `json:"phone"`
	Mail          string `json:"mail"`
}

type GetTeacher struct {
	Id            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Subject_id    string `json:"subject_id"`
	Start_working string `json:"start_working"`
	Phone         string `json:"phone"`
	Mail          string `json:"mail"`
	Created_at    string `json:"created_at"`
	Updated       string `json:"updated"`
}

type GetAllTeachersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllTeachersResponse struct {
	Teachers []GetTeacher `json:"teachers"`
	Count    int64        `json:"count"`
}
