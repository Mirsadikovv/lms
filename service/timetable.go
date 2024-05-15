package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
)

type timetableService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewTimetableService(storage storage.IStorage, logger logger.ILogger) timetableService {
	return timetableService{storage: storage}
}

func (s timetableService) Create(timetable models.Timetable) (string, error) {
	// business logic
	id, err := s.storage.TimetableStorage().Create(timetable)
	if err != nil {
		s.logger.Error("error while creating timetable, err: ", logger.Error(err))
		return "", err
	}
	// business logic
	return id, nil
}

func (s timetableService) Update(timetable models.Timetable, id string) (string, error) {
	// business logic
	id, err := s.storage.TimetableStorage().Update(timetable, id)
	if err != nil {
		s.logger.Error("error while updating timetable, err: ", logger.Error(err))
		return "", err
	}
	// business logic
	return id, nil
}

//	func (s timetableService) GetAllTimetables(timetable models.GetAllTimetablesRequest) (models.GetAllTimetablesResponse, error) {
//		// business logic
//		timetables, err := s.storage.TimetableStorage().GetAll(models.GetAllTimetablesRequest{})
//		if err != nil {
//			s.logger.Error("error while getting all timetable, err: ",logger.Error(err))
//			return timetables, err
//		}
//		// business logic
//		return timetables, nil
//	}
func (s timetableService) GetTimetableById(id string) (models.GetTimetable, error) {
	// business logic
	resp, err := s.storage.TimetableStorage().GetTimetableById(id)
	if err != nil {
		s.logger.Error("error while getting timetable, err: ", logger.Error(err))
		return resp, err
	}
	// business logic
	return resp, nil
}

func (s timetableService) Delete(id string) error {
	// business logic
	_, err := s.storage.TimetableStorage().Delete(id)
	if err != nil {
		s.logger.Error("error while deletting all timetable, err: ", logger.Error(err))
		return err
	}
	// business logic
	return nil
}
