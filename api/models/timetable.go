package models

type Timetable struct {
	Id         string `json:"id"`
	Teacher_id string `json:"teacher_id"`
	Student_id string `json:"student_id"`
	Subject_id string `json:"subject_id"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
}

type GetTimetable struct {
	Id       string `json:"id"`
	Teacher  string `json:"teacher"`
	Student  string `json:"student"`
	Subject  string `json:"subject"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type GetAllTimetablesRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllTimetablesResponse struct {
	Timetables []GetTimetable `json:"timetables"`
	Count      int64          `json:"count"`
}
