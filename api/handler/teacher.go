package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Teacher().Create(c.Request.Context(), teacher)
	if err != nil {
		handleResponse(c, h.Log, "error while creating teacher", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
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
		handleResponse(c, h.Log, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}
	teacher.Id = id

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Teacher().Update(c.Request.Context(), teacher, id)
	if err != nil {
		handleResponse(c, h.Log, "error while updating teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Router		/teacher [get]
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
		handleResponse(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Service.Teacher().GetAllTeachers(c.Request.Context(), models.GetAllTeachersRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all teachers", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "request successful", http.StatusOK, resp)
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

	// teacher := models.GetTeacher{}
	id := c.Param("id")

	// teacher.Id = cast.ToString(id)

	resp, err := h.Service.Teacher().GetTeacherById(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			handleResponse(c, h.Log, "teacher not found", http.StatusNotFound, err.Error())
			return
		}
		handleResponse(c, h.Log, "error while getting teacher", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "request successful", http.StatusOK, resp)
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
		handleResponse(c, h.Log, "error while validating teacherId", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Teacher().Delete(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Deleted successfully", http.StatusOK, nil)
}

// @Router		/teacher/getlesson/{id} [get]
// @Summary		get lesson now
// @Description	This api get lesson now and returns its
// @Tags		teacher
// @Accept		json
// @Produce		json
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) TeacherCheckLessonNow(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.Service.Teacher().CheckLessonNow(c.Request.Context(), id)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	handleResponse(c, h.Log, "lesson not found", http.StatusNotFound, err.Error())
		// 	return
		// }
		handleResponse(c, h.Log, "error while getting lesson now", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "request successful", http.StatusOK, resp)
}
