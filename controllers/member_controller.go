package controllers

import (
	"encoding/json"
	"go-pos/model"
	"net/http"
	"strconv"
	"time"
)

// MemberController handles Member CRUD operations
type MemberController struct {
	BaseController
}

// Create adds a new member
func (c *MemberController) Create() {
	var member model.Member
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &member); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Set default values
	member.JoinDate = time.Now()
	member.Points = 0
	
	// TODO: Hash password and generate token
	// TODO: Implement repository call to save the member
	
	// For now, mock the response
	member.ID = 1 // Mocked ID
	
	c.JSONResponse(http.StatusCreated, "Member created successfully", member)
}

// Get retrieves a member by ID
func (c *MemberController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to fetch the member
	// For now, mock the response
	member := &model.Member{
		ID:       id,
		Name:     "Sample Member",
		Phone:    123456789,
		JoinDate: time.Now(),
		Points:   100,
	}
	
	c.JSONResponse(http.StatusOK, "Member retrieved successfully", member)
}

// GetAll retrieves all members
func (c *MemberController) GetAll() {
	// TODO: Implement repository call to fetch all members
	// For now, mock the response
	members := []model.Member{
		{ID: 1, Name: "Member 1", Phone: 123456789, JoinDate: time.Now(), Points: 100},
		{ID: 2, Name: "Member 2", Phone: 987654321, JoinDate: time.Now(), Points: 200},
	}
	
	c.JSONResponse(http.StatusOK, "Members retrieved successfully", members)
}

// Update updates a member
func (c *MemberController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	var member model.Member
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &member); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	member.ID = id
	
	// TODO: Implement repository call to update the member
	
	c.JSONResponse(http.StatusOK, "Member updated successfully", member)
}

// Delete deletes a member
func (c *MemberController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to delete the member
	_ = id // Temporarily use the id variable to avoid unused variable error
	
	c.JSONResponse(http.StatusOK, "Member deleted successfully", nil)
}
