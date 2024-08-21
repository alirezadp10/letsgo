package main

import (
    "github.com/alirezadp10/letsgo/authentication"
    "github.com/alirezadp10/letsgo/internal/db"
    "github.com/alirezadp10/letsgo/internal/models"
    _ "github.com/go-sql-driver/mysql"
    "github.com/labstack/echo/v4"
)

func main() {
    db.Connection().AutoMigrate(&models.User{})
    e := echo.New()

    e.POST("/login", authentication.Login)
    e.POST("/register", authentication.Register)
    e.Logger.Fatal(e.Start(":8000"))
}
