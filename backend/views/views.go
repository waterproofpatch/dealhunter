package views

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"deals/database"
	"deals/decorators"
	"deals/logging"
	"deals/models"
	"deals/tokens"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// init views handle to the db
func Init() (*http.Handler, *mux.Router) {
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
	var req SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user already exists
	var user models.User
	if err := database.GetDb().Where("email = ?", req.Email).First(&user).Error; err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// Hash and salt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create the new user
	newUser := models.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Reputation:   0,
	}
	if err := database.GetDb().Create(&newUser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessToken := tokens.GenerateAccessToken(newUser)
	logging.GetLogger().Debug().Msgf("Generated access token: %v", accessToken)
	// Create a response object with the access token
	response := models.JwtAccessToken{
		AccessToken: accessToken,
	}

	// Convert the response object to JSON and write it to the response writer
	json.NewEncoder(w).Encode(response)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var req SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the user with the given email
	var user models.User
	if err := database.GetDb().Where("email = ?", req.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Compare the stored hashed password with the password provided by the user
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	accessToken := tokens.GenerateAccessToken(user)
	logging.GetLogger().Debug().Msgf("Generated access token: %v", accessToken)
	// Create a response object with the access token
	response := models.JwtAccessToken{
		AccessToken: accessToken,
	}

	// Convert the response object to JSON and write it to the response writer
	json.NewEncoder(w).Encode(response)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
}

func GetDeals(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token")
	fmt.Printf("Token=%v", token)
	var deals []models.Deal
	database.GetDb().Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func CreateDeal(w http.ResponseWriter, r *http.Request) {
	var deal models.Deal
	_ = json.NewDecoder(r.Body).Decode(&deal)
	database.GetDb().Create(&deal)
	var deals []models.Deal
	database.GetDb().Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func DeleteDeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var deal models.Deal
	database.GetDb().First(&deal, id)
	database.GetDb().Delete(&deal)
	var deals []models.Deal
	database.GetDb().Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func UpdateDeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	vote := r.URL.Query().Get("vote")

	var deal models.Deal
	database.GetDb().First(&deal, id)

	if vote == "up" {
		deal.Upvotes++
		deal.LastUpvoteTime = time.Now()
		database.GetDb().Save(&deal)
	}
	if vote == "down" {
		deal.Upvotes--
		deal.LastUpvoteTime = time.Now()
		database.GetDb().Save(&deal)
	}

	var deals []models.Deal
	database.GetDb().Preload("Location").Find(&deals)
	json.NewEncoder(w).Encode(deals)
}
