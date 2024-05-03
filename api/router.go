package api

import (
	"backend_course/lms/api/handler"
	"backend_course/lms/storage"

	"github.com/gin-gonic/gin"
)

func New(store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store)

	r := gin.Default()

	r.POST("/student", h.CreateStudent)
	r.PUT("/update/:id", h.UpdateStudent)
	r.GET("/getall", h.GetAllStudents)
	// r.GET("/getstudent/:external_id", h.GetStudent)
	r.DELETE("/delete/:id", h.DeleteStudent)

	return r
}
