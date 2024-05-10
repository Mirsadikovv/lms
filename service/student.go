package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"context"
	"fmt"
)

type studentService struct {
	storage storage.IStorage
}

func NewStudentService(storage storage.IStorage) studentService {
	return studentService{storage: storage}
}

func (s studentService) Create(ctx context.Context, student models.Student) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Create(ctx, student)
	if err != nil {
		fmt.Println("error while creating student, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) Update(ctx context.Context, student models.UpdateStudent, id string) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Update(ctx, student, id)
	if err != nil {
		fmt.Println("error while updating student, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) GetAllStudents(ctx context.Context, student models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
	// business logic
	students, err := s.storage.StudentStorage().GetAll(ctx, models.GetAllStudentsRequest{})
	if err != nil {
		fmt.Println("error while getting all student, err: ", err)
		return students, err
	}
	// business logic
	return students, nil
}
func (s studentService) GetStudentById(ctx context.Context, external_id string) (models.GetStudent, error) {
	// business logic
	resp, err := s.storage.StudentStorage().GetStudentById(ctx, external_id)
	if err != nil {
		fmt.Println("error while getting all student, err: ", err)
		return resp, err
	}
	// business logic
	return resp, nil
}

func (s studentService) Delete(ctx context.Context, id string) error {
	// business logic
	_, err := s.storage.StudentStorage().Delete(ctx, id)
	if err != nil {
		fmt.Println("error while deletting all student, err: ", err)
		return err
	}
	// business logic
	return nil
}

func (s studentService) UpdateActivity(ctx context.Context, student models.Activity) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().UpdateActivity(ctx, student)
	if err != nil {
		fmt.Println("error while updating student, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}
