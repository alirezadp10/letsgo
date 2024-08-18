package db

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "sync"
)

func Connection() *gorm.DB {
    var once sync.Once
    var db *gorm.DB
    var err error

    once.Do(func() {
        dsn := "root:password@tcp(127.0.0.1:3306)/letsgo?charset=utf8mb4&parseTime=True&loc=Local"
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatal("Failed to connect to the database:", err)
        }
        log.Println("Database connection established")
    })

    return db
}
