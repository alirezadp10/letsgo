package main

import (
    "github.com/alirezadp10/letsgo/authentication"
    "github.com/alirezadp10/letsgo/internal/db"
    "github.com/alirezadp10/letsgo/internal/models"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
    db.Connection().AutoMigrate(&models.User{})
    r := mux.NewRouter()
    r.HandleFunc("/login", authentication.Login).Methods("POST")
    r.HandleFunc("/register", authentication.Register).Methods("POST")
    r.HandleFunc("/foo", authentication.Foo).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", r))
}
