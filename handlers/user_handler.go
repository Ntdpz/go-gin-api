package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin-api/services"
)

// UserHandler struct สำหรับจัดการกับ User
type UserHandler struct {
	Service *services.UserService
}

// ฟังก์ชันสร้าง UserHandler
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// ฟังก์ชันสำหรับดึงข้อมูลผู้ใช้ทั้งหมด
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.Service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(users),
		"data":  users,
	})
}
