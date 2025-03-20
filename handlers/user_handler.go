package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin-api/models"
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

// ฟังก์ชันสำหรับสร้างผู้ใช้ใหม่
func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser models.UserInput // ใช้ Model UserInput เพื่อตรวจสอบข้อมูล

	// เช็คว่าข้อมูลที่รับเข้ามาถูกต้องหรือไม่
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// เรียกใช้ Service เพื่อบันทึกข้อมูล
	user, err := h.Service.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	// ดึง email จาก URL Parameter
	email := c.Param("email")

	// ค้นหาผู้ใช้จากฐานข้อมูล
	existingUser, err := h.Service.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// สร้างตัวแปรสำหรับรับข้อมูลที่ต้องการอัปเดต
	var userInput models.User

	// Bind ข้อมูลที่ส่งมาจาก Request Body
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// ถ้าไม่ส่งค่า name หรือ email มา, ให้ใช้ข้อมูลเดิม
	if userInput.Name == "" {
		userInput.Name = existingUser.Name
	}
	if userInput.Email == "" {
		userInput.Email = existingUser.Email
	}

	// เรียกใช้ Service ในการอัปเดตข้อมูลผู้ใช้
	updatedUser, err := h.Service.UpdateUser(email, userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update user"})
		return
	}

	// ส่งข้อมูลผู้ใช้ที่ถูกอัปเดต
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}

// ฟังก์ชันสำหรับลบผู้ใช้
func (h *UserHandler) DeleteUser(c *gin.Context) {
	// ดึง email จาก URL Parameter
	email := c.Param("email")

	// เรียกใช้ Service ในการลบผู้ใช้
	err := h.Service.DeleteUser(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// ส่งผลลัพธ์ว่าได้ลบข้อมูลผู้ใช้เรียบร้อย
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
