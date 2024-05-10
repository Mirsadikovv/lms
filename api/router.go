package api

import (
	"backend_course/lms/api/handler"
	"backend_course/lms/service"
	"backend_course/lms/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(store storage.IStorage, service service.IServiceManager) *gin.Engine {
	h := handler.NewStrg(store, service)

	r := gin.Default()

	r.POST("/student", h.CreateStudent)
	r.PUT("/student/update/:id", h.UpdateStudent)
	r.PATCH("/student/activity/:id", h.UpdateStudentActivity)
	r.GET("/student", h.GetAllStudents)
	r.GET("student/:external_id", h.GetStudent)
	r.DELETE("/student/:id", h.DeleteStudent)

	r.POST("/teacher", h.CreateTeacher)
	r.PUT("/teacher/update/:id", h.UpdateTeacher)
	r.GET("/teacher", h.GetAllTeachers)
	r.GET("teacher/:id", h.GetTeacher)
	r.DELETE("/teacher/:id", h.DeleteTeacher)

	r.POST("/subject", h.CreateSubject)
	r.PUT("/subject/update/:id", h.UpdateSubject)
	r.GET("/subject", h.GetAllSubjects)
	r.GET("subject/:id", h.GetSubject)
	r.DELETE("/subject/:id", h.DeleteSubject)

	r.POST("/timetable", h.CreateTimetable)
	r.PUT("/timetable/update/:id", h.UpdateTimetable)
	r.GET("timetable/:id", h.GetTimetable)
	r.DELETE("/timetable/:id", h.DeleteTimetable)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
