package form_requests

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/alirezadp10/letsgo/internal/models"
    "net/http"
)

func LoginFormRequest(r *http.Request) (models.User, error) {
    var userReq models.User

    if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
        return models.User{}, fmt.Errorf("failed to decode request body: %w", err)
    }

    // Validate the user request
    if err := validateLoginForm(userReq); err != nil {
        return models.User{}, fmt.Errorf("validation failed: %w", err)
    }

    return userReq, nil
}

func validateLoginForm(u models.User) error {
    if u.Username == "" || u.Password == "" {
        return errors.New("missing required fields")
    }
    return nil
}
