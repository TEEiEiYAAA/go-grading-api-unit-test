package grade

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) SubmitGradeHandler(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	res, err := h.Service.SubmitGrade(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to save grade"})
		return
	}

	c.JSON(200, res)
}

func (h *Handler) GetGradeHandler(c *gin.Context) {
	studentID := c.Param("studentId")

	grade, err := h.Service.CheckGrade(studentID)
	if err != nil {
		c.JSON(404, gin.H{"error": "grade not found"})
		return
	}

	c.JSON(200, grade)
}
