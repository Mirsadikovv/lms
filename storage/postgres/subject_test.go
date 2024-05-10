package postgres

import (
	"backend_course/lms/api/models"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateSubject(t *testing.T) {
	subjectRepo := NewSubject(db)
	id := uuid.New()
	reqSubject := models.Subject{
		Id:   id.String(),
		Name: faker.Word(),
		Type: faker.Word()}

	_, err := subjectRepo.Create(reqSubject)
	if assert.NoError(t, err) {
		createdSubject, err := subjectRepo.GetSubjectById(id.String())
		if assert.NoError(t, err) {
			assert.Equal(t, reqSubject.Name, createdSubject.Name)
			assert.Equal(t, reqSubject.Type, createdSubject.Type)
		} else {
			return
		}
		fmt.Println("Created subject", createdSubject)
	}

}

func TestUpdateSubject(t *testing.T) {
	subjectRepo := NewSubject(db)

	reqSubject := models.Subject{
		Name: faker.Word(),
		Type: faker.Word()}

	_, err := subjectRepo.Update(reqSubject, "0e074d6b-9c4d-4a98-b5ff-ab9daa929d5c")
	if assert.NoError(t, err) {
		createdSubject, err := subjectRepo.GetSubjectById("0e074d6b-9c4d-4a98-b5ff-ab9daa929d5c")
		if assert.NoError(t, err) {
			assert.Equal(t, reqSubject.Name, createdSubject.Name)
			assert.Equal(t, reqSubject.Type, createdSubject.Type)
		} else {
			return
		}
		fmt.Println("Created subject", createdSubject)
	}
}

func TestGetSubject(t *testing.T) {
	subjectRepo := NewSubject(db)

	reqSubject := models.Subject{
		Name: faker.Word(),
		Type: faker.Word()}

	_, err := subjectRepo.Update(reqSubject, "0e074d6b-9c4d-4a98-b5ff-ab9daa929d5c")
	if assert.NoError(t, err) {
		createdSubject, err := subjectRepo.GetSubjectById("0e074d6b-9c4d-4a98-b5ff-ab9daa929d5c")
		if assert.NoError(t, err) {
			assert.Equal(t, reqSubject.Name, createdSubject.Name)
			assert.Equal(t, reqSubject.Type, createdSubject.Type)
		} else {
			return
		}
		fmt.Println("Created subject", createdSubject)
	}
}

func TestDeleteSubject(t *testing.T) {
	subjectRepo := NewSubject(db)

	id, err := subjectRepo.Delete("d20ffed0-e186-4512-86aa-4575bdc51fbb")
	if assert.NoError(t, err) {
		assert.Equal(t, "d20ffed0-e186-4512-86aa-4575bdc51fbb", id)
		return
	}
	fmt.Println("deleted subject", id)
}
