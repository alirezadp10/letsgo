package authentication

import (
    "encoding/json"
    "errors"
    "github.com/alirezadp10/letsgo/internal/db"
    "github.com/alirezadp10/letsgo/internal/form_requests"
    "github.com/alirezadp10/letsgo/internal/models"
    "github.com/alirezadp10/letsgo/internal/utils"
    "gorm.io/gorm"
    "net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    userReq, err := form_requests.LoginFormRequest(r)

    if err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        w.Write(utils.Error(err.Error()))
        return
    }

    var user models.User

    result := db.Connection().Where("username = ?", userReq.Username).Find(&user)

    if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(utils.Error("Username or password is incorrect."))
        return
    }

    if !utils.Verify(userReq.Password, user.Password) {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(utils.Error("Username or password is incorrect."))
        return
    }

    token, _ := utils.GenerateJWT(user.Username)

    jsonResponse, _ := json.Marshal(map[string]interface{}{
        "status":  "success",
        "message": "User logged in successfully",
        "data": map[string]interface{}{
            "access_token": token.AccessToken,
            "expire_at":    token.ExpireAt,
        },
    })
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write(jsonResponse)
}
