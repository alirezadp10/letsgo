package main

import (
    "github.com/alirezadp10/letsgo/internal/configs"
    "github.com/alirezadp10/letsgo/internal/controllers"
    "github.com/alirezadp10/letsgo/internal/db"
    "github.com/alirezadp10/letsgo/internal/models"
    _ "github.com/go-sql-driver/mysql"
    "github.com/labstack/echo-jwt/v4"
    "github.com/labstack/echo/v4"
)

func main() {
    db.Connection().AutoMigrate(&models.User{})

    e := echo.New()
    e.POST("/login", controllers.Login)
    e.POST("/register", controllers.Register)
    e.GET("/home", controllers.Home, echojwt.WithConfig(echojwt.Config{
        SigningKey: []byte(configs.JWT()["secret"]),
    }))

    e.Logger.Fatal(e.Start(":8000"))
}
