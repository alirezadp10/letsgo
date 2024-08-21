package controllers

import (
    "errors"
    "github.com/alirezadp10/letsgo/internal/db"
    "github.com/alirezadp10/letsgo/internal/form_requests"
    "github.com/alirezadp10/letsgo/internal/models"
    "github.com/alirezadp10/letsgo/pkg/utils"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
    "net/http"
)

func Login(c echo.Context) error {
    userReq, err := form_requests.LoginFormRequest(c)

    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
            "message": err.Error(),
        })
    }

    var user models.User

    result := db.Connection().Where("username = ?", userReq.Username).Find(&user)

    if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
            "message": "Username or password is incorrect.",
        })
    }

    if !utils.Verify(userReq.Password, user.Password) {
        return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
            "message": "Username or password is incorrect.",
        })
    }

    token, _ := utils.GenerateJWT(user.Username)

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  "success",
        "message": "User logged in successfully",
        "data": map[string]interface{}{
            "access_token": token.AccessToken,
            "expire_at":    token.ExpireAt,
        },
    })
}

func Register(c echo.Context) error {
    newUser, err := form_requests.RegisterFormRequest(c)

    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
            "message": err.Error(),
        })
    }

    result := db.Connection().Create(&newUser)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": result.Error.Error(),
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status":  "success",
        "message": "User registered successfully",
        "data": map[string]interface{}{
            "id":         newUser.ID,
            "name":       newUser.Name,
            "username":   newUser.Username,
            "created_at": newUser.CreatedAt,
            "updated_at": newUser.UpdatedAt,
        },
    })
}
