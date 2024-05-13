package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"context"
	"fmt"
)

type teacherService struct {
	storage storage.IStorage
}

func NewTeacherService(storage storage.IStorage) teacherService {
	return teacherService{storage: storage}
}

func (s teacherService) Create(ctx context.Context, teacher models.Teacher) (string, error) {
	// business logic
	id, err := s.storage.TeacherStorage().Create(ctx, teacher)
	if err != nil {
		fmt.Println("error while creating teacher, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s teacherService) Update(ctx context.Context, teacher models.Teacher, id string) (string, error) {
	// business logic
	id, err := s.storage.TeacherStorage().Update(ctx, teacher, id)
	if err != nil {
		fmt.Println("error while updating teacher, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s teacherService) GetAllTeachers(ctx context.Context, teacher models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {
	// business logic
	teachers, err := s.storage.TeacherStorage().GetAll(ctx, models.GetAllTeachersRequest{})
	if err != nil {
		fmt.Println("error while getting all teacher, err: ", err)
		return teachers, err
	}
	// business logic
	return teachers, nil
}
func (s teacherService) GetTeacherById(ctx context.Context, id string) (models.GetTeacher, error) {
	// business logic
	resp, err := s.storage.TeacherStorage().GetTeacherById(ctx, id)
	if err != nil {
		fmt.Println("error while getting all teacher, err: ", err)
		return resp, err
	}
	// business logic
	return resp, nil
}

func (s teacherService) Delete(ctx context.Context, id string) error {
	// business logic
	_, err := s.storage.TeacherStorage().Delete(ctx, id)
	if err != nil {
		fmt.Println("error while deletting all teacher, err: ", err)
		return err
	}
	// business logic
	return nil
}

func (s teacherService) CheckLessonNow(ctx context.Context, id string) (models.TeacherLessonNow, error) {
	// business logic
	resp, err := s.storage.TeacherStorage().CheckLessonNow(ctx, id)
	if err != nil {
		fmt.Println("error while checking lesson now, err: ", err)
		return resp, err
	}
	// business logic
	return resp, nil
}
