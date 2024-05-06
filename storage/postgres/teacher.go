package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"

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

func (s *teacherRepo) Create(teacher models.Teacher) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO teachers (
		id,
		first_name,
		last_name,
		subject_id,
		start_working,
		phone,
		mail) VALUES ($1, $2, $3, $4, $5, $6, $7) `

	_, err := s.db.Exec(context.Background(), query,
		id,
		teacher.FirstName,
		teacher.LastName,
		teacher.Subject_id,
		teacher.Start_working,
		teacher.Phone,
		teacher.Mail)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *teacherRepo) Update(teacher models.Teacher) (string, error) {

	query := `UPDATE teachers SET 
	first_name = $1, 
	last_name =$2, 
	subject_id = $3, 
	start_working = $4,
	phone = $5, 
	mail = $6 WHERE id = $7`

	_, err := s.db.Exec(context.Background(), query,
		teacher.FirstName,
		teacher.LastName,
		teacher.Subject_id,
		teacher.Start_working,
		teacher.Phone,
		teacher.Mail,
		teacher.Id)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (s *teacherRepo) GetAll(req models.GetAllTeachersRequest) (models.GetAllTeachersResponse, error) {
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
	rows, err := s.db.Query(context.Background(), query, offest, req.Limit)
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

	err = s.db.QueryRow(context.Background(), `SELECT count(*) from teachers WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *teacherRepo) GetTeacherById(teacher models.GetTeacher) (models.GetTeacher, error) {
	query := `SELECT id,
					first_name,
					last_name,
					subject_id
				FROM teachers
				WHERE id = $1 LIMIT 1`
	rows := s.db.QueryRow(context.Background(), query, teacher.Id)

	err := rows.Scan(
		&teacher.Id,
		&teacher.FirstName,
		&teacher.LastName,
		&teacher.Subject_id)
	if err != nil {
		return teacher, err
	}

	return teacher, nil
}

func (s *teacherRepo) Delete(id string) error {
	query := `DELETE FROM teachers WHERE id = $1`

	_, err := s.db.Exec(context.Background(), query, id)

	if err != nil {
		return err
	}
	return nil
}

// func (s *teacherRepo) UpdateActivity( models.Activity) (string, error) {
// 	query := "UPDATE teachers SET isactive = $1 WHERE id = $2"
// 	_, err := s.db.Exec(context.Background(), query, .IsActive, .Id)
// 	if err != nil {
// 		return "", err
// 	}

// 	return .Id, nil
// }
