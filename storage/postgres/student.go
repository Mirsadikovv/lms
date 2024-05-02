package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"database/sql"

	"github.com/google/uuid"
)

type studentRepo struct {
	db *sql.DB
}

func NewStudent(db *sql.DB) studentRepo {
	return studentRepo{
		db: db,
	}
}

func (s *studentRepo) Create(student models.Student) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO students (
		id,
		first_name,
		last_name,
		age,
		external_id,
		phone,
		mail,
		pasword) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) `

	_, err := s.db.Exec(query,
		id,
		student.FirstName,
		student.LastName,
		student.Age,
		student.External_id,
		student.Phone,
		student.Mail,
		student.Pasword)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *studentRepo) Update(student models.Student) (string, error) {

	query := `UPDATE students SET first_name = $1, 
	last_name =$2, 
	age = $3, 
	external_id = $4,
	phone = $5, 
	mail = $6,
	pasword = $7 WHERE id = $8`

	_, err := s.db.Exec(query, student.FirstName,
		student.LastName,
		student.Age,
		student.External_id,
		student.Phone,
		student.Mail,
		student.Pasword,
		student.Id)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (s *studentRepo) GetAll(req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
	resp := models.GetAllStudentsResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT id,
					first_name,
					last_name
				FROM students
				WHERE TRUE ` + filter + `
				OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			student  models.GetStudent
			lastName sql.NullString
		)
		if err := rows.Scan(
			&student.Id,
			&student.FirstName,
			&lastName); err != nil {
			return resp, err
		}

		student.LastName = pkg.NullStringToString(lastName)
		resp.Students = append(resp.Students, student)
	}

	err = s.db.QueryRow(`SELECT count(*) from students WHERE TRUE ` + filter + ``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *studentRepo) Delete(student models.Student) error {
	query := `DELETE FROM students WHERE id = $1`

	_, err := s.db.Exec(query, student.Id)

	if err != nil {
		return err
	}
	return nil
}
