package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloHandler - สำหรับจัดการ API สำหรับข้อความ "Hello GoGin"
type HelloHandler struct{}

// ฟังก์ชันสำหรับแสดงข้อความ "Hello GoGin"
func (h *HelloHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello GoGin",
	})
}
