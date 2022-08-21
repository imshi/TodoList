package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string `gorm:"password"` // 密文存储
}
