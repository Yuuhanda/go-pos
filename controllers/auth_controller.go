package controllers

import (
	"encoding/json"
	"go-pos/model"
	"go-pos/repository"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

// AuthController handles authentication operations
type AuthController struct {
	BaseController
}

// LoginRequest represents the login request body
type LoginRequest struct {
	NIK      int    `json:"nik"`
	Password string `json:"password"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}

// Login authenticates a user and returns a token
func (c *AuthController) Login() {
	var loginReq LoginRequest
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &loginReq); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Create repository instances
	userRepo := repository.NewUserRepository()
	userLogRepo := repository.NewUserLogRepository()
	
	// Find user by NIK
	user, err := userRepo.GetUserByNIK(loginReq.NIK)
	if err != nil {
		c.JSONResponse(http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}
	
	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password))
	if err != nil {
		c.JSONResponse(http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}
	
	// Generate a new token
	token := uuid.New().String()
	
	// Update the user's token
	user.Token = token
	_, err = userRepo.UpdateUser(user)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to update user token", nil)
		return
	}
	
	// Create a user log entry
	userLog := &model.UserLog{
		UserID:          user.ID,
		Date:            time.Now(),
		IP:              c.Ctx.Input.IP(),
		PlatformBrowser: c.Ctx.Request.UserAgent(),
	}
	
	// Save the user log
	_, err = userLogRepo.CreateUserLog(userLog)
	if err != nil {
		// Log the error but continue - non-critical operation
		// In a production environment, you might want to log this error properly
		// fmt.Printf("Error logging user login: %v\n", err)
	}
	
	// Don't include password hash in response
	user.PasswordHash = ""
	
	response := LoginResponse{
		User:  *user,
		Token: token,
	}
	
	c.JSONResponse(http.StatusOK, "Login successful", response)
}

// Logout logs out a user
func (c *AuthController) Logout() {
	// Get the authorization token
	token := c.Ctx.Input.Header("Authorization")
	if token == "" {
		c.JSONResponse(http.StatusBadRequest, "No authorization token provided", nil)
		return
	}
	
	// Create repository instance
	userRepo := repository.NewUserRepository()
	
	// Find the user by token
	user, err := userRepo.GetUserByToken(token)
	if err != nil {
		c.JSONResponse(http.StatusUnauthorized, "Invalid token", nil)
		return
	}
	
	// Invalidate the token by generating a new one
	user.Token = uuid.New().String() // Generate a new token that won't match what the client has
	
	// Update the user
	_, err = userRepo.UpdateUser(user)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to update user token", nil)
		return
	}
	
	// Create user log for logout
	userLogRepo := repository.NewUserLogRepository()
	userLog := &model.UserLog{
		UserID:          user.ID,
		Date:            time.Now(),
		IP:              c.Ctx.Input.IP(),
		PlatformBrowser: c.Ctx.Request.UserAgent(),
		Action:          "LOGOUT", // Add this field to the UserLog model
	}
	
	// Save the logout log
	_, err = userLogRepo.CreateUserLog(userLog)
	if err != nil {
		// Log the error but continue - non-critical operation
	}
	
	c.JSONResponse(http.StatusOK, "Logout successful", nil)
}
