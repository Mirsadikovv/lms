package storage

import "backend_course/lms/api/models"

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
}

type StudentStorage interface {
	Create(student models.Student) (string, error)
	Update(student models.Student) (string, error)
	GetAll(student models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	GetStudentById(student models.GetStudent) (models.GetStudent, error)
	Delete(id string) error
}
