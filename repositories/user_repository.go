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

// ฟังก์ชันสร้างผู้ใช้ใหม่
func (r *UserRepository) CreateUser(user models.User) (models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

// ฟังก์ชัน FindByEmail เพื่อค้นหาผู้ใช้จาก Email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ฟังก์ชันเพื่อค้นหาผู้ใช้จากฐานข้อมูลโดยใช้เงื่อนไข
func (r *UserRepository) Where(query string, args ...interface{}) *gorm.DB {
	return r.DB.Where(query, args...)
}

// ฟังก์ชันเพื่อดึงข้อมูลผู้ใช้คนแรกจากฐานข้อมูล
func (r *UserRepository) First(user *models.User) error {
	return r.DB.First(user).Error
}

// ฟังก์ชันสำหรับบันทึกข้อมูลผู้ใช้
func (r *UserRepository) Save(user *models.User) error {
	// ใช้ Save โดยไม่ต้องรับค่าผลลัพธ์สองตัว
	if err := r.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}
