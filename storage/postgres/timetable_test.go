package postgres

import (
	"backend_course/lms/api/models"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateTimetable(t *testing.T) {
	timetableRepo := NewTimetable(db)
	id := uuid.New()
	reqTimetable := models.Timetable{
		Id:         id.String(),
		Teacher_id: "d093a028-3220-42b0-b09e-97e87cfad26d",
		Student_id: "06d289df-ae75-478e-ad15-278399b91a79",
		Subject_id: "ede6e2ad-879d-46f2-9848-4850d86eff11",
		FromDate:   "2024-05-10 11:00:00 AM",
		ToDate:     "2024-05-10 12:50:00 AM"}

	_, err := timetableRepo.Create(reqTimetable)
	if assert.NoError(t, err) {
		createdTimetable, err := timetableRepo.GetTimetableById(id.String())
		if assert.NoError(t, err) {
			// assert.Equal(t, reqTimetable.Teacher_id, createdTimetable.Teacher_id)
			// assert.Equal(t, reqTimetable.Student_id, createdTimetable.Student_id)
			// assert.Equal(t, reqTimetable.Subject_id, createdTimetable.Subject_id)
			assert.Equal(t, reqTimetable.FromDate, createdTimetable.FromDate)
			assert.Equal(t, reqTimetable.ToDate, createdTimetable.ToDate)
		} else {
			return
		}
		fmt.Println("Created timetable", createdTimetable)
	}

}

// func TestUpdateTimetable(t *testing.T) {
// 	timetableRepo := NewTimetable(db)

// 	reqTimetable := models.Timetable{
// 		Id:            "d093a028-3220-42b0-b09e-97e87cfad26d",
// 		FirstName:     faker.Name(),
// 		LastName:      faker.LastName(),
// 		Subject_id:    "2e23f8ad-7d20-443e-819f-780c0545003b",
// 		Start_working: faker.Date(),
// 		Phone:         faker.Phonenumber(),
// 		Mail:          faker.Email()}

// 	_, err := timetableRepo.Update(reqTimetable, "d093a028-3220-42b0-b09e-97e87cfad26d")
// 	if assert.NoError(t, err) {
// 		createdTimetable, err := timetableRepo.GetTimetableById("d093a028-3220-42b0-b09e-97e87cfad26d")
// 		if assert.NoError(t, err) {
// 			assert.Equal(t, reqTimetable.FirstName, createdTimetable.FirstName)
// 			assert.Equal(t, reqTimetable.LastName, createdTimetable.LastName)
// 			assert.Equal(t, reqTimetable.Subject_id, createdTimetable.Subject_id)
// 			assert.Equal(t, reqTimetable.Start_working, createdTimetable.Start_working)
// 			assert.Equal(t, reqTimetable.Phone, createdTimetable.Phone)
// 			assert.Equal(t, reqTimetable.Mail, createdTimetable.Mail)
// 		} else {
// 			return
// 		}
// 		fmt.Println("Created timetable", createdTimetable)
// 	}
// }

// func TestGetTimetable(t *testing.T) {
// 	timetableRepo := NewTimetable(db)

// 	reqTimetable := models.Timetable{
// 		Id:            "d093a028-3220-42b0-b09e-97e87cfad26d",
// 		FirstName:     faker.Name(),
// 		LastName:      faker.LastName(),
// 		Subject_id:    "2e23f8ad-7d20-443e-819f-780c0545003b",
// 		Start_working: faker.Date(),
// 		Phone:         faker.Phonenumber(),
// 		Mail:          faker.Email()}

// 	_, err := timetableRepo.Update(reqTimetable, "d093a028-3220-42b0-b09e-97e87cfad26d")
// 	if assert.NoError(t, err) {
// 		createdTimetable, err := timetableRepo.GetTimetableById("d093a028-3220-42b0-b09e-97e87cfad26d")
// 		if assert.NoError(t, err) {
// 			assert.Equal(t, reqTimetable.FirstName, createdTimetable.FirstName)
// 			assert.Equal(t, reqTimetable.LastName, createdTimetable.LastName)
// 			assert.Equal(t, reqTimetable.Subject_id, createdTimetable.Subject_id)
// 			assert.Equal(t, reqTimetable.Start_working, createdTimetable.Start_working)
// 			assert.Equal(t, reqTimetable.Phone, createdTimetable.Phone)
// 			assert.Equal(t, reqTimetable.Mail, createdTimetable.Mail)
// 		} else {
// 			return
// 		}
// 		fmt.Println("Created timetable", createdTimetable)
// 	}
// }

func TestDeleteTimetable(t *testing.T) {
	timetableRepo := NewTimetable(db)

	id, err := timetableRepo.Delete("c5ec7dc7-69d2-4001-b8cd-7cc6c183e774")
	if assert.NoError(t, err) {
		assert.Equal(t, "c5ec7dc7-69d2-4001-b8cd-7cc6c183e774", id)
		return
	}
	fmt.Println("deleted timetable", id)
}
