package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type studentRepo struct {
	db *pgxpool.Pool
}

func NewStudent(db *pgxpool.Pool) studentRepo {
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

	_, err := s.db.Exec(context.Background(),
		query,
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

	_, err := s.db.Exec(context.Background(), query, student.FirstName,
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
					last_name,
					external_id
				FROM students
				WHERE TRUE ` + filter + `
				OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(context.Background(), query, offest, req.Limit)
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
			&lastName,
			&student.External_id); err != nil {
			return resp, err
		}

		student.LastName = pkg.NullStringToString(lastName)
		resp.Students = append(resp.Students, student)
	}

	err = s.db.QueryRow(context.Background(), `SELECT count(*) from students WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *studentRepo) GetStudentById(student models.GetStudent) (models.GetStudent, error) {
	query := `SELECT id,
					first_name,
					last_name,
					external_id
				FROM students
				WHERE external_id = $1 LIMIT 1`
	rows := s.db.QueryRow(context.Background(), query, student.External_id)

	err := rows.Scan(
		&student.Id,
		&student.FirstName,
		&student.LastName,
		&student.External_id)
	if err != nil {
		return student, err
	}

	return student, nil
}

func (s *studentRepo) Delete(id string) error {
	query := `DELETE FROM students WHERE id = $1`

	_, err := s.db.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepo) UpdateActivity(student models.Activity) (string, error) {
	query := "UPDATE students SET isactive = $1 WHERE id = $2"
	_, err := s.db.Exec(context.Background(), query, student.IsActive, student.Id)
	if err != nil {
		return "", err
	}

	return student.Id, nil
}
