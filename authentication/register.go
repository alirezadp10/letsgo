package authentication

import (
    "encoding/json"
    "fmt"
    "github.com/alirezadp10/letsgo/internal/db"
    "github.com/alirezadp10/letsgo/internal/form_requests"
    "github.com/alirezadp10/letsgo/internal/models"
    "github.com/alirezadp10/letsgo/internal/utils"
    "net/http"
)

func Foo(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    db.Connection().Find(&users)
    jsonData, _ := json.Marshal(users)
    fmt.Println(string(jsonData))
    w.Header().Set("Content-Type", "application/json")
    _, _ = w.Write(jsonData)
}

func Register(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    newUser, err := form_requests.RegisterFormRequest(r)

    if err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        w.Write(utils.Error(err.Error()))
        return
    }

    result := db.Connection().Create(&newUser)
    if result.Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(utils.Error(result.Error.Error()))
        return
    }

    jsonResponse, _ := json.Marshal(map[string]interface{}{
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
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write(jsonResponse)
}
