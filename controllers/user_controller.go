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
	// Create a new repository instance
	repo := repository.NewUserRepository()
	
	// Fetch all users from database
	users, err := repo.GetAllUsers()
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve users: "+err.Error(), nil)
		return
	}
	
	// Don't include password hash in response for security
	for i := range users {
		users[i].PasswordHash = ""
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
	
	// Create a new repository instance
	repo := repository.NewUserRepository()
	
	// Check if user exists
	existingUser, err := repo.GetUser(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "User not found", nil)
		return
	}
	
	// If password is being updated, hash it
	if user.PasswordHash != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			c.JSONResponse(http.StatusInternalServerError, "Failed to hash password", nil)
			return
		}
		user.PasswordHash = string(hashedPassword)
	} else {
		// Keep the existing password
		user.PasswordHash = existingUser.PasswordHash
	}
	
	// Update user in database
	updatedUser, err := repo.UpdateUser(&user)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to update user: "+err.Error(), nil)
		return
	}
	
	// Don't return sensitive information
	updatedUser.PasswordHash = ""
	
	c.JSONResponse(http.StatusOK, "User updated successfully", updatedUser)
}

// Delete deletes a user
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Create a new repository instance
	repo := repository.NewUserRepository()
	
	// Check if user exists
	_, err = repo.GetUser(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "User not found", nil)
		return
	}
	
	// Delete user from database
	err = repo.DeleteUser(id)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to delete user: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "User deleted successfully", nil)
}