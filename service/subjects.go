package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"fmt"
)

type subjectService struct {
	storage storage.IStorage
}

func NewSubjectService(storage storage.IStorage) subjectService {
	return subjectService{storage: storage}
}

func (s subjectService) Create(subject models.Subject) (string, error) {
	// business logic
	id, err := s.storage.SubjectStorage().Create(subject)
	if err != nil {
		fmt.Println("error while creating subject, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s subjectService) Update(subject models.Subject, id string) (string, error) {
	// business logic
	id, err := s.storage.SubjectStorage().Update(subject, id)
	if err != nil {
		fmt.Println("error while updating subject, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s subjectService) GetAllSubjects(subject models.GetAllSubjectsRequest) (models.GetAllSubjectsResponse, error) {
	// business logic
	subjects, err := s.storage.SubjectStorage().GetAll(models.GetAllSubjectsRequest{})
	if err != nil {
		fmt.Println("error while getting all subject, err: ", err)
		return subjects, err
	}
	// business logic
	return subjects, nil
}
func (s subjectService) GetSubjectById(external_id string) (models.GetSubject, error) {
	// business logic
	resp, err := s.storage.SubjectStorage().GetSubjectById(external_id)
	if err != nil {
		fmt.Println("error while getting all subject, err: ", err)
		return resp, err
	}
	// business logic
	return resp, nil
}

func (s subjectService) Delete(id string) error {
	// business logic
	_, err := s.storage.SubjectStorage().Delete(id)
	if err != nil {
		fmt.Println("error while deletting all subject, err: ", err)
		return err
	}
	// business logic
	return nil
}
