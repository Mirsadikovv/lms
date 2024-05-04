package storage

import "backend_course/lms/api/models"

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
	TeacherStorage() TeacherStorage
}

type StudentStorage interface {
	Create(student models.Student) (string, error)
	Update(student models.Student) (string, error)
	UpdateActivity(student models.Activity) (string, error)
	GetAll(student models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	GetStudentById(student models.GetStudent) (models.GetStudent, error)
	Delete(id string) error
}

type TeacherStorage interface {
	Create(teacher models.Teacher) (string, error)
	Update(teacher models.Teacher) (string, error)
	GetAll(teacher models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error)
	GetTeacherById(teacher models.GetTeacher) (models.GetTeacher, error)
	Delete(id string) error
}
