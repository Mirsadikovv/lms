package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
	"context"
)

type studentService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewStudentService(storage storage.IStorage, log logger.ILogger) studentService {
	return studentService{storage: storage,
		logger: log}
}

func (s studentService) Create(ctx context.Context, student models.Student) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Create(ctx, student)
	if err != nil {
		s.logger.Error("error while creating student, err: ", logger.Error(err))
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) Update(ctx context.Context, student models.UpdateStudent, id string) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Update(ctx, student, id)
	if err != nil {
		s.logger.Error("error while updating student, err: ", logger.Error(err))
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) GetAllStudents(ctx context.Context, student models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
	// business logic
	students, err := s.storage.StudentStorage().GetAll(ctx, models.GetAllStudentsRequest{})
	if err != nil {
		s.logger.Error("error while getting all student, err: ", logger.Error(err))
		return students, err
	}
	// business logic
	return students, nil
}
func (s studentService) GetStudentById(ctx context.Context, external_id string) (models.GetStudent, error) {
	// business logic
	resp, err := s.storage.StudentStorage().GetStudentById(ctx, external_id)
	if err != nil {
		s.logger.Error("error while getting all student, err: ", logger.Error(err))
		return resp, err
	}
	// business logic
	return resp, nil
}

func (s studentService) Delete(ctx context.Context, id string) error {
	// business logic
	_, err := s.storage.StudentStorage().Delete(ctx, id)
	if err != nil {
		s.logger.Error("error while deletting all student, err: ", logger.Error(err))
		return err
	}
	// business logic
	return nil
}

func (s studentService) UpdateActivity(ctx context.Context, student models.Activity) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().UpdateActivity(ctx, student)
	if err != nil {
		s.logger.Error("error while updating student, err: ", logger.Error(err))
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) CheckLessonNow(ctx context.Context, id string) (models.StudentLessonNow, error) {
	// business logic
	resp, err := s.storage.StudentStorage().CheckLessonNow(ctx, id)
	if err != nil {
		s.logger.Error("error while checking lesson now, err: ", logger.Error(err))
		return resp, err
	}
	// business logic
	return resp, nil
}
