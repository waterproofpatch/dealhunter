package views

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"deals/cookies"
	"deals/database"
	"deals/decorators"
	"deals/logging"
	"deals/models"
	"deals/tokens"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	r.HandleFunc("/deals", decorators.LogDecorator(decorators.TokenDecorator(true, GetDeals))).Methods("GET")
	r.HandleFunc("/deals", decorators.LogDecorator(decorators.TokenDecorator(false, CreateDeal))).Methods("POST")
	r.HandleFunc("/deals/{id}", decorators.LogDecorator(decorators.TokenDecorator(false, DeleteDeal))).Methods("DELETE")
	r.HandleFunc("/deals/{id}", decorators.LogDecorator(decorators.TokenDecorator(false, UpdateDeal))).Methods("PUT")
	r.HandleFunc("/auth/signin", decorators.LogDecorator(SignIn)).Methods("POST")
	r.HandleFunc("/auth/logout", decorators.LogDecorator(SignOut)).Methods("POST")
	r.HandleFunc("/auth/signup", decorators.LogDecorator(SignUp)).Methods("POST")
	r.HandleFunc("/auth/refresh", decorators.LogDecorator(Refresh)).Methods("GET")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://localhost:4300", "https://main.d1gmmk0f0epwh2.amplifyapp.com"}, // your website
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Origin", "Authorization"},
		AllowCredentials: true,
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
		IsAdmin:      false,
		Reputation:   0,
	}
	if err := database.GetDb().Create(&newUser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// mint some tokens
	accessToken := tokens.GenerateAccessToken(newUser)
	logging.GetLogger().Debug().Msgf("Generated access token: %v", accessToken)

	refreshToken := tokens.GenerateRefrehToken(newUser)
	logging.GetLogger().Debug().Msgf("Generated refresh token: %v", refreshToken)

	cookie := cookies.SetRefreshTokenCookie(refreshToken)
	http.SetCookie(w, cookie)
	logging.GetLogger().Debug().Msgf("set refreshToken cookie")

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

	refreshToken := tokens.GenerateRefrehToken(user)
	logging.GetLogger().Debug().Msgf("Generated refresh token: %v", refreshToken)

	cookie := cookies.SetRefreshTokenCookie(refreshToken)
	http.SetCookie(w, cookie)
	logging.GetLogger().Debug().Msgf("set refreshToken cookie")

	response := models.JwtAccessToken{
		AccessToken: accessToken,
	}

	// Convert the response object to JSON and write it to the response writer
	json.NewEncoder(w).Encode(response)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	// Read the refresh token from the cookie
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Missing refresh token", http.StatusBadRequest)
		return
	}

	// Verify the refresh token
	claims, err := tokens.VerifyRefreshToken(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := database.GetDb().Where("email = ?", claims["email"]).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Generate a new access token
	accessToken := tokens.GenerateAccessToken(user)
	logging.GetLogger().Debug().Msgf("Generated access token: %v", accessToken)

	// Generate a new refresh token
	refreshToken := tokens.GenerateRefrehToken(user)
	logging.GetLogger().Debug().Msgf("Generated refresh token: %v", refreshToken)

	// Set the new refresh token as a cookie
	newCookie := cookies.SetRefreshTokenCookie(refreshToken)
	http.SetCookie(w, newCookie)
	logging.GetLogger().Debug().Msgf("Set refresh token cookie")

	// Create the response object
	response := models.JwtAccessToken{
		AccessToken: accessToken,
	}

	// Convert the response object to JSON and write it to the response writer
	json.NewEncoder(w).Encode(response)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
}

func GetDeals(w http.ResponseWriter, r *http.Request) {
	var deals []models.Deal
	db := database.GetDb().Preload("Location").Preload("User")
	if db.Error != nil {
		// Handle preload error
		logging.GetLogger().Error().Msgf("Error: %v", db.Error)
		http.Error(w, "Failed fetching deals.", http.StatusInternalServerError)
		return
	}
	db.Find(&deals)
	logging.GetLogger().Debug().Msgf("Returning %d deals...", len(deals))
	json.NewEncoder(w).Encode(deals)
}

func CreateDeal(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Context().Value("token").(jwt.MapClaims)
	var deal models.Deal
	_ = json.NewDecoder(r.Body).Decode(&deal)

	// Assign the deal's user to the user from token["id"]
	tokenUserId, err := strconv.ParseInt(token["id"].(string), 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}
	deal.UserID = uint(tokenUserId)

	logging.GetLogger().Debug().Msgf("Adding deal for user %d (%s)", deal.UserID, token["email"])

	// Check if creating the deal fails
	result := database.GetDb().Create(&deal)
	if result.Error != nil {
		logging.GetLogger().Error().Msgf("Failed creating deal: %v", result.Error)
		http.Error(w, "Failed creating deal", http.StatusInternalServerError)
		return
	}

	var deals []models.Deal
	db := database.GetDb().Preload("Location").Preload("User")
	if db.Error != nil {
		// Handle preload error
		logging.GetLogger().Error().Msgf("Error: %v", db.Error)
		http.Error(w, "Failed fetching deals.", http.StatusInternalServerError)
		return
	}
	db.Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func DeleteDeal(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Context().Value("token").(jwt.MapClaims)
	params := mux.Vars(r)
	id := params["id"]
	var deal models.Deal
	tokenUserId, err := strconv.ParseInt(token["id"].(string), 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}
	database.GetDb().Preload("User").First(&deal, id)
	if deal.User.ID != uint(tokenUserId) {
		http.Error(w, "Cannot delete a deal that you didn't post!", http.StatusUnauthorized)
		return
	}
	database.GetDb().Delete(&deal)
	var deals []models.Deal
	db := database.GetDb().Preload("Location").Preload("User")
	if db.Error != nil {
		// Handle preload error
		logging.GetLogger().Error().Msgf("Error: %v", db.Error)
		http.Error(w, "Failed fetching deals.", http.StatusInternalServerError)
		return
	}
	db.Find(&deals)
	json.NewEncoder(w).Encode(deals)
}

func UpdateDeal(w http.ResponseWriter, r *http.Request) {
	token, _ := r.Context().Value("token").(jwt.MapClaims)
	userIDStr, ok := token["id"].(string)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	dealID := params["id"]
	vote := r.URL.Query().Get("vote")

	var existingVote models.Vote
	result := database.GetDb().Where("user_id = ? AND deal_id = ?", userID, dealID).First(&existingVote)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Failed checking existing votes.", http.StatusInternalServerError)
		return
	}

	if existingVote.ID != 0 {
		http.Error(w, "You have already voted on this deal!", http.StatusBadRequest)
		return
	}

	var deal models.Deal
	database.GetDb().Preload("User").First(&deal, dealID)

	if deal.User.ID == uint(userID) {
		http.Error(w, "Cannot vote on a deal that you posted!", http.StatusBadRequest)
		return
	}

	if vote == "up" {
		deal.Upvotes++
		deal.LastUpvoteTime = time.Now()
		database.GetDb().Save(&deal)
	} else if vote == "down" {
		deal.Upvotes--
		deal.LastUpvoteTime = time.Now()
		database.GetDb().Save(&deal)
	}

	newVote := models.Vote{
		UserID: uint(userID),
		DealID: deal.ID,
	}
	database.GetDb().Create(&newVote)

	var deals []models.Deal
	db := database.GetDb().Preload("Location").Preload("User")
	if db.Error != nil {
		logging.GetLogger().Error().Msgf("Error: %v", db.Error)
		http.Error(w, "Failed fetching deals.", http.StatusInternalServerError)
		return
	}
	db.Find(&deals)
	json.NewEncoder(w).Encode(deals)
}
