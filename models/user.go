package models

import (
	"errors"
	"strings"
	"time"
)

var (
	Err_UsernameCannotBeEmpty = errors.New("Username cannot be empty")
	Err_PasswordRange         = errors.New("Password length must be great than 0 and less than 5 characters")
)

type User struct {
	Id       int    `json:"id,omitempty" gorm:"column:id"`
	UserName string `json:"userName" gorm:"column:user_name"`
	Email    string `json:"email" gorm:"column:email"`
	//FirstName string `json:"firstName" gorm: "column:firstName"`
	//LastName  string `json:"lastName" gorm: "column:lastName"`
	EncryptedPassword string    `json:"encryptedPassword" gorm:"column:encrypted_password"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy         string    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy         string    `json:"updated_by" gorm:"column:updated_by"`
	DeleteFlag        bool      `json:"deleted_flag" gorm:"column:deleted_flag"`
}

// vì ng dùng ko dùng hết cái user đó nên có 1 struct DTO để lấy 1 số trường user nhập
// Model này dành cho user nhập liệu (DTO)
type UserCreation struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type UserUpdateParams struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func (u *UserCreation) Validate() error {
	u.UserName = strings.TrimSpace(u.UserName)
	if u.UserName == "" {
		return Err_UsernameCannotBeEmpty
	}

	u.Password = strings.TrimSpace(u.Password)

	if len(u.Password) <= 0 || len(u.Password) > 5 {
		return Err_PasswordRange
	}

	return nil
}
