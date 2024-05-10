package postgres

import (
	"backend_course/lms/api/models"
	"context"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	studentRepo := NewStudent(db)

	reqStudent := models.Student{
		Id:          "5f6b2b47-b599-4c16-9216-5a3695645df0",
		FirstName:   faker.Name(),
		Age:         "10",
		LastName:    faker.Word(),
		External_id: "80",
		Phone:       "1231313123",
		Mail:        "string",
		Pasword:     "string",
		IsActive:    true}

	_, err := studentRepo.Create(context.Background(), reqStudent)
	if assert.NoError(t, err) {
		createdStudent, err := studentRepo.GetStudentById(context.Background(), "80")
		if assert.NoError(t, err) {
			assert.Equal(t, reqStudent.FirstName, createdStudent.FirstName)
			assert.Equal(t, reqStudent.Age, createdStudent.Age)
			assert.Equal(t, reqStudent.LastName, createdStudent.LastName)
		} else {
			return
		}
		fmt.Println("Created student", createdStudent)
	}

}

func TestUpdateStudent(t *testing.T) {
	studentRepo := NewStudent(db)

	reqStudent := models.UpdateStudent{
		FirstName:   faker.Name(),
		Age:         "10",
		LastName:    faker.Word(),
		External_id: "78",
		Phone:       "1231313123",
		Mail:        "string",
		Pasword:     "string",
		IsActive:    true}

	_, err := studentRepo.Update(context.Background(), reqStudent, "f542b02c-5318-44a4-9c16-c54375f32985")
	if assert.NoError(t, err) {
		createdStudent, err := studentRepo.GetStudentById(context.Background(), "78")
		if assert.NoError(t, err) {
			assert.Equal(t, reqStudent.FirstName, createdStudent.FirstName)
			assert.Equal(t, reqStudent.Age, createdStudent.Age)
			assert.Equal(t, reqStudent.LastName, createdStudent.LastName)
		} else {
			return
		}
		fmt.Println("Created student", createdStudent)
	}
}

func TestGetStudent(t *testing.T) {
	studentRepo := NewStudent(db)

	reqStudent := models.UpdateStudent{
		FirstName:   faker.Name(),
		Age:         "10",
		LastName:    faker.Word(),
		External_id: "78",
		Phone:       "1231313123",
		Mail:        "string",
		Pasword:     "string",
		IsActive:    true}

	_, err := studentRepo.Update(context.Background(), reqStudent, "f542b02c-5318-44a4-9c16-c54375f32985")
	if assert.NoError(t, err) {
		createdStudent, err := studentRepo.GetStudentById(context.Background(), "78")
		if assert.NoError(t, err) {
			assert.Equal(t, reqStudent.FirstName, createdStudent.FirstName)
			assert.Equal(t, reqStudent.Age, createdStudent.Age)
			assert.Equal(t, reqStudent.LastName, createdStudent.LastName)
		} else {
			return
		}
		fmt.Println("Created student", createdStudent)
	}
}

func TestDeleteStudent(t *testing.T) {
	studentRepo := NewStudent(db)

	id, err := studentRepo.Delete(context.Background(), "f542b02c-5318-44a4-9c16-c54375f32985")
	if assert.NoError(t, err) {
		assert.Equal(t, "f542b02c-5318-44a4-9c16-c54375f32985", id)
		return
	}
	fmt.Println("deleted student", id)
}

func TestUpdateActivity(t *testing.T) {
	studentRepo := NewStudent(db)

	reqStudent := models.Activity{
		Id:       "dc0f279e-ff3e-4a82-af21-9032a98c2b01",
		IsActive: true}

	_, err := studentRepo.UpdateActivity(context.Background(), reqStudent)
	if assert.NoError(t, err) {
		createdStudent, err := studentRepo.GetStudentById(context.Background(), "35")
		if assert.NoError(t, err) {
			assert.Equal(t, reqStudent.Id, createdStudent.Id)
		} else {
			return
		}
		fmt.Println("Created student", createdStudent)
	}
}
