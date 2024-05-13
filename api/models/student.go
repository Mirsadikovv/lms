package models

type Student struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         string `json:"age"`
	External_id string `json:"external_id"`
	Phone       string `json:"phone"`
	Mail        string `json:"mail"`
	Pasword     string `json:"pasword"`
	IsActive    bool   `json:"isactive"`
}

type UpdateStudent struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         string `json:"age"`
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
	Age         string `json:"age,omitempty"`
	External_id string `json:"external_id"`
	Phone       string `json:"phone,omitempty"`
	Mail        string `json:"mail,omitempty"`
	Created_at  string `json:"created_at"`
	Updated     string `json:"updated"`
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

type StudentLessonNow struct {
	SubjectName  string `json:"subject_name,omitempty"`
	TeacherName  string `json:"teacher_name,omitempty"`
	TimeUntilEnd string `json:"time_until_end,omitempty"`
}
