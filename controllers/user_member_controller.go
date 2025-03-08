package controllers

import (
	"encoding/json"
	"go-pos/model"
	"net/http"
	"strconv"
	"time"
)

// UserMemberController handles UserMember CRUD operations
type UserMemberController struct {
	BaseController
}

// Create adds a new user member log entry
func (c *UserMemberController) Create() {
	var userMember model.UserMember
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userMember); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Set current time if not provided
	if userMember.Date.IsZero() {
		userMember.Date = time.Now()
	}
	
	// Set IP from request if not provided
	if userMember.IP == "" {
		userMember.IP = c.Ctx.Input.IP()
	}
	
	// Set user agent from request if not provided
	if userMember.PlatformBrowser == "" {
		userMember.PlatformBrowser = c.Ctx.Request.UserAgent()
	}
	
	// TODO: Implement repository call to save the user member log
	// For now, mock the response
	userMember.ID = 1 // Mocked ID
	
	c.JSONResponse(http.StatusCreated, "User member log created successfully", userMember)
}

// Get retrieves a user member log by ID
func (c *UserMemberController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to fetch the user member log
	// For now, mock the response
	userMember := &model.UserMember{
		ID:              id,
		MemberID:        1,
		Date:            time.Now(),
		IP:              "192.168.1.1",
		PlatformBrowser: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}
	
	c.JSONResponse(http.StatusOK, "User member log retrieved successfully", userMember)
}

// GetAll retrieves all user member logs
func (c *UserMemberController) GetAll() {
	// Check for optional filter by member ID
	memberIDStr := c.GetString("member_id")
	var memberID int
	var err error
	
	if memberIDStr != "" {
		memberID, err = strconv.Atoi(memberIDStr)
		if err != nil {
			c.JSONResponse(http.StatusBadRequest, "Invalid member ID format", nil)
			return
		}
	}
	
	// TODO: Implement repository call to fetch user member logs, filtered by member ID if provided
	// For now, mock the response
	userMembers := []model.UserMember{
		{
			ID:              1,
			MemberID:        1,
			Date:            time.Now().Add(-1 * time.Hour),
			IP:              "192.168.1.1",
			PlatformBrowser: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		},
		{
			ID:              2,
			MemberID:        2,
			Date:            time.Now(),
			IP:              "192.168.1.2",
			PlatformBrowser: "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3)",
		},
	}
	
	// Filter by member ID if specified
	if memberIDStr != "" {
		var filteredLogs []model.UserMember
		for _, log := range userMembers {
			if log.MemberID == memberID {
				filteredLogs = append(filteredLogs, log)
			}
		}
		userMembers = filteredLogs
	}
	
	c.JSONResponse(http.StatusOK, "User member logs retrieved successfully", userMembers)
}

// Update updates a user member log
func (c *UserMemberController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	var userMember model.UserMember
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userMember); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	userMember.ID = id
	
	// TODO: Implement repository call to update the user member log
	
	c.JSONResponse(http.StatusOK, "User member log updated successfully", userMember)
}

// Delete deletes a user member log
func (c *UserMemberController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to delete the user member log
	_ = id //temporary fix
	
	c.JSONResponse(http.StatusOK, "User member log deleted successfully", nil)
}
