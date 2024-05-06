package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"fmt"
)

type studentService struct {
	storage storage.IStorage
}

func NewStudentService(storage storage.IStorage) studentService {
	return studentService{storage: storage}
}

func (s studentService) Create(student models.Student) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Create(student)
	if err != nil {
		fmt.Println("error while creating student, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) Update(student models.UpdateStudent, id string) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Update(student, id)
	if err != nil {
		fmt.Println("error while updating student, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) GetAllStudents(student models.Student) (models.GetAllStudentsResponse, error) {
	// business logic
	students, err := s.storage.StudentStorage().GetAll(models.GetAllStudentsRequest{})
	if err != nil {
		fmt.Println("error while getting all student, err: ", err)
		return students, err
	}
	// business logic
	return students, nil
}
func (s studentService) GetStudentById(student models.GetStudent) (models.GetStudent, error) {
	// business logic
	resp, err := s.storage.StudentStorage().GetStudentById(student)
	if err != nil {
		fmt.Println("error while getting all student, err: ", err)
		return resp, err
	}
	// business logic
	return resp, nil
}

func (s studentService) Delete(id string) error {
	// business logic
	err := s.storage.StudentStorage().Delete(id)
	if err != nil {
		fmt.Println("error while deletting all student, err: ", err)
		return err
	}
	// business logic
	return nil
}
