package controllers

import (
	"encoding/json"
	"go-pos/database" // Add this import
	"go-pos/model"
	"go-pos/repository"
	"net/http"
	"strconv"
	"time"
)

// SalesBasketController handles SalesBasket CRUD operations
type SalesBasketController struct {
	BaseController
	repo *repository.SalesBasketRepository
	itemRepo *repository.SalesItemRepository
}

// Prepare initializes the controller
func (c *SalesBasketController) Prepare() {
	// Initialize the repositories
	c.repo = repository.NewSalesBasketRepository()
	c.itemRepo = repository.NewSalesItemRepository()
}

// Create adds a new sales basket
func (c *SalesBasketController) Create() {
	var salesBasket model.SalesBasket
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &salesBasket); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Validate required fields
	if salesBasket.UserID <= 0 {
		c.JSONResponse(http.StatusBadRequest, "User ID is required", nil)
		return
	}
	
	// Set current time for sales date if not provided
	if salesBasket.SalesDate == 0 {
		salesBasket.SalesDate = int(time.Now().Unix())
	}
	
	// Check that we have at least one item
	if len(salesBasket.Items) == 0 {
		c.JSONResponse(http.StatusBadRequest, "At least one item is required", nil)
		return
	}
	
	// Calculate total if not provided
	if salesBasket.Total == 0 {
		for _, item := range salesBasket.Items {
			salesBasket.Total += item.TotalAmount
		}
	}
	
	// Create transaction
	tx, err := database.DB.Begin()
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to start transaction: "+err.Error(), nil)
		return
	}
	
	// Save sales basket first
	newSalesBasket, err := c.repo.CreateSalesBasketTx(tx, &salesBasket)
	if err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to create sales basket: "+err.Error(), nil)
		return
	}
	
	// Save each sales item with the new sales basket ID
	var savedItems []model.SalesItem
	for _, item := range salesBasket.Items {
		item.SalesID = newSalesBasket.ID
		newItem, err := c.itemRepo.CreateSalesItemTx(tx, &item)
		if err != nil {
			tx.Rollback()
			c.JSONResponse(http.StatusInternalServerError, "Failed to create sales item: "+err.Error(), nil)
			return
		}
		savedItems = append(savedItems, *newItem)
	}
	
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to commit transaction: "+err.Error(), nil)
		return
	}
	
	// Add the items to the response
	newSalesBasket.Items = savedItems
	
	c.JSONResponse(http.StatusCreated, "Sales basket created successfully", newSalesBasket)
}

// Get retrieves a sales basket by ID
func (c *SalesBasketController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	salesBasket, err := c.repo.GetSalesBasket(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Sales basket not found", nil)
		return
	}
	
	// Get items for this sales basket
	items, err := c.itemRepo.GetSalesItemsBySales(id)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve sales items: "+err.Error(), nil)
		return
	}
	
	salesBasket.Items = items
	
	c.JSONResponse(http.StatusOK, "Sales basket retrieved successfully", salesBasket)
}

// GetAll retrieves all sales baskets
func (c *SalesBasketController) GetAll() {
	// Check for optional filters
	userIDStr := c.GetString("user_id")
	memberIDStr := c.GetString("member_id")
	
	var userID, memberID int
	var err error
	
	if userIDStr != "" {
		userID, err = strconv.Atoi(userIDStr)
		if err != nil {
			c.JSONResponse(http.StatusBadRequest, "Invalid user ID format", nil)
			return
		}
	}
	
	if memberIDStr != "" {
		memberID, err = strconv.Atoi(memberIDStr)
		if err != nil {
			c.JSONResponse(http.StatusBadRequest, "Invalid member ID format", nil)
			return
		}
	}
	
	var salesBaskets []model.SalesBasket
	
	if userIDStr != "" && memberIDStr != "" {
		salesBaskets, err = c.repo.GetSalesBasketsByUserAndMember(userID, memberID)
	} else if userIDStr != "" {
		salesBaskets, err = c.repo.GetSalesBasketsByUser(userID)
	} else if memberIDStr != "" {
		salesBaskets, err = c.repo.GetSalesBasketsByMember(memberID)
	} else {
		salesBaskets, err = c.repo.GetAllSalesBaskets()
	}
	
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve sales baskets: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Sales baskets retrieved successfully", salesBaskets)
}

// Update updates a sales basket
func (c *SalesBasketController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	var salesBasket model.SalesBasket
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &salesBasket); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	salesBasket.ID = id
	
	// Check if sales basket exists
	_, err = c.repo.GetSalesBasket(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Sales basket not found", nil)
		return
	}
	
	// Update sales basket
	updatedSalesBasket, err := c.repo.UpdateSalesBasket(&salesBasket)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to update sales basket: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Sales basket updated successfully", updatedSalesBasket)
}

// Delete deletes a sales basket
func (c *SalesBasketController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Check if sales basket exists
	_, err = c.repo.GetSalesBasket(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Sales basket not found", nil)
		return
	}
	
	// Create transaction
	tx, err := database.DB.Begin()
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to start transaction: "+err.Error(), nil)
		return
	}
	
	// Delete related sales items first
	err = c.itemRepo.DeleteSalesItemsBySalesTx(tx, id)
	if err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to delete sales items: "+err.Error(), nil)
		return
	}
	
	// Delete sales basket
	err = c.repo.DeleteSalesBasketTx(tx, id)
	if err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to delete sales basket: "+err.Error(), nil)
		return
	}
	
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.JSONResponse(http.StatusInternalServerError, "Failed to commit transaction: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Sales basket deleted successfully", nil)
}