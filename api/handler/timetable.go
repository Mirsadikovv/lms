package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router		/timetable [post]
// @Summary		creates a timetable
// @Description	This api creates a timetable and returns its id
// @Tags		timetable
// @Accept		json
// @Produce		json
// @Param		timetable body models.Timetable true "timetable"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateTimetable(c *gin.Context) {
	timetable := models.Timetable{}

	if err := c.ShouldBindJSON(&timetable); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Timetable().Create(timetable)
	if err != nil {
		handleResponse(c, "error while creating timetable", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

// @Router		/timetable/update/{id} [put]
// @Summary		updates a timetable
// @Description	This api updates a timetable and returns its id
// @Tags		timetable
// @Accept		json
// @Produce		json
// @Param		timetable body models.Timetable true "timetable"
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateTimetable(c *gin.Context) {

	timetable := models.Timetable{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating timetableId", http.StatusBadRequest, err.Error())
		return
	}
	timetable.Id = id

	if err := c.ShouldBindJSON(&timetable); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Timetable().Update(timetable, id)
	if err != nil {
		handleResponse(c, "error while updating timetable", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}

// @Router		/timetable/{id} [get]
// @Summary		get one timetable
// @Description	This api get timetable and returns its
// @Tags		timetable
// @Accept		json
// @Produce		json
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetTimetable(c *gin.Context) {

	// timetable := models.GetTimetable{}
	id := c.Param("id")

	// timetable.Id = cast.ToString(id)

	resp, err := h.Service.Timetable().GetTimetableById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			handleResponse(c, "timetable not found", http.StatusNotFound, err.Error())
			return
		}
		handleResponse(c, "error while getting timetable", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "request successful", http.StatusOK, resp)
}

// @Router		/timetable/{id} [delete]
// @Summary		delete one timetable
// @Description	This api for delete timetable
// @Tags		timetable
// @Accept		json
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteTimetable(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "error while validating timetableId", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Timetable().Delete(id)
	if err != nil {
		handleResponse(c, "error while deleting timetable", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Deleted successfully", http.StatusOK, nil)
}
