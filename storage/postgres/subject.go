package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type subjectRepo struct {
	db *pgxpool.Pool
}

func NewSubject(db *pgxpool.Pool) subjectRepo {
	return subjectRepo{
		db: db,
	}
}

func (s *subjectRepo) Create(ctx context.Context, subject models.Subject) (string, error) {

	id := uuid.New()

	query := ` INSERT INTO subjects (
		id,
		name,
		type,
		created_at) VALUES ($1, $2, $3, 'NOW()') `

	_, err := s.db.Exec(ctx,
		query,
		id,
		subject.Name,
		subject.Type)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *subjectRepo) Update(ctx context.Context, subject models.Subject, id string) (string, error) {

	query := `UPDATE subjects SET name = $1, 
	type =$2, 
	updated = 'NOW()' WHERE id = $3`

	_, err := s.db.Exec(ctx, query,
		subject.Name,
		subject.Type,
		id)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (s *subjectRepo) GetSubjectById(ctx context.Context, id string) (models.GetSubject, error) {

	var (
		updated sql.NullString
		subject models.GetSubject
	)
	query := `SELECT id,
					name,
					type,
					to_char(created_at, 'YYYY-MM-DD HH:MM:SS AM'),
					to_char(updated, 'YYYY-MM-DD HH:MM:SS AM')
				FROM subjects
				WHERE id = $1 LIMIT 1`
	rows := s.db.QueryRow(ctx, query, id)

	err := rows.Scan(
		&subject.Id,
		&subject.Name,
		&subject.Type,
		&subject.Created_at,
		&updated)
	if err != nil {
		return subject, err
	}
	subject.Updated = pkg.NullStringToString(updated)

	return subject, nil
}

func (s *subjectRepo) GetAll(ctx context.Context, req models.GetAllSubjectsRequest) (models.GetAllSubjectsResponse, error) {
	resp := models.GetAllSubjectsResponse{}
	var updated sql.NullString
	offest := (req.Page - 1) * req.Limit
	query := `SELECT id,
					name,
					type,
					to_char(created_at, 'YYYY-MM-DD HH:MM:SS AM'),
					to_char(updated, 'YYYY-MM-DD HH:MM:SS AM')
				FROM subjects
				OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(ctx, query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var subject models.GetSubject
		if err := rows.Scan(
			&subject.Id,
			&subject.Name,
			&subject.Type,
			&subject.Created_at,
			&updated); err != nil {
			return resp, err
		}

		subject.Updated = pkg.NullStringToString(updated)

		resp.Subjects = append(resp.Subjects, subject)
	}

	err = s.db.QueryRow(ctx, `SELECT count(*) from subjects`).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *subjectRepo) Delete(ctx context.Context, id string) (string, error) {
	query := `DELETE FROM subjects WHERE id = $1`

	_, err := s.db.Exec(ctx, query, id)

	if err != nil {
		return "", err
	}
	return id, nil
}
