package grade

import (
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
		{"Grade B", 80, 70, 60, "B"},
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
