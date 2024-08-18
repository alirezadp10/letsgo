package authentication

import (
    "encoding/json"
    "fmt"
    "github.com/alirezadp10/letsgo/internal/db"
    "github.com/alirezadp10/letsgo/internal/models"
    "github.com/alirezadp10/letsgo/internal/utils"
    "io"
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
    var userReq models.User
    _ = json.NewDecoder(r.Body).Decode(&userReq)
    w.Header().Set("Content-Type", "application/json")

    newUser := models.User{Name: userReq.Name, Username: userReq.Username, Password: userReq.Password}
    result := db.Connection().Create(&newUser)
    if result.Error != nil {
        w.WriteHeader(http.StatusInternalServerError)
        io.WriteString(w, string(utils.JsonStatus(result.Error.Error())))
        return
    }

    w.WriteHeader(http.StatusOK)
    response := map[string]interface{}{
        "status":  "success",
        "message": "User registered successfully",
        "data": map[string]interface{}{
            "id":         newUser.ID,
            "name":       newUser.Name,
            "username":   newUser.Username,
            "created_at": newUser.CreatedAt,
            "updated_at": newUser.UpdatedAt,
        },
    }
    jsonResponse, _ := json.Marshal(response)
    _, _ = w.Write(jsonResponse)
}
