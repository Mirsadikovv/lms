package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

// @Router		/teacher [post]
// @Summary		creates a teacher
// @Description	This api creates a teacher and returns its id
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		teacher body models.Teacher true "teacher"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateTeacher(c *gin.Context) {
	teacher := models.Teacher{}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.TeacherStorage().Create(teacher)
	if err != nil {
		handleResponse(c, "error while creating teacher", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

// @Router		/teacher/update/{id} [put]
// @Summary		updates a teacher
// @Description	This api updates a teacher and returns its id
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param		teacher body models.Teacher true "teacher"
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateTeacher(c *gin.Context) {

	teacher := models.Teacher{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}
	teacher.Id = id

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Store.TeacherStorage().Update(teacher)
	if err != nil {
		handleResponse(c, "error while updating teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}

// @Router		/teacher/ [get]
// @Summary		get all teachers
// @Description	This api get all teachers
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllTeachers(c *gin.Context) {
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

	resp, err := h.Store.TeacherStorage().GetAll(models.GetAllTeachersRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, "error while getting all teachers", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "request successful", http.StatusOK, resp)
}

// @Router		/teacher/{id} [get]
// @Summary		get one teacher
// @Description	This api get teacher and returns its
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetTeacher(c *gin.Context) {

	teacher := models.GetTeacher{}

	id := c.Param("id")

	teacher.Id = cast.ToString(id)

	resp, err := h.Store.TeacherStorage().GetTeacherById(teacher)
	if err != nil {
		if err == sql.ErrNoRows {
			handleResponse(c, "teacher not found", http.StatusNotFound, err.Error())
			return
		}
		handleResponse(c, "error while getting teacher", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "request successful", http.StatusOK, resp)
}

// @Router		/teacher/{id} [delete]
// @Summary		delete one teacher
// @Description	This api for delete teacher
// @Tags		teacher
// @Accept		json
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteTeacher(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Store.TeacherStorage().Delete(id)
	if err != nil {
		handleResponse(c, "error while deleting teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Deleted successfully", http.StatusOK, nil)
}
