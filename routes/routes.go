package routes

import (
	"github.com/gin-gonic/gin"

	"go-gin-api/handlers"
)

// ฟังก์ชันตั้งค่า Routes
func SetupRouter(userHandler *handlers.UserHandler, helloHandler *handlers.HelloHandler) *gin.Engine {
	r := gin.Default()

	// เพิ่ม Route สำหรับ "/users"
	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:email", userHandler.UpdateUser)

	// เพิ่ม Route สำหรับ "/"
	r.GET("/", helloHandler.HelloHandler)

	// สามารถเพิ่ม Route อื่นๆ ได้ที่นี่ เช่น เพิ่มการจัดการกับ `POST /users`
	// r.POST("/users", userHandler.CreateUser)

	return r
}
