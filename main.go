package main

import (
	"go-gin-api/config"
	"go-gin-api/handlers"
	"go-gin-api/models"
	"go-gin-api/repositories"
	"go-gin-api/routes"
	"go-gin-api/services"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDatabase()

	// Auto Migrate Database (สร้างตารางอัตโนมัติ)
	config.DB.AutoMigrate(&models.User{})

	// สร้าง instance ของ repository, service และ handler สำหรับ User
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// สร้าง instance ของ handler สำหรับ Hello
	helloHandler := &handlers.HelloHandler{}

	// ตั้งค่า Router
	r := routes.SetupRouter(userHandler, helloHandler)

	// รันเซิร์ฟเวอร์
	r.Run(":8080")
}
