package postgres

import (
	"backend_course/lms/api/models"
	"context"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateTeacher(t *testing.T) {
	teacherRepo := NewTeacher(db)
	id := uuid.New()
	reqTeacher := models.Teacher{
		Id:            id.String(),
		FirstName:     faker.Name(),
		LastName:      faker.LastName(),
		Subject_id:    "2e23f8ad-7d20-443e-819f-780c0545003b",
		Start_working: faker.Date(),
		Phone:         faker.Phonenumber(),
		Mail:          faker.Email()}

	_, err := teacherRepo.Create(context.Background(), reqTeacher)
	if assert.NoError(t, err) {
		createdTeacher, err := teacherRepo.GetTeacherById(context.Background(), id.String())
		if assert.NoError(t, err) {
			assert.Equal(t, reqTeacher.FirstName, createdTeacher.FirstName)
			assert.Equal(t, reqTeacher.LastName, createdTeacher.LastName)
			assert.Equal(t, reqTeacher.Subject_id, createdTeacher.Subject_id)
			assert.Equal(t, reqTeacher.Start_working, createdTeacher.Start_working)
			assert.Equal(t, reqTeacher.Phone, createdTeacher.Phone)
			assert.Equal(t, reqTeacher.Mail, createdTeacher.Mail)
		} else {
			return
		}
		fmt.Println("Created teacher", createdTeacher)
	}

}

func TestUpdateTeacher(t *testing.T) {
	teacherRepo := NewTeacher(db)

	reqTeacher := models.Teacher{
		Id:            "d093a028-3220-42b0-b09e-97e87cfad26d",
		FirstName:     faker.Name(),
		LastName:      faker.LastName(),
		Subject_id:    "2e23f8ad-7d20-443e-819f-780c0545003b",
		Start_working: faker.Date(),
		Phone:         faker.Phonenumber(),
		Mail:          faker.Email()}

	_, err := teacherRepo.Update(context.Background(), reqTeacher, "d093a028-3220-42b0-b09e-97e87cfad26d")
	if assert.NoError(t, err) {
		createdTeacher, err := teacherRepo.GetTeacherById(context.Background(), "d093a028-3220-42b0-b09e-97e87cfad26d")
		if assert.NoError(t, err) {
			assert.Equal(t, reqTeacher.FirstName, createdTeacher.FirstName)
			assert.Equal(t, reqTeacher.LastName, createdTeacher.LastName)
			assert.Equal(t, reqTeacher.Subject_id, createdTeacher.Subject_id)
			assert.Equal(t, reqTeacher.Start_working, createdTeacher.Start_working)
			assert.Equal(t, reqTeacher.Phone, createdTeacher.Phone)
			assert.Equal(t, reqTeacher.Mail, createdTeacher.Mail)
		} else {
			return
		}
		fmt.Println("Created teacher", createdTeacher)
	}
}

func TestGetTeacher(t *testing.T) {
	teacherRepo := NewTeacher(db)

	reqTeacher := models.Teacher{
		Id:            "d093a028-3220-42b0-b09e-97e87cfad26d",
		FirstName:     faker.Name(),
		LastName:      faker.LastName(),
		Subject_id:    "2e23f8ad-7d20-443e-819f-780c0545003b",
		Start_working: faker.Date(),
		Phone:         faker.Phonenumber(),
		Mail:          faker.Email()}

	_, err := teacherRepo.Update(context.Background(), reqTeacher, "d093a028-3220-42b0-b09e-97e87cfad26d")
	if assert.NoError(t, err) {
		createdTeacher, err := teacherRepo.GetTeacherById(context.Background(), "d093a028-3220-42b0-b09e-97e87cfad26d")
		if assert.NoError(t, err) {
			assert.Equal(t, reqTeacher.FirstName, createdTeacher.FirstName)
			assert.Equal(t, reqTeacher.LastName, createdTeacher.LastName)
			assert.Equal(t, reqTeacher.Subject_id, createdTeacher.Subject_id)
			assert.Equal(t, reqTeacher.Start_working, createdTeacher.Start_working)
			assert.Equal(t, reqTeacher.Phone, createdTeacher.Phone)
			assert.Equal(t, reqTeacher.Mail, createdTeacher.Mail)
		} else {
			return
		}
		fmt.Println("Created teacher", createdTeacher)
	}
}

func TestDeleteTeacher(t *testing.T) {
	teacherRepo := NewTeacher(db)

	id, err := teacherRepo.Delete(context.Background(), "a41b2094-e87a-4669-84a5-b8bd1f329938")
	if assert.NoError(t, err) {
		assert.Equal(t, "a41b2094-e87a-4669-84a5-b8bd1f329938", id)
		return
	}
	fmt.Println("deleted teacher", id)
}
