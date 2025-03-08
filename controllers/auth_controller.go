package controllers

import (
	"encoding/json"
	"go-pos/model"
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
	
	// TODO: Implement repository call to find user by NIK
	// For now, mock the user
	user := &model.User{
		ID:           1,
		NIK:          loginReq.NIK,
		Name:         "John Doe",
		Address:      "123 Main St",
		Phone:        987654321,
		Gender:       model.GenderMale,
		IsAdmin:      false,
		PasswordHash: "$2a$10$XgNp5NSnK8WbgJZJGyfkQe1J3yMT9reKP/Ia6Q4qbsFMLZU2USc26", // hashed "password"
		Token:        "",
	}
	
	// Verify password
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password))
	if err != nil {
		c.JSONResponse(http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}
	
	// Generate a new token
	token := uuid.New().String()
	
	// Update the user's token
	user.Token = token
	
	// TODO: Save the updated token to the database
	
	// Create a user log entry
	userLog := &model.UserLog{
		UserID:          user.ID,
		Date:            time.Now(),
		IP:              c.Ctx.Input.IP(),
		PlatformBrowser: c.Ctx.Request.UserAgent(),
	}
	
	// TODO: Save the user log
	_ = userLog //temporary fix
	
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
	
	// TODO: Implement repository call to invalidate the token
	// This could involve:
	// 1. Finding the user by token
	// 2. Setting the token to empty or generating a new one
	
	c.JSONResponse(http.StatusOK, "Logout successful", nil)
}
