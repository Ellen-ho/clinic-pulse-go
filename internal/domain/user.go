package domain

import "gorm.io/gorm"

type User struct {
    gorm.Model
    FirstName string `gorm:"size:255"`
    LastName  string `gorm:"size:255"`
    Email     string `gorm:"uniqueIndex;size:255"`
}