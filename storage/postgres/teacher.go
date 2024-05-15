package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type teacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacher(db *pgxpool.Pool) teacherRepo {
	return teacherRepo{
		db: db,
	}
}

func (s *teacherRepo) Create(ctx context.Context, teacher models.Teacher) (string, error) {

	id := uuid.New()

	query := `INSERT INTO teachers (
		id,
		first_name,
		last_name,
		subject_id,
		start_working,
		phone,
		mail,
		pasword) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) `

	_, err := s.db.Exec(ctx, query,
		id,
		teacher.FirstName,
		teacher.LastName,
		teacher.Subject_id,
		teacher.Start_working,
		teacher.Phone,
		teacher.Mail,
		teacher.Password)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *teacherRepo) Update(ctx context.Context, teacher models.Teacher, id string) (string, error) {

	query := `UPDATE teachers SET 
	first_name = $1, 
	last_name = $2, 
	subject_id = $3, 
	start_working = $4,
	phone = $5, 
	mail = $6 WHERE id = $7`

	_, err := s.db.Exec(ctx, query,
		teacher.FirstName,
		teacher.LastName,
		teacher.Subject_id,
		teacher.Start_working,
		teacher.Phone,
		teacher.Mail,
		id)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (s *teacherRepo) GetAll(ctx context.Context, req models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {
	resp := models.GetAllTeachersResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT id,
					first_name,
					last_name,
					subject_id
				FROM teachers
				WHERE TRUE ` + filter + `
				OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(ctx, query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			teacher  models.GetTeacher
			lastName sql.NullString
		)
		if err := rows.Scan(
			&teacher.Id,
			&teacher.FirstName,
			&lastName,
			&teacher.Subject_id); err != nil {
			return resp, err
		}

		teacher.LastName = pkg.NullStringToString(lastName)
		resp.Teachers = append(resp.Teachers, teacher)
	}

	err = s.db.QueryRow(ctx, `SELECT count(*) from teachers WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *teacherRepo) GetTeacherById(ctx context.Context, id string) (models.GetTeacher, error) {
	teacher := models.GetTeacher{}
	var start_work sql.NullString

	query := `SELECT id,
					first_name,
					last_name,
					subject_id,
					to_char(start_working, 'YYYY-MM-DD'),
					phone,
					mail
				FROM teachers
				WHERE id = $1 LIMIT 1`
	rows := s.db.QueryRow(ctx, query, id)

	err := rows.Scan(
		&teacher.Id,
		&teacher.FirstName,
		&teacher.LastName,
		&teacher.Subject_id,
		&start_work,
		&teacher.Phone,
		&teacher.Mail)

	if err != nil {
		return teacher, err
	}
	teacher.Start_working = pkg.NullStringToString(start_work)

	return teacher, nil
}

func (s *teacherRepo) Delete(ctx context.Context, id string) (string, error) {
	query := `DELETE FROM teachers WHERE id = $1`

	_, err := s.db.Exec(ctx, query, id)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *teacherRepo) CheckLessonNow(ctx context.Context, id string) (models.TeacherLessonNow, error) {
	lessonInfo := models.TeacherLessonNow{}
	teacherInfo := models.TeacherInfo{}
	var timeUntilEnd sql.NullString
	query := `
	SELECT 	sub.name AS subject,
			t.first_name,
			t.last_name,
			t.phone,
			NOW() - tt.to_date
		FROM 
			time_table tt
		JOIN 
			teachers t ON tt.teacher_id = t.id
		JOIN 
			students s ON tt.student_id = s.id
		JOIN 
			subjects sub ON tt.subject_id = sub.id 
		WHERE t.id = $1 AND from_date <= NOW() AND to_date >= NOW() LIMIT 1`

	rows := s.db.QueryRow(ctx, query, id)
	fmt.Println(id)
	err := rows.Scan(
		&lessonInfo.SubjectName,
		&teacherInfo.FirstName,
		&teacherInfo.LastName,
		&teacherInfo.Phone,
		&timeUntilEnd)
	if err != nil {
		return lessonInfo, err
	}
	lessonInfo.Teacher = teacherInfo
	lessonInfo.TimeUntilEnd = pkg.NullStringToString(timeUntilEnd)

	return lessonInfo, nil
}

func (s *teacherRepo) GetTeacherByLogin(ctx context.Context, login string) (models.Teacher, error) {

	query := `
	SELECT
		id,
		first_name,
		last_name,
		subject_id,
		TO_CHAR(start_working,'YYYY-MM-DD HH:MM:SS'),
		phone,
		mail,
		TO_CHAR(created_at,'YYYY-MM-DD HH:MM:SS'),
		pasword
	FROM
		teachers
	WHERE
		mail = $1;
`
	row := s.db.QueryRow(ctx, query, login)

	var (
		teacher models.Teacher
		firstName,
		lastName,
		subjectId,
		startWorking,
		phone,
		mail,
		createdAt sql.NullString
	)

	err := row.Scan(
		&teacher.Id,
		&firstName,
		&lastName,
		&subjectId,
		&startWorking,
		&phone,
		&mail,
		&createdAt,
		&teacher.Password,
	)

	teacher.FirstName = pkg.NullStringToString(firstName)
	teacher.LastName = pkg.NullStringToString(lastName)
	teacher.Subject_id = pkg.NullStringToString(subjectId)
	teacher.Start_working = pkg.NullStringToString(startWorking)
	teacher.Phone = pkg.NullStringToString(phone)
	teacher.Mail = pkg.NullStringToString(mail)

	if err != nil {
		return teacher, err
	}
	return teacher, nil
}
