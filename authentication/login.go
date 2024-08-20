package authentication

import (
    "encoding/json"
    "fmt"
    "github.com/alirezadp10/letsgo/internal/utils"
    "io"
    "net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)

    fmt.Printf("%v", string(body))

    //utils.Verify()

    //passwordToCheck := "mySecretPassword"

    w.Header().Set("Content-Type", "application/json")
    token, _ := utils.GenerateJWT("123")

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
