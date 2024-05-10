package service

import "backend_course/lms/storage"

type IServiceManager interface {
	Student() studentService
	Subject() subjectService
	Teacher() teacherService
	Timetable() timetableService
}

type Service struct {
	studentService   studentService
	subjectService   subjectService
	teacherService   teacherService
	timetableService timetableService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.studentService = NewStudentService(storage)
	services.teacherService = NewTeacherService(storage)
	services.subjectService = NewSubjectService(storage)
	services.timetableService = NewTimetableService(storage)

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
