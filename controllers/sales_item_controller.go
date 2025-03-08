package controllers

import (
	"encoding/json"
	"go-pos/model"
	"net/http"
	"strconv"
)

// SalesItemController handles SalesItem CRUD operations
type SalesItemController struct {
	BaseController
}

// Create adds a new sales item
func (c *SalesItemController) Create() {
	var salesItem model.SalesItem
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &salesItem); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// TODO: Implement repository call to save the sales item
	// For now, mock the response
	salesItem.ID = 1 // Mocked ID
	
	c.JSONResponse(http.StatusCreated, "Sales item created successfully", salesItem)
}

// Get retrieves a sales item by ID
func (c *SalesItemController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to fetch the sales item
	// For now, mock the response
	salesItem := &model.SalesItem{
		ID:          id,
		SalesID:     1,
		ItemID:      1,
		Qty:         2,
		TotalAmount: 2000,
	}
	
	c.JSONResponse(http.StatusOK, "Sales item retrieved successfully", salesItem)
}

// GetAll retrieves all sales items
func (c *SalesItemController) GetAll() {
	// TODO: Implement repository call to fetch all sales items
	// For now, mock the response
	salesItems := []model.SalesItem{
		{
			ID:          1,
			SalesID:     1,
			ItemID:      1,
			Qty:         2,
			TotalAmount: 2000,
		},
		{
			ID:          2,
			SalesID:     1,
			ItemID:      2,
			Qty:         3,
			TotalAmount: 3000,
		},
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
	
	// TODO: Implement repository call to update the sales item
	
	c.JSONResponse(http.StatusOK, "Sales item updated successfully", salesItem)
}

// Delete deletes a sales item
func (c *SalesItemController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to delete the sales item
	_ = id // Temporarily use the id variable to avoid unused variable error
	
	c.JSONResponse(http.StatusOK, "Sales item deleted successfully", nil)
}
