package models

import "gorm.io/gorm"

// User ใช้กับ GORM สำหรับเก็บข้อมูลในฐานข้อมูล
type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
}

// UserInput ใช้สำหรับตรวจสอบค่าที่รับมาก่อนบันทึกลงฐานข้อมูล
type UserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
