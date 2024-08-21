package authentication

import (
    "github.com/alirezadp10/letsgo/internal/db"
    "github.com/alirezadp10/letsgo/internal/form_requests"
    "github.com/labstack/echo/v4"
    "net/http"
)

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
