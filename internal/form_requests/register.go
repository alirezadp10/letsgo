package form_requests

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/alirezadp10/letsgo/internal/models"
    "github.com/alirezadp10/letsgo/internal/utils"
    "net/http"
)

func RegisterFormRequest(r *http.Request) (models.User, error) {
    var userReq models.User

    // Decode JSON body and handle errors
    if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
        return models.User{}, fmt.Errorf("failed to decode request body: %w", err)
    }

    // Hash the password and handle errors
    hash, err := utils.Hash(userReq.Password)
    if err != nil {
        return models.User{}, fmt.Errorf("failed to hash password: %w", err)
    }

    // Validate the user request
    if err := validateRegisterForm(userReq); err != nil {
        return models.User{}, fmt.Errorf("validation failed: %w", err)
    }

    // Create new user with hashed password
    newUser := models.User{
        Name:     userReq.Name,
        Username: userReq.Username,
        Password: hash,
    }

    return newUser, nil
}

func validateRegisterForm(u models.User) error {
    if u.Username == "" || u.Name == "" || u.Password == "" {
        return errors.New("missing required fields")
    }
    return nil
}
