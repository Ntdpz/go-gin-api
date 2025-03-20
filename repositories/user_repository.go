package repositories

import (
	"gorm.io/gorm"

	"go-gin-api/models"
)

// Struct UserRepository ใช้ GORM
type UserRepository struct {
	DB *gorm.DB
}

// ฟังก์ชันสร้าง instance ของ UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// ฟังก์ชันดึงข้อมูลผู้ใช้ทั้งหมดจาก Database
func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := repo.DB.Find(&users) // ใช้ GORM Query แทน SQL ดั้งเดิม
	return users, result.Error
}
