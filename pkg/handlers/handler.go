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

func GetUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var gousers []User
		_ = db.Table("gouser").Select("id, name").Scan(&gousers)
		json.NewEncoder(w).Encode(gousers)
	}
}

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var RequestBody UserBody
		json.NewDecoder(r.Body).Decode(&RequestBody)
		_ = db.Table("gouser").Create(&RequestBody)
		fmt.Println("Created User")
		json.NewEncoder(w).Encode(RequestBody)
	}
}

func UpdateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var PutBody User
		json.NewDecoder(r.Body).Decode(&PutBody)
		_ = db.Table("gouser").Where("id=?", PutBody.Id).Update("name", PutBody.Name).Scan(&PutBody)
		fmt.Printf("Updated User with id %d\n", PutBody.Id)
		json.NewEncoder(w).Encode(PutBody)
	}
}

func DeleteUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var DeleteBody User
		json.NewDecoder(r.Body).Decode(&DeleteBody)
		_ = db.Table("gouser").Delete(&DeleteBody)
		fmt.Printf("Deleted User with id %d\n", DeleteBody.Id)
		json.NewEncoder(w).Encode(DeleteBody)
	}
}
