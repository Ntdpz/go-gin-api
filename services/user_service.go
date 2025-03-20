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

// ฟังก์ชันสำหรับดึงข้อมูลผู้ใช้ตาม Email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.Repo.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ฟังก์ชันสำหรับสร้างผู้ใช้ใหม่
func (s *UserService) CreateUser(userInput models.UserInput) (models.User, error) {
	user := models.User{
		Name:  userInput.Name,
		Email: userInput.Email,
	}
	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(email string, userInput models.User) (*models.User, error) {
	// ค้นหาผู้ใช้จากฐานข้อมูล
	var user models.User
	if err := s.Repo.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// อัปเดตเฉพาะฟิลด์ที่มีการส่งมาจาก client
	if userInput.Name != "" {
		user.Name = userInput.Name
	}

	// ตรวจสอบว่า email ที่ส่งมามีค่า ถ้าไม่มีให้ใช้ email เดิม
	if userInput.Email != "" {
		user.Email = userInput.Email
	} else {
		user.Email = user.Email // ใช้ email เดิม
	}

	// บันทึกการเปลี่ยนแปลง
	if err := s.Repo.Save(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// ฟังก์ชันลบผู้ใช้ตาม Email
func (s *UserService) DeleteUser(email string) error {
	var user models.User
	// ค้นหาผู้ใช้จากฐานข้อมูลตาม Email
	if err := s.Repo.Where("email = ?", email).First(&user).Error; err != nil {
		return err // หากไม่พบผู้ใช้
	}

	// ลบผู้ใช้
	if err := s.Repo.Delete(&user); err != nil {
		return err // หากเกิดข้อผิดพลาดในการลบ
	}

	return nil
}
