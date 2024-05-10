package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"fmt"
)

type timetableService struct {
	storage storage.IStorage
}

func NewTimetableService(storage storage.IStorage) timetableService {
	return timetableService{storage: storage}
}

func (s timetableService) Create(timetable models.Timetable) (string, error) {
	// business logic
	id, err := s.storage.TimetableStorage().Create(timetable)
	if err != nil {
		fmt.Println("error while creating timetable, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s timetableService) Update(timetable models.Timetable, id string) (string, error) {
	// business logic
	id, err := s.storage.TimetableStorage().Update(timetable, id)
	if err != nil {
		fmt.Println("error while updating timetable, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

//	func (s timetableService) GetAllTimetables(timetable models.GetAllTimetablesRequest) (models.GetAllTimetablesResponse, error) {
//		// business logic
//		timetables, err := s.storage.TimetableStorage().GetAll(models.GetAllTimetablesRequest{})
//		if err != nil {
//			fmt.Println("error while getting all timetable, err: ", err)
//			return timetables, err
//		}
//		// business logic
//		return timetables, nil
//	}
func (s timetableService) GetTimetableById(id string) (models.GetTimetable, error) {
	// business logic
	resp, err := s.storage.TimetableStorage().GetTimetableById(id)
	if err != nil {
		fmt.Println("error while getting timetable, err: ", err)
		return resp, err
	}
	// business logic
	return resp, nil
}

func (s timetableService) Delete(id string) error {
	// business logic
	_, err := s.storage.TimetableStorage().Delete(id)
	if err != nil {
		fmt.Println("error while deletting all timetable, err: ", err)
		return err
	}
	// business logic
	return nil
}
