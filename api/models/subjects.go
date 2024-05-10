package models

type Subject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetSubject struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Created_at string `json:"created_at"`
	Updated    string `json:"updated"`
}

type GetAllSubjectsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllSubjectsResponse struct {
	Subjects []GetSubject `json:"subjects"`
	Count    int64        `json:"count"`
}
