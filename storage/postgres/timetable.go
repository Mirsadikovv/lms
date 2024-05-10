package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type timetableRepo struct {
	db *pgxpool.Pool
}

func NewTimetable(db *pgxpool.Pool) timetableRepo {
	return timetableRepo{
		db: db,
	}
}

func (s *timetableRepo) Create(timetable models.Timetable) (string, error) {

	id := uuid.New()

	query := `INSERT INTO time_table (
		id,
		teacher_id,
		student_id,
		subject_id,
		from_date,
		to_date) 
		VALUES ($1, $2, $3, $4, $5, $6) `

	_, err := s.db.Exec(context.Background(), query,
		timetable.Id,
		timetable.Teacher_id,
		timetable.Student_id,
		timetable.Subject_id,
		timetable.FromDate,
		timetable.ToDate)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *timetableRepo) Update(timetable models.Timetable, id string) (string, error) {

	query := `UPDATE time_table SET 
	teacher_id = $1, 
	student_id = $2, 
	subject_id = $3, 
	from_date = $4,
	to_date = $5, 
	WHERE id = $6`

	_, err := s.db.Exec(context.Background(), query,
		timetable.Teacher_id,
		timetable.Student_id,
		timetable.Subject_id,
		timetable.FromDate,
		timetable.ToDate,
		id)
	if err != nil {
		return "", err
	}

	return "", nil
}

// func (s *timetableRepo) GetAll(req models.GetAllTimetablesRequest) (models.GetAllTimetablesResponse, error) {
// 	resp := models.GetAllTimetablesResponse{}
// 	filter := ""
// 	offest := (req.Page - 1) * req.Limit

// 	if req.Search != "" {
// 		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
// 	}

// 	query := `SELECT id,
// 					first_name,
// 					last_name,
// 					subject_id
// 				FROM timetables
// 				WHERE TRUE ` + filter + `
// 				OFFSET $1 LIMIT $2
// 					`
// 	rows, err := s.db.Query(context.Background(), query, offest, req.Limit)
// 	if err != nil {
// 		return resp, err
// 	}
// 	for rows.Next() {
// 		var (
// 			timetable models.GetTimetable
// 			lastName  sql.NullString
// 		)
// 		if err := rows.Scan(
// 			&timetable.Id,
// 			&timetable.FirstName,
// 			&lastName,
// 			&timetable.Subject_id); err != nil {
// 			return resp, err
// 		}

// 		timetable.LastName = pkg.NullStringToString(lastName)
// 		resp.Timetables = append(resp.Timetables, timetable)
// 	}

// 	err = s.db.QueryRow(context.Background(), `SELECT count(*) from timetables WHERE TRUE `+filter+``).Scan(&resp.Count)
// 	if err != nil {
// 		return resp, err
// 	}

// 	return resp, nil
// }

func (s *timetableRepo) GetTimetableById(id string) (models.GetTimetable, error) {
	timetable := models.GetTimetable{}
	var fromdate, todate sql.NullString

	query := `
	SELECT 	tt.id,
			t.first_name AS teacher,
			s.first_name AS student,
			sub.name AS subject,			
			to_char(tt.from_date, 'YYYY-MM-DD HH:MM:SS AM'),
			to_char(tt.to_date, 'YYYY-MM-DD HH:MM:SS AM')
		FROM 
			time_table tt
		JOIN 
			teachers t ON tt.teacher_id = t.id
		JOIN 
			students s ON tt.student_id = s.id
		JOIN 
			subjects sub ON tt.subject_id = sub.id 
		WHERE tt.id = $1 LIMIT 1`

	rows := s.db.QueryRow(context.Background(), query, id)

	err := rows.Scan(
		&timetable.Id,
		&timetable.Teacher,
		&timetable.Student,
		&timetable.Subject,
		&fromdate,
		&todate)

	if err != nil {
		return timetable, err
	}
	timetable.FromDate = pkg.NullStringToString(fromdate)
	timetable.ToDate = pkg.NullStringToString(todate)

	return timetable, nil
}

func (s *timetableRepo) Delete(id string) (string, error) {
	query := `DELETE FROM time_table WHERE id = $1`

	_, err := s.db.Exec(context.Background(), query, id)

	if err != nil {
		return "", err
	}
	return id, nil
}
