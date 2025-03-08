package controllers

import (
	"encoding/json"
	"go-pos/model"
	"go-pos/repository"
	"net/http"
	"strconv"
	"time"
)

// UserLogController handles UserLog CRUD operations
type UserLogController struct {
	BaseController
}

// Create adds a new user log entry
func (c *UserLogController) Create() {
	var userLog model.UserLog
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userLog); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Set current time if not provided
	if userLog.Date.IsZero() {
		userLog.Date = time.Now()
	}
	
	// Set IP from request if not provided
	if userLog.IP == "" {
		userLog.IP = c.Ctx.Input.IP()
	}
	
	// Set user agent from request if not provided
	if userLog.PlatformBrowser == "" {
		userLog.PlatformBrowser = c.Ctx.Request.UserAgent()
	}
	
	// Create repository instance
	repo := repository.NewUserLogRepository()
	
	// Save the user log
	newUserLog, err := repo.CreateUserLog(&userLog)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to create user log: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusCreated, "User log created successfully", newUserLog)
}

// Get retrieves a user log by ID
func (c *UserLogController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Create repository instance
	repo := repository.NewUserLogRepository()
	
	// Fetch the user log
	userLog, err := repo.GetUserLog(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "User log not found", nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "User log retrieved successfully", userLog)
}

// GetAll retrieves all user logs
func (c *UserLogController) GetAll() {
	// Check for optional filter by user ID
	userIDStr := c.GetString("user_id")
	var userID int
	var err error
	
	if userIDStr != "" {
		userID, err = strconv.Atoi(userIDStr)
		if err != nil {
			c.JSONResponse(http.StatusBadRequest, "Invalid user ID format", nil)
			return
		}
	}
	
	// Create repository instance
	repo := repository.NewUserLogRepository()
	
	var userLogs []model.UserLog
	
	// Fetch user logs, filtered by user ID if provided
	if userIDStr != "" {
		userLogs, err = repo.GetUserLogsByUserID(userID)
	} else {
		userLogs, err = repo.GetAllUserLogs()
	}
	
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve user logs: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "User logs retrieved successfully", userLogs)
}

// Update updates a user log
func (c *UserLogController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	var userLog model.UserLog
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userLog); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	userLog.ID = id
	
	// Create repository instance
	repo := repository.NewUserLogRepository()
	
	// Check if user log exists
	_, err = repo.GetUserLog(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "User log not found", nil)
		return
	}
	
	// Update user log
	updatedUserLog, err := repo.UpdateUserLog(&userLog)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to update user log: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "User log updated successfully", updatedUserLog)
}

// Delete deletes a user log
func (c *UserLogController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Create repository instance
	repo := repository.NewUserLogRepository()
	
	// Check if user log exists
	_, err = repo.GetUserLog(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "User log not found", nil)
		return
	}
	
	// Delete user log
	err = repo.DeleteUserLog(id)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to delete user log: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "User log deleted successfully", nil)
}
