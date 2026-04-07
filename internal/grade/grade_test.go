package grade

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------------- ทดสอบเกรด A ----------------
func TestCalculateGradeA(t *testing.T) {
	_, grade := CalculateGrade(80, 70, 90)
	assert.Equal(t, "A", grade)
}

// ---------------- ทดสอบเกรด B ----------------
func TestCalculateGradeB(t *testing.T) {
	_, grade := CalculateGrade(80, 70, 75)
	assert.Equal(t, "B", grade)
}

// ---------------- ทดสอบเกรด C ----------------
func TestCalculateGradeC(t *testing.T) {
	_, grade := CalculateGrade(70, 60, 65)
	assert.Equal(t, "C", grade)
}

// ---------------- ทดสอบเกรด D ----------------
func TestCalculateGradeD(t *testing.T) {
	_, grade := CalculateGrade(60, 50, 55)
	assert.Equal(t, "D", grade)
}

// ---------------- ทดสอบเกรด F ----------------
func TestCalculateGradeF(t *testing.T) {
	_, grade := CalculateGrade(40, 40, 40)
	assert.Equal(t, "F", grade)
}

// ---------------- Invalid ----------------
func TestInvalidScore(t *testing.T) {
	_, grade := CalculateGrade(-10, 20, 30)
	assert.Equal(t, "Invalid", grade)
}

// ---------------- Boundary ----------------
func TestBoundaryScore(t *testing.T) {
	_, grade := CalculateGrade(999, 20, 30)
	assert.Equal(t, "Invalid", grade)
}

// ---------------- Task 4 ------------------
func TestCalculateGrade_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		homework float64
		midterm  float64
		final    float64
		expected string
	}{
		{"Grade A", 80, 70, 90, "A"},
		{"Grade B", 80, 70, 75, "B"},
		{"Grade C", 70, 60, 65, "C"},
		{"Grade D", 60, 50, 55, "D"},
		{"Grade F", 40, 40, 40, "F"},
		{"Invalid More than 100", 150, 100, 100, "Invalid"},
		{"Invalid Negative Number", -10, 50, 50, "Invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, grade := CalculateGrade(tt.homework, tt.midterm, tt.final)
			assert.Equal(t, tt.expected, grade)
		})
	}
}

// --------------- Task 5 ----------------------
type MockRepository struct{}

func (m *MockRepository) GetGradeByStudentID(studentID string) (*Response, error) {
	if studentID == "error_case" {
		return nil, errors.New("mock db error")
	}
	return &Response{StudentID: studentID, Total: 90, Grade: "A"}, nil
}

func (m *MockRepository) InsertGrade(g Response, homework, midterm, final float64) error {
	if g.StudentID == "error_case" {
		return errors.New("mock insert error")
	}
	return nil
}

func TestCheckGrade(t *testing.T) {
	mockRepo := &MockRepository{}
	service := &GradeService{Repo: mockRepo}

	t.Run("Success", func(t *testing.T) {
		res, err := service.CheckGrade("65001")
		assert.NoError(t, err)
		assert.Equal(t, "65001", res.StudentID)
		assert.Equal(t, "A", res.Grade)
	})

	t.Run("Empty Student ID", func(t *testing.T) {
		res, err := service.CheckGrade("")
		assert.Error(t, err)
		assert.Equal(t, "student ID is required", err.Error())
		assert.Nil(t, res)
	})

	t.Run("Database Error", func(t *testing.T) {
		res, err := service.CheckGrade("error_case")
		assert.Error(t, err)
		assert.Equal(t, "mock db error", err.Error())
		assert.Nil(t, res)
	})
}

func TestSubmitGrade(t *testing.T) {
	mockRepo := &MockRepository{}
	service := &GradeService{Repo: mockRepo}

	t.Run("Success", func(t *testing.T) {
		req := Request{StudentID: "65001", Homework: 80, Midterm: 70, Final: 90}
		res, err := service.SubmitGrade(req)

		assert.NoError(t, err)
		assert.Equal(t, "65001", res.StudentID)
		assert.Equal(t, "A", res.Grade)
	})

	t.Run("Insert Error", func(t *testing.T) {
		req := Request{StudentID: "error_case", Homework: 80, Midterm: 70, Final: 90}
		res, err := service.SubmitGrade(req)

		assert.Error(t, err)
		assert.Equal(t, "mock insert error", err.Error())
		assert.Nil(t, res)
	})
}
