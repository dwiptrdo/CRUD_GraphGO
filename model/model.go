package model

import (
	"time"
)

// type User struct {
//     ID       uint   `gorm:"primary_key"`
//     Username string `gorm:"type:varchar(100);unique_index"`
//     Email    string `gorm:"type:varchar(100);unique_index"`
// }

type User struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	Nama       string    `json:"nama"`
	Username   string    `json:"username"`
	Kata_Sandi string    `json:"kata_sandi"`
	Email      string    `json:"email"`
	Status     string    `json:"status" gorm:"default:'active'"`
}
