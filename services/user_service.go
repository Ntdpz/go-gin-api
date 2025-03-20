package services

import (
	"go-gin-api/models"
	"go-gin-api/repositories"
)

// Struct UserService ใช้ GORM
type UserService struct {
	Repo *repositories.UserRepository
}

// ฟังก์ชันสร้าง instance ของ UserService
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// ฟังก์ชันดึงข้อมูลผู้ใช้ทั้งหมด
func (service *UserService) GetUsers() ([]models.User, error) {
	return service.Repo.GetAllUsers()
}
