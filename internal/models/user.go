package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name     string `gorm:"size:100"`
    Username string `gorm:"uniqueIndex"`
    Password string `gorm:"size:100"`
}
