package grade

import (
	"errors"
)

type GradeService struct {
	Repo Repository
}

type Service interface {
	SubmitGrade(req Request) (*Response, error)
	CheckGrade(studentID string) (*Response, error)
}

func (s *GradeService) SubmitGrade(req Request) (*Response, error) {
	total, gradeLetter := CalculateGrade(req.Homework, req.Midterm, req.Final)

	res := Response{
		StudentID: req.StudentID,
		Total:     total,
		Grade:     gradeLetter,
	}

	err := s.Repo.InsertGrade(res, req.Homework, req.Midterm, req.Final)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *GradeService) CheckGrade(studentID string) (*Response, error) {
	if studentID == "" {
		return nil, errors.New("student ID is required")
	}

	grade, err := s.Repo.GetGradeByStudentID(studentID)
	if err != nil {
		return nil, err
	}

	return grade, nil
}

func CalculateGrade(homework, midterm, final float64) (float64, string) {
	if homework < 0 || midterm < 0 || final < 0 {
		return 0, "Invalid"
	}

	if homework > 100 || midterm > 100 || final > 100 {
		return 0, "Invalid"
	}

	total := homework*0.3 + midterm*0.3 + final*0.4

	var grade string
	if total >= 80 {
		grade = "A"
	} else if total >= 70 {
		grade = "B"
	} else if total >= 60 {
		grade = "C"
	} else if total >= 50 {
		grade = "D"
	} else {
		grade = "F"
	}

	return total, grade
}
