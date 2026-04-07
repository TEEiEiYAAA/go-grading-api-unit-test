package grade

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetGradeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockService{}
	handler := NewHandler(mockService)

	router := gin.Default()
	router.GET("/grade/:studentId", handler.GetGradeHandler)

	req, _ := http.NewRequest("GET", "/grade/65001", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
