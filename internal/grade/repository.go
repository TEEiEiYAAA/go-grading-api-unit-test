package grade

import (
	"go-grading-api/config"
)

type GradeRepository struct{}

func NewGradeService(repo Repository) *GradeService {
	return &GradeService{Repo: repo}
}

type Repository interface {
	InsertGrade(g Response, homework, midterm, final float64) error
	GetGradeByStudentID(studentID string) (*Response, error)
}

func (r *GradeRepository) InsertGrade(g Response, homework, midterm, final float64) error {
	query := `
	INSERT INTO grades (student_id, homework, midterm, final, total, grade)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := config.DB.Exec(
		query,
		g.StudentID,
		homework,
		midterm,
		final,
		g.Total,
		g.Grade,
	)

	return err
}

func (r *GradeRepository) GetGradeByStudentID(studentID string) (*Response, error) {
	query := `
	SELECT student_id, total, grade
	FROM grades
	WHERE student_id = ?
	ORDER BY id DESC
	LIMIT 1
	`

	row := config.DB.QueryRow(query, studentID)

	var res Response
	err := row.Scan(&res.StudentID, &res.Total, &res.Grade)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
