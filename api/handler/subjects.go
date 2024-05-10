package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router		/subject [post]
// @Summary		creates a subject
// @Description	This api creates a subject and returns its id
// @Tags		subject
// @Accept		json
// @Produce		json
// @Param		subject body models.Subject true "subject"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateSubject(c *gin.Context) {
	subject := models.Subject{}

	if err := c.ShouldBindJSON(&subject); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	// if err := check.ValidateYear(subject.Age); err != nil {
	// 	handleResponse(c, h.Log, "error while validating subject age, year: "+strconv.Itoa(subject.Age), http.StatusBadRequest, err.Error())

	// 	return
	// }

	id, err := h.Service.Subject().Create(c.Request.Context(), subject)
	if err != nil {
		handleResponse(c, h.Log, "error while creating subject", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Router		/subject/update/{id} [put]
// @Summary		updates a subject
// @Description	This api updates a subject and returns its id
// @Tags		subject
// @Accept		json
// @Produce		json
// @Param		subject body models.Subject true "subject"
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateSubject(c *gin.Context) {

	subject := models.Subject{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating subjectId", http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&subject); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Subject().Update(c.Request.Context(), subject, id)
	if err != nil {
		handleResponse(c, h.Log, "error while updating subject", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Router		/subject/ [get]
// @Summary		get all subjects
// @Description	This api get all subjects
// @Tags		subject
// @Accept		json
// @Produce		json
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllSubjects(c *gin.Context) {
	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "erro/subjectr while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Service.Subject().GetAllSubjects(c.Request.Context(), models.GetAllSubjectsRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all subjects", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "request successful", http.StatusOK, resp)
}

// @Router		/subject/{external_id} [get]
// @Summary		get one subject
// @Description	This api get subject and returns its
// @Tags		subject
// @Accept		json
// @Produce		json
// @Param 		external_id path string true "external_id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetSubject(c *gin.Context) {

	id := c.Param("external_id")

	resp, err := h.Service.Subject().GetSubjectById(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			handleResponse(c, h.Log, "subject not found", http.StatusNotFound, err.Error())
			return
		}
		handleResponse(c, h.Log, "error while getting subject", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "request successful", http.StatusOK, resp)
}

// @Router		/subject/{id} [delete]
// @Summary		delete one subject
// @Description	This api for delete subject
// @Tags		subject
// @Accept		json
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteSubject(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating subjectId", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Subject().Delete(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting subject", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Deleted successfully", http.StatusOK, nil)
}
