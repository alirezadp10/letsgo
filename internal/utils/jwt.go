package utils

import (
    "github.com/alirezadp10/letsgo/internal/configs"
    "github.com/golang-jwt/jwt/v5"
    "strconv"
    "time"
)

type Token struct {
    AccessToken string
    ExpireAt    string
}

func GenerateJWT(userID string) (Token, error) {
    var jwtSecret = []byte(configs.JWT()["secret"])

    // Create the token
    tokenLifeTime, _ := strconv.Atoi(configs.JWT()["tokenLifeTime"])
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Unix(),
    })

    // Sign the token with the secret key
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return Token{}, err
    }

    return Token{
        AccessToken: tokenString,
        ExpireAt:    time.Now().Add(time.Hour * time.Duration(tokenLifeTime)).Format("2006-01-02 15:04:05"),
    }, nil
}
