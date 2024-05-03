package handler

import (
	"backend_course/lms/api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

func (h Handler) CreateStudent(c *gin.Context) {
	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.StudentStorage().Create(student)
	if err != nil {
		handleResponse(c, "error while creating student", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

func (h Handler) UpdateStudent(c *gin.Context) {

	student := models.Student{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	student.Id = id

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Store.StudentStorage().Update(student)
	if err != nil {
		handleResponse(c, "error while updating student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}

func (h Handler) GetAllStudents(c *gin.Context) {
	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Store.StudentStorage().GetAll(models.GetAllStudentsRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, "error while getting all students", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "request successful", http.StatusOK, resp)
}

func (h Handler) GetStudent(c *gin.Context) {

	student := models.GetStudent{}

	id := c.Param("external_id")

	student.External_id = cast.ToString(id)

	resp, err := h.Store.StudentStorage().GetStudentById(student)
	if err != nil {
		if err == sql.ErrNoRows {
			handleResponse(c, "student not found", http.StatusNotFound, err.Error())
			return
		}
		handleResponse(c, "error while getting student", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "request successful", http.StatusOK, resp)
}

func (h Handler) DeleteStudent(c *gin.Context) {

	student := models.Student{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}
	student.Id = id

	err := h.Store.StudentStorage().Delete(student)
	if err != nil {
		handleResponse(c, "error while updating student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, nil)
}
