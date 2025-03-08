package controllers

import (
	"encoding/json"
	"go-pos/model"
	"go-pos/repository"
	"net/http"
	"strconv"
)

// SalesItemController handles SalesItem CRUD operations
type SalesItemController struct {
	BaseController
	repo *repository.SalesItemRepository
}

// Prepare initializes the controller
func (c *SalesItemController) Prepare() {
	// Initialize the repository
	c.repo = repository.NewSalesItemRepository()
}

// Create adds a new sales item
func (c *SalesItemController) Create() {
	var salesItem model.SalesItem
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &salesItem); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Validate required fields
	if salesItem.SalesID <= 0 {
		c.JSONResponse(http.StatusBadRequest, "Sales ID is required", nil)
		return
	}
	
	if salesItem.ItemID <= 0 {
		c.JSONResponse(http.StatusBadRequest, "Item ID is required", nil)
		return
	}
	
	if salesItem.Qty <= 0 {
		c.JSONResponse(http.StatusBadRequest, "Quantity must be greater than zero", nil)
		return
	}
	
	// Save the sales item to database
	newSalesItem, err := c.repo.CreateSalesItem(&salesItem)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to create sales item: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusCreated, "Sales item created successfully", newSalesItem)
}

// Get retrieves a sales item by ID
func (c *SalesItemController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	salesItem, err := c.repo.GetSalesItem(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Sales item not found", nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Sales item retrieved successfully", salesItem)
}

// GetAll retrieves all sales items
func (c *SalesItemController) GetAll() {
	// Check for optional sales basket filter
	salesIDStr := c.GetString("sales_id")
	var salesItems []model.SalesItem
	var err error
	
	if salesIDStr != "" {
		var salesID int
		salesID, err = strconv.Atoi(salesIDStr)
		if err != nil {
			c.JSONResponse(http.StatusBadRequest, "Invalid sales ID format", nil)
			return
		}
		
		salesItems, err = c.repo.GetSalesItemsBySales(salesID)
	} else {
		salesItems, err = c.repo.GetAllSalesItems()
	}
	
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve sales items: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Sales items retrieved successfully", salesItems)
}

// Update updates a sales item
func (c *SalesItemController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	var salesItem model.SalesItem
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &salesItem); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	salesItem.ID = id
	
	// Check if sales item exists
	_, err = c.repo.GetSalesItem(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Sales item not found", nil)
		return
	}
	
	// Update sales item
	updatedSalesItem, err := c.repo.UpdateSalesItem(&salesItem)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to update sales item: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Sales item updated successfully", updatedSalesItem)
}

// Delete deletes a sales item
func (c *SalesItemController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Check if sales item exists
	_, err = c.repo.GetSalesItem(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Sales item not found", nil)
		return
	}
	
	// Delete sales item
	err = c.repo.DeleteSalesItem(id)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to delete sales item: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Sales item deleted successfully", nil)
}
