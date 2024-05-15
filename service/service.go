package service

import (
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
)

type IServiceManager interface {
	Student() studentService
	Subject() subjectService
	Teacher() teacherService
	Timetable() timetableService
	Auth() authService
}

type Service struct {
	studentService   studentService
	subjectService   subjectService
	teacherService   teacherService
	timetableService timetableService
	authService      authService
	logger           logger.ILogger
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	services := Service{}
	services.studentService = NewStudentService(storage, log)
	services.teacherService = NewTeacherService(storage, log)
	services.subjectService = NewSubjectService(storage, log)
	services.timetableService = NewTimetableService(storage, log)
	services.authService = NewAuthService(storage, log)
	services.logger = log

	return services
}

func (s Service) Student() studentService {
	return s.studentService
}

func (s Service) Subject() subjectService {
	return s.subjectService
}

func (s Service) Teacher() teacherService {
	return s.teacherService
}

func (s Service) Timetable() timetableService {
	return s.timetableService
}

func (s Service) Auth() authService {
	return s.authService
}
