package views

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"deals/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

var db *gorm.DB

func TokenDecorator(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		ctx := context.WithValue(r.Context(), "token", token)
		h(w, r.WithContext(ctx))
	}
}

// init views handle to the db
func Init(_db *gorm.DB) (*http.Handler, *mux.Router) {
	db = _db
	r := mux.NewRouter()

	r.HandleFunc("/deals", TokenDecorator(GetDeals)).Methods("GET")
	r.HandleFunc("/deals", TokenDecorator(CreateDeal)).Methods("POST")
	r.HandleFunc("/deals/{id}", TokenDecorator(DeleteDeal)).Methods("DELETE")
	r.HandleFunc("/deals/{id}", TokenDecorator(UpdateDeal)).Methods("PUT")
	r.HandleFunc("/auth/login", TokenDecorator(Login)).Methods("POST")
	r.HandleFunc("/auth/logout", TokenDecorator(Logout)).Methods("POST")
	r.HandleFunc("/auth/register", TokenDecorator(Register)).Methods("POST")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // your website
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	handler := corsOptions.Handler(r)
	return &handler, r
}

func Register(w http.ResponseWriter, r *http.Request) {
}

func Logout(w http.ResponseWriter, r *http.Request) {
}

func Login(w http.ResponseWriter, r *http.Request) {
}

func GetDeals(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token")
	fmt.Printf("Token=%v", token)
	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func CreateDeal(w http.ResponseWriter, r *http.Request) {
	var deal models.Deal
	_ = json.NewDecoder(r.Body).Decode(&deal)
	fmt.Printf("db=%v", db)
	db.Create(&deal)
	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func DeleteDeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var deal models.Deal
	db.First(&deal, id)
	db.Delete(&deal)
	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func UpdateDeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	vote := r.URL.Query().Get("vote")

	var deal models.Deal
	db.First(&deal, id)

	if vote == "up" {
		deal.Upvotes++
		deal.LastUpvoteTime = time.Now()
		db.Save(&deal)
	}
	if vote == "down" {
		deal.Upvotes--
		deal.LastUpvoteTime = time.Now()
		db.Save(&deal)
	}

	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}
