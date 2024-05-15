package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"backend_course/lms/pkg/hash"
	"context"
	"database/sql"
	"fmt"

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

func (s *studentRepo) Create(ctx context.Context, student models.Student) (string, error) {

	id := uuid.New()
	pasword, err := hash.HashPassword(student.Pasword)
	if err != nil {
		return "", err
	}

	query := ` INSERT INTO students (
		id,
		first_name,
		last_name,
		age,
		external_id,
		phone,
		mail,
		pasword) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) `

	_, err = s.db.Exec(ctx,
		query,
		id,
		student.FirstName,
		student.LastName,
		student.Age,
		student.External_id,
		student.Phone,
		student.Mail,
		pasword)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *studentRepo) Update(ctx context.Context, student models.UpdateStudent, id string) (string, error) {
	pasword, err := hash.HashPassword(student.Pasword)
	if err != nil {
		return "", err
	}

	query := `UPDATE students SET first_name = $1, 
	last_name =$2, 
	age = $3, 
	external_id = $4,
	phone = $5, 
	mail = $6,
	pasword = $7,
	updated = 'NOW()' WHERE id = $8`

	_, err = s.db.Exec(ctx, query,
		student.FirstName,
		student.LastName,
		student.Age,
		student.External_id,
		student.Phone,
		student.Mail,
		pasword,
		id)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (s *studentRepo) GetAll(ctx context.Context, req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
	resp := models.GetAllStudentsResponse{}
	var created sql.NullString
	filter := ""
	offest := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT id,
					first_name,
					last_name,
					age,
					external_id,
					to_char(created_at, 'YYYY-MM-DD HH:MM:SS AM')
				FROM students
				WHERE TRUE ` + filter + `
				OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(ctx, query, offest, req.Limit)
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
			&student.Age,
			&student.External_id,
			&created); err != nil {
			return resp, err
		}

		student.LastName = pkg.NullStringToString(lastName)
		student.Created_at = pkg.NullStringToString(created)

		resp.Students = append(resp.Students, student)
	}

	err = s.db.QueryRow(ctx, `SELECT count(*) from students WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *studentRepo) GetStudentById(ctx context.Context, external_id string) (models.GetStudent, error) {
	student := models.GetStudent{}
	var updated sql.NullString
	query := `SELECT id,
					first_name,
					last_name,
					age,
					external_id,
					phone,
					mail,
					to_char(created_at, 'YYYY-MM-DD HH:MM:SS AM'),
					to_char(updated, 'YYYY-MM-DD HH:MM:SS AM')
				FROM students
				WHERE external_id = $1 LIMIT 1`
	rows := s.db.QueryRow(ctx, query, external_id)
	fmt.Println(external_id)
	err := rows.Scan(
		&student.Id,
		&student.FirstName,
		&student.LastName,
		&student.Age,
		&student.External_id,
		&student.Phone,
		&student.Mail,
		&student.Created_at,
		&updated)
	if err != nil {
		return student, err
	}
	student.Updated = pkg.NullStringToString(updated)

	return student, nil
}

func (s *studentRepo) Delete(ctx context.Context, id string) (string, error) {
	query := `DELETE FROM students WHERE id = $1`

	_, err := s.db.Exec(ctx, query, id)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *studentRepo) UpdateActivity(ctx context.Context, student models.Activity) (string, error) {
	query := "UPDATE students SET isactive = $1 WHERE id = $2"
	_, err := s.db.Exec(ctx, query, student.IsActive, student.Id)
	if err != nil {
		return "", err
	}

	return student.Id, nil
}

func (s *studentRepo) CheckLessonNow(ctx context.Context, id string) (models.StudentLessonNow, error) {
	lessonInfo := models.StudentLessonNow{}
	var timeUntilEnd sql.NullString
	query := `
	SELECT 	sub.name AS subject,
			t.first_name AS teacher,
			tt.to_date - NOW()
		FROM 
			time_table tt
		JOIN 
			teachers t ON tt.teacher_id = t.id
		JOIN 
			students s ON tt.student_id = s.id
		JOIN 
			subjects sub ON tt.subject_id = sub.id 
		WHERE s.id = $1 AND from_date <= NOW() AND to_date >= NOW() LIMIT 1`

	rows := s.db.QueryRow(ctx, query, id)
	fmt.Println(id)
	err := rows.Scan(
		&lessonInfo.SubjectName,
		&lessonInfo.TeacherName,
		&timeUntilEnd)

	if err != nil {
		return lessonInfo, err
	}
	lessonInfo.TimeUntilEnd = pkg.NullStringToString(timeUntilEnd)

	return lessonInfo, nil
}
