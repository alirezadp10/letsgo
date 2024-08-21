package controllers

import (
    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"
    "net/http"
)

func Home(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    name := claims["name"].(string)
    return c.String(http.StatusOK, "Welcome "+name+"!")
}
