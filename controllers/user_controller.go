package controllers

import (
	"encoding/json"
	"go-pos/model"
	"net/http"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
	"go-pos/repository"
)

// UserController handles User CRUD operations
type UserController struct {
	BaseController
}

// Create adds a new user
func (c *UserController) Create() {
	var user model.User
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to hash password", nil)
		return
	}
	user.PasswordHash = string(hashedPassword)
	
	// Generate a unique token using UUID
	token := uuid.New().String()
	user.Token = token
	
	// Create new user repository instance
	repo := repository.NewUserRepository()
	
	// Save the user to database
	newUser, err := repo.CreateUser(&user)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to create user: "+err.Error(), nil)
		return
	}
	
	// Don't return sensitive information
	newUser.PasswordHash = ""
	
	c.JSONResponse(http.StatusCreated, "User created successfully", newUser)
}

// Get retrieves a user by ID
func (c *UserController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Create a new repository instance
	repo := repository.NewUserRepository()
	
	// Fetch the user from database
	user, err := repo.GetUser(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "User not found", nil)
		return
	}
	
	// Don't include password hash in response for security
	user.PasswordHash = ""
	
	c.JSONResponse(http.StatusOK, "User retrieved successfully", user)
}

// GetAll retrieves all users
func (c *UserController) GetAll() {
	// TODO: Implement repository call to fetch all users
	// For now, mock the response
	users := []model.User{
		{
			ID:       1,
			NIK:      123456789,
			Name:     "John Doe",
			Address:  "123 Main St",
			Phone:    987654321,
			Gender:   model.GenderMale,
			IsAdmin:  false,
			Token:    "user-token-1",
			// Don't include password hash in response
		},
		{
			ID:       2,
			NIK:      987654321,
			Name:     "Jane Smith",
			Address:  "456 Oak Ave",
			Phone:    123456789,
			Gender:   model.GenderFemale,
			IsAdmin:  true,
			Token:    "user-token-2",
			// Don't include password hash in response
		},
	}
	
	c.JSONResponse(http.StatusOK, "Users retrieved successfully", users)
}

// Update updates a user
func (c *UserController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	var user model.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	user.ID = id
	
	// If password is being updated, hash it
	if user.PasswordHash != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			c.JSONResponse(http.StatusInternalServerError, "Failed to hash password", nil)
			return
		}
		user.PasswordHash = string(hashedPassword)
	}
	
	// TODO: Implement repository call to update the user
	
	// Don't return sensitive information
	user.PasswordHash = ""
	
	c.JSONResponse(http.StatusOK, "User updated successfully", user)
}

// Delete deletes a user
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to delete the user
	_ = id //temporary fix
	
	c.JSONResponse(http.StatusOK, "User deleted successfully", nil)
}
