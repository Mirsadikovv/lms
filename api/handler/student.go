package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/check"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

// @Router		/student [post]
// @Summary		creates a student
// @Description	This api creates a student and returns its id
// @Tags		student
// @Accept		json
// @Produce		json
// @Param		student body models.Student true "student"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateStudent(c *gin.Context) {
	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateYear(student.Age); err != nil {
		handleResponse(c, "error while validating student age, year: "+strconv.Itoa(student.Age), http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.Service.Student().Create(student)
	if err != nil {
		handleResponse(c, "error while creating student", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

// @Router		/student/update/{id} [put]
// @Summary		updates a student
// @Description	This api updates a student and returns its id
// @Tags		student
// @Accept		json
// @Produce		json
// @Param		student body models.Student true "student"
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateStudent(c *gin.Context) {

	student := models.UpdateStudent{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Student().Update(student, id)
	if err != nil {
		handleResponse(c, "error while updating student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}

// @Router		/student/ [get]
// @Summary		get all students
// @Description	This api get all students
// @Tags		student
// @Accept		json
// @Produce		json
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllStudents(c *gin.Context) {
	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "erro/studentr while parsing page", http.StatusBadRequest, err.Error())
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

// @Router		/student/{external_id} [get]
// @Summary		get one student
// @Description	This api get student and returns its
// @Tags		student
// @Accept		json
// @Produce		json
// @Param 		external_id path string true "external_id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetStudent(c *gin.Context) {

	student := models.GetStudent{}

	id := c.Param("external_id")

	student.External_id = cast.ToString(id)

	resp, err := h.Service.Student().GetStudentById(student)
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

// @Router		/student/{id} [delete]
// @Summary		delete one student
// @Description	This api for delete student
// @Tags		student
// @Accept		json
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteStudent(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Student().Delete(id)
	if err != nil {
		handleResponse(c, "error while deleting student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Deleted successfully", http.StatusOK, nil)
}

// UpdateStudentActivity обновляет активность студента
// @Router		/student/activity/{id} [patch]
// @Summary		updates student activity
// @Description	This api updates student activity and returns its id
// @Tags		student
// @Accept		json
// @Produce		json
// @Param		id path string true "student id"
// @Param		activity body models.Activity true "student activity"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateStudentActivity(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating studentId", http.StatusBadRequest, err.Error())
		return
	}

	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.Store.StudentStorage().UpdateActivity(activity); err != nil {
		handleResponse(c, "error while updating student activity", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated student activity successfully", http.StatusOK, id)
}
