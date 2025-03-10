package controllers

import (
	"encoding/json"
	"go-pos/model"
	"go-pos/repository"
	"net/http"
	"strconv"
	"time"
)

// MemberController handles Member CRUD operations
type MemberController struct {
	BaseController
	repo *repository.MemberRepository
}

// Prepare initializes the controller
func (c *MemberController) Prepare() {
	// Initialize the repository
	c.repo = repository.NewMemberRepository()
}

// Create adds a new member
func (c *MemberController) Create() {
	var member model.Member
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &member); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Validate required fields
	if member.Name == "" {
		c.JSONResponse(http.StatusBadRequest, "Member name is required", nil)
		return
	}
	
	if member.Phone <= 0 {
		c.JSONResponse(http.StatusBadRequest, "Valid phone number is required", nil)
		return
	}
	
	// Set default values
	member.JoinDate = time.Now()
	member.Points = 0
	
	// Save the member to database
	newMember, err := c.repo.CreateMember(&member)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to create member: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusCreated, "Member created successfully", newMember)
}

// Get retrieves a member by ID
func (c *MemberController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	member, err := c.repo.GetMember(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Member not found", nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Member retrieved successfully", member)
}

// GetAll retrieves all members
func (c *MemberController) GetAll() {
	// Check for optional phone filter
	phoneStr := c.GetString("phone")
	
	var members []model.Member
	var err error

	if phoneStr != "" {
		var phone int
		phone, err = strconv.Atoi(phoneStr)
		if err != nil {
			c.JSONResponse(http.StatusBadRequest, "Invalid phone number format", nil)
			return
		}
		
		members, err = c.repo.GetMembersByPhone(phone)
	} else {
		members, err = c.repo.GetAllMembers()
	}

	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve members: "+err.Error(), nil)
		return
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
	
	// Check if member exists
	_, err = c.repo.GetMember(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Member not found", nil)
		return
	}
	
	// Update member
	updatedMember, err := c.repo.UpdateMember(&member)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to update member: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Member updated successfully", updatedMember)
}

// Delete deletes a member
func (c *MemberController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Check if member exists
	_, err = c.repo.GetMember(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Member not found", nil)
		return
	}
	
	// Check if member has any sales baskets (can't delete a member with transactions)
	hasSales, err := c.repo.MemberHasSales(id)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to check member sales: "+err.Error(), nil)
		return
	}
	
	if hasSales {
		c.JSONResponse(http.StatusBadRequest, "Cannot delete member: has existing sales records", nil)
		return
	}
	
	// Delete member
	err = c.repo.DeleteMember(id)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to delete member: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Member deleted successfully", nil)
}

// GetAllPoints retrieves all member points
func (c *MemberController) GetAllPoints() {
    // Get member points from repository
    memberPointRepo := repository.NewMemberPointRepository()
    points, err := memberPointRepo.GetAllMemberPoints()
    
    if err != nil {
        c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve member points: "+err.Error(), nil)
        return
    }
    
    c.JSONResponse(http.StatusOK, "Member points retrieved successfully", points)
}
