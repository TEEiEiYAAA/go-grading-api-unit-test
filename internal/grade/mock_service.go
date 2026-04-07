package grade

type MockService struct{}

func (m *MockService) CheckGrade(studentID string) (*Response, error) {
	return &Response{
		StudentID: studentID,
		Total:     90,
		Grade:     "A",
	}, nil
}

func (m *MockService) SubmitGrade(req Request) (*Response, error) {
	return &Response{
		StudentID: req.StudentID,
		Total:     85,
		Grade:     "A",
	}, nil
}
