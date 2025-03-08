package controllers

import (
	"encoding/json"
	"go-pos/model"
	"net/http"
	"strconv"
)

// ItemBatchController handles ItemBatch CRUD operations
type ItemBatchController struct {
	BaseController
}

// Create adds a new item batch
func (c *ItemBatchController) Create() {
	var itemBatch model.ItemBatch
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &itemBatch); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// TODO: Implement repository call to save the item batch
	// For now, mock the response
	itemBatch.ID = 1 // Mocked ID
	
	c.JSONResponse(http.StatusCreated, "Item batch created successfully", itemBatch)
}

// Get retrieves an item batch by ID
func (c *ItemBatchController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to fetch the item batch
	// For now, mock the response
	itemBatch := &model.ItemBatch{
		ID:     id,
		ItemID: 1,
		Qty:    100,
	}
	
	c.JSONResponse(http.StatusOK, "Item batch retrieved successfully", itemBatch)
}

// GetAll retrieves all item batches
func (c *ItemBatchController) GetAll() {
	// TODO: Implement repository call to fetch all item batches
	// For now, mock the response
	itemBatches := []model.ItemBatch{
		{ID: 1, ItemID: 1, Qty: 100},
		{ID: 2, ItemID: 2, Qty: 200},
	}
	
	c.JSONResponse(http.StatusOK, "Item batches retrieved successfully", itemBatches)
}

// Update updates an item batch
func (c *ItemBatchController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	var itemBatch model.ItemBatch
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &itemBatch); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	itemBatch.ID = id
	
	// TODO: Implement repository call to update the item batch
	
	c.JSONResponse(http.StatusOK, "Item batch updated successfully", itemBatch)
}

// Delete deletes an item batch
func (c *ItemBatchController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to delete the item batch
	_ = id // Temporary fix: use the id variable to avoid unused variable error
	
	c.JSONResponse(http.StatusOK, "Item batch deleted successfully", nil)
}