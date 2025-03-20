package models

import "gorm.io/gorm"

// Struct User ใช้กับ GORM
type User struct {
	gorm.Model        // เพิ่ม Model นี้เพื่อให้มี ID, CreatedAt, UpdatedAt, DeletedAt อัตโนมัติ
	Name       string `json:"name"`
	Email      string `json:"email"`
}
