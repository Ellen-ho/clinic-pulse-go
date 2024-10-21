package domain

import "time"  

type User struct {
    ID        string    `gorm:"primary_key"`
    Email     string
    CreatedAt time.Time  
    UpdatedAt time.Time  
}