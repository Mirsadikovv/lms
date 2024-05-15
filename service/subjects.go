package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
	"context"
)

type subjectService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewSubjectService(storage storage.IStorage, log logger.ILogger) subjectService {
	return subjectService{storage: storage,
		logger: log}
}

func (s subjectService) Create(ctx context.Context, subject models.Subject) (string, error) {
	// business logic
	id, err := s.storage.SubjectStorage().Create(ctx, subject)
	if err != nil {
		s.logger.Error("error while creating subject, err: ", logger.Error(err))
		return "", err
	}
	// business logic
	return id, nil
}

func (s subjectService) Update(ctx context.Context, subject models.Subject, id string) (string, error) {
	// business logic
	id, err := s.storage.SubjectStorage().Update(ctx, subject, id)
	if err != nil {
		s.logger.Error("error while updating subject, err: ", logger.Error(err))
		return "", err
	}
	// business logic
	return id, nil
}

func (s subjectService) GetAllSubjects(ctx context.Context, subject models.GetAllSubjectsRequest) (models.GetAllSubjectsResponse, error) {
	// business logic
	subjects, err := s.storage.SubjectStorage().GetAll(ctx, models.GetAllSubjectsRequest{})
	if err != nil {
		s.logger.Error("error while getting all subject, err: ", logger.Error(err))
		return subjects, err
	}
	// business logic
	return subjects, nil
}
func (s subjectService) GetSubjectById(ctx context.Context, external_id string) (models.GetSubject, error) {
	// business logic
	resp, err := s.storage.SubjectStorage().GetSubjectById(ctx, external_id)
	if err != nil {
		s.logger.Error("error while getting all subject, err: ", logger.Error(err))
		return resp, err
	}
	// business logic
	return resp, nil
}

func (s subjectService) Delete(ctx context.Context, id string) error {
	// business logic
	_, err := s.storage.SubjectStorage().Delete(ctx, id)
	if err != nil {
		s.logger.Error("error while deletting all subject, err: ", logger.Error(err))
		return err
	}
	// business logic
	return nil
}
