package storage

import (
	"backend_course/lms/api/models"
	"context"
	"time"
)

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
	TeacherStorage() TeacherStorage
	SubjectStorage() SubjectStorage
	TimetableStorage() TimetableStorage
	Redis() IRedisStorage
}

type StudentStorage interface {
	Create(ctx context.Context, student models.Student) (string, error)
	Update(ctx context.Context, student models.UpdateStudent, id string) (string, error)
	UpdateActivity(ctx context.Context, student models.Activity) (string, error)
	GetAll(ctx context.Context, student models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	GetStudentById(ctx context.Context, external_id string) (models.GetStudent, error)
	Delete(ctx context.Context, id string) (string, error)
	CheckLessonNow(ctx context.Context, id string) (models.StudentLessonNow, error)
}

type TeacherStorage interface {
	Create(ctx context.Context, teacher models.Teacher) (string, error)
	Update(ctx context.Context, teacher models.Teacher, id string) (string, error)
	GetAll(ctx context.Context, teacher models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error)
	GetTeacherById(ctx context.Context, id string) (models.GetTeacher, error)
	Delete(ctx context.Context, id string) (string, error)
	CheckLessonNow(ctx context.Context, id string) (models.TeacherLessonNow, error)
	GetTeacherByLogin(ctx context.Context, login string) (models.Teacher, error)
}

type SubjectStorage interface {
	Create(ctx context.Context, subject models.Subject) (string, error)
	Update(ctx context.Context, subject models.Subject, id string) (string, error)
	GetAll(ctx context.Context, req models.GetAllSubjectsRequest) (models.GetAllSubjectsResponse, error)
	GetSubjectById(ctx context.Context, id string) (models.GetSubject, error)
	Delete(ctx context.Context, id string) (string, error)
}

type TimetableStorage interface {
	Create(subject models.Timetable) (string, error)
	Update(subject models.Timetable, id string) (string, error)
	// GetAll(req models.GetAllTimetablesRequest) (models.GetAllTimetablesResponse, error)
	GetTimetableById(id string) (models.GetTimetable, error)
	Delete(id string) (string, error)
}

type IRedisStorage interface {
	SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) interface{}
	Del(ctx context.Context, key string) error
}
