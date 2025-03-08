package controllers

import (
	"encoding/json"
	"go-pos/model"
	"go-pos/repository"
	"net/http"
	"strconv"
	"go-pos/database"
)

// MemberPointController handles MemberPoint CRUD operations
type MemberPointController struct {
	BaseController
	repo *repository.MemberPointRepository
	memberRepo *repository.MemberRepository
}

// Prepare initializes the controller
func (c *MemberPointController) Prepare() {
	// Initialize the repositories
	c.repo = repository.NewMemberPointRepository()
	c.memberRepo = repository.NewMemberRepository()
}

// Create adds a new member point transaction
func (c *MemberPointController) Create() {
	var memberPoint model.MemberPoint
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &memberPoint); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Validate required fields
	if memberPoint.MemberID <= 0 {
		c.JSONResponse(http.StatusBadRequest, "Member ID is required", nil)
		return
	}
	
	if memberPoint.Points <= 0 {
		c.JSONResponse(http.StatusBadRequest, "Points must be greater than zero", nil)
		return
	}
	
	if memberPoint.Type != model.PointTypeEarned && memberPoint.Type != model.PointTypeRedeemed {
		c.JSONResponse(http.StatusBadRequest, "Point type must be either EARNED or REDEEMED", nil)
		return
	}
	
	// Check if member exists
	member, err := c.memberRepo.GetMember(memberPoint.MemberID)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Member not found", nil)
		return
	}
	
	// For REDEEMED points, check if member has sufficient points
	if memberPoint.Type == model.PointTypeRedeemed && member.Points < memberPoint.Points {
		c.JSONResponse(http.StatusBadRequest, "Member does not have sufficient points to redeem", nil)
		return
	}
	
	// Create transaction
	tx, err := database.DB.Begin()
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to start transaction: "+err.Error(), nil)
		return
	}
	
	// Save the point transaction
	newMemberPoint, err := c.repo.CreateMemberPointTx(tx, &memberPoint)
	if err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to create point transaction: "+err.Error(), nil)
		return
	}
	
	// Update member's points balance
	if memberPoint.Type == model.PointTypeEarned {
		member.Points += memberPoint.Points
	} else {
		member.Points -= memberPoint.Points
	}
	
	// Update the member record
	_, err = c.memberRepo.UpdateMemberTx(tx, member)
	if err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to update member points: "+err.Error(), nil)
		return
	}
	
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to commit transaction: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusCreated, "Member point transaction created successfully", newMemberPoint)
}

// Get retrieves a member point transaction by ID
func (c *MemberPointController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	memberPoint, err := c.repo.GetMemberPoint(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Member point transaction not found", nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Member point transaction retrieved successfully", memberPoint)
}

// GetAll retrieves all member point transactions
func (c *MemberPointController) GetAll() {
	// Check for optional member filter
	memberIDStr := c.GetString("member_id")
	typeFilter := c.GetString("type")
	
	var memberPoints []model.MemberPoint
	var err error
	
	if memberIDStr != "" {
		memberID, err := strconv.Atoi(memberIDStr)
		if err != nil {
			c.JSONResponse(http.StatusBadRequest, "Invalid member ID format", nil)
			return
		}
		
		if typeFilter != "" {
			// Validate type
			if typeFilter != string(model.PointTypeEarned) && typeFilter != string(model.PointTypeRedeemed) {
				c.JSONResponse(http.StatusBadRequest, "Invalid point type filter", nil)
				return
			}
			
			memberPoints, err = c.repo.GetMemberPointsByMemberAndType(memberID, model.PointType(typeFilter))
		} else {
			memberPoints, err = c.repo.GetMemberPointsByMember(memberID)
		}
	} else {
		if typeFilter != "" {
			// Validate type
			if typeFilter != string(model.PointTypeEarned) && typeFilter != string(model.PointTypeRedeemed) {
				c.JSONResponse(http.StatusBadRequest, "Invalid point type filter", nil)
				return
			}
			
			memberPoints, err = c.repo.GetMemberPointsByType(model.PointType(typeFilter))
		} else {
			memberPoints, err = c.repo.GetAllMemberPoints()
		}
	}
	
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve member point transactions: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Member point transactions retrieved successfully", memberPoints)
}

// Update updates a member point transaction
func (c *MemberPointController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	var memberPoint model.MemberPoint
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &memberPoint); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	memberPoint.ID = id
	
	// Get the original point transaction
	originalPoint, err := c.repo.GetMemberPoint(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Member point transaction not found", nil)
		return
	}
	
	// Get the member
	member, err := c.memberRepo.GetMember(originalPoint.MemberID)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve member: "+err.Error(), nil)
		return
	}
	
	// Calculate the points adjustment needed
	var pointsAdjustment int
	
	// Reverse the original transaction effect
	if originalPoint.Type == model.PointTypeEarned {
		pointsAdjustment -= originalPoint.Points
	} else {
		pointsAdjustment += originalPoint.Points
	}
	
	// Apply the new transaction effect
	if memberPoint.Type == model.PointTypeEarned {
		pointsAdjustment += memberPoint.Points
	} else {
		pointsAdjustment -= memberPoint.Points
	}
	
	// Check if the adjustment would result in negative points
	if member.Points + pointsAdjustment < 0 {
		c.JSONResponse(http.StatusBadRequest, "Update would result in negative member points", nil)
		return
	}
	
	// Create transaction
	tx, err := database.DB.Begin()
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to start transaction: "+err.Error(), nil)
		return
	}
	
	// Update the point transaction
	updatedMemberPoint, err := c.repo.UpdateMemberPointTx(tx, &memberPoint)
	if err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to update point transaction: "+err.Error(), nil)
		return
	}
	
	// Update member's points balance
	member.Points += pointsAdjustment
	
	// Update the member record
	_, err = c.memberRepo.UpdateMemberTx(tx, member)
	if err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to update member points: "+err.Error(), nil)
		return
	}
	
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to commit transaction: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Member point transaction updated successfully", updatedMemberPoint)
}

// Delete deletes a member point transaction
func (c *MemberPointController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Get the point transaction
	memberPoint, err := c.repo.GetMemberPoint(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Member point transaction not found", nil)
		return
	}
	
	// Get the member
	member, err := c.memberRepo.GetMember(memberPoint.MemberID)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve member: "+err.Error(), nil)
		return
	}
	
	// Calculate the points adjustment needed
	var pointsAdjustment int
	
	// Reverse the transaction effect
	if memberPoint.Type == model.PointTypeEarned {
		pointsAdjustment = -memberPoint.Points
	} else {
		pointsAdjustment = memberPoint.Points
	}
	
	// Check if the adjustment would result in negative points
	if member.Points + pointsAdjustment < 0 {
		c.JSONResponse(http.StatusBadRequest, "Deletion would result in negative member points", nil)
		return
	}
	
	// Create transaction
	tx, err := database.DB.Begin()
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to start transaction: "+err.Error(), nil)
		return
	}
	
	// Delete the point transaction
	err = c.repo.DeleteMemberPointTx(tx, id)
	if err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to delete point transaction: "+err.Error(), nil)
		return
	}
	
	// Update member's points balance
	member.Points += pointsAdjustment
	
	// Update the member record
	_, err = c.memberRepo.UpdateMemberTx(tx, member)
	if err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to update member points: "+err.Error(), nil)
		return
	}
	
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to commit transaction: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Member point transaction deleted successfully", nil)
}
