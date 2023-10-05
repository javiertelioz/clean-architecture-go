package model

import (
	"database/sql"
	"fmt"
	"github.com/javiertelioz/clean-architecture-go/config"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name        string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Surname     string `gorm:"not null"`
	Phone       string `gorm:"unique;not null"`
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	ActivatedAt sql.NullTime
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time `gorm:"index"`
}

func (u *User) TableName() string {
	schema, _ := config.GetConfig[string]("Database.schema")
	return fmt.Sprintf("%s.users", schema)
}
