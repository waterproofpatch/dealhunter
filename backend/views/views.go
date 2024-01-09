package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"deals/decorators"
	"deals/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

var db *gorm.DB

// init views handle to the db
func Init(_db *gorm.DB) (*http.Handler, *mux.Router) {
	db = _db
	r := mux.NewRouter()

	r.HandleFunc("/deals", decorators.LogDecorator(decorators.TokenDecorator(GetDeals))).Methods("GET")
	r.HandleFunc("/deals", decorators.LogDecorator(decorators.TokenDecorator(CreateDeal))).Methods("POST")
	r.HandleFunc("/deals/{id}", decorators.LogDecorator(decorators.TokenDecorator(DeleteDeal))).Methods("DELETE")
	r.HandleFunc("/deals/{id}", decorators.LogDecorator(decorators.TokenDecorator(UpdateDeal))).Methods("PUT")
	r.HandleFunc("/auth/signin", decorators.LogDecorator(SignIn)).Methods("POST")
	r.HandleFunc("/auth/logout", decorators.LogDecorator(SignOut)).Methods("POST")
	r.HandleFunc("/auth/signup", decorators.LogDecorator(SignUp)).Methods("POST")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // your website
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	handler := corsOptions.Handler(r)
	return &handler, r
}

func SignUp(w http.ResponseWriter, r *http.Request) {
}

func SignOut(w http.ResponseWriter, r *http.Request) {
}

func SignIn(w http.ResponseWriter, r *http.Request) {
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
