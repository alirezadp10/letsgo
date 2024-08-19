package response

import "encoding/json"

func Error(message string) []byte {
    m, _ := json.Marshal(struct {
        Message string `json:"message"`
    }{
        Message: message,
    })
    return m
}
