package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

type UserBody struct {
	Name string `json"name"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var gouser []User
		_ = db.Table("gouser").Select("id, name").Where("id=4").Scan(&gouser)
		fmt.Println(gouser)
	}
}

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var RequestBody UserBody
		json.NewDecoder(r.Body).Decode(&RequestBody)
		fmt.Println(RequestBody)
		json.NewEncoder(w).Encode(RequestBody)
	}
}
