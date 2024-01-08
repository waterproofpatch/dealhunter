package views

import (
	"context"
	"crypto/rand"
	"encoding/hex"
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

	r.HandleFunc("/user-meta", TokenDecorator(UserMeta)).Methods("GET")
	r.HandleFunc("/deals", TokenDecorator(GetDeals)).Methods("GET")
	r.HandleFunc("/deals", TokenDecorator(CreateDeal)).Methods("POST")
	r.HandleFunc("/deals/{id}", TokenDecorator(UpdateDeal)).Methods("PUT")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // your website
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	handler := corsOptions.Handler(r)
	return &handler, r
}

func GetDeals(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token")
	fmt.Printf("Token=%v", token)
	var deals []models.Deal
	db.Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func UserMeta(w http.ResponseWriter, r *http.Request) {
	var userMeta models.UserMeta
	bytes := make([]byte, 16) // generate 16 bytes
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	userMeta.Token = hex.EncodeToString(bytes)
	json.NewEncoder(w).Encode(userMeta)
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
