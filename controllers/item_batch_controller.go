package controllers

import (
	"encoding/json"
	"go-pos/model"
	"go-pos/repository"
	"net/http"
	"strconv"
	"time"
)

// ItemBatchController handles ItemBatch CRUD operations
type ItemBatchController struct {
	BaseController
	repo *repository.ItemBatchRepository
}

// Prepare initializes the controller
func (c *ItemBatchController) Prepare() {
	// Initialize the repository
	c.repo = repository.NewItemBatchRepository()
}

// Create adds a new item batch
func (c *ItemBatchController) Create() {
	var itemBatch model.ItemBatch
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &itemBatch); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Validate required fields
	if itemBatch.ItemID <= 0 {
		c.JSONResponse(http.StatusBadRequest, "Item ID is required", nil)
		return
	}
	
	if itemBatch.Qty <= 0 {
		c.JSONResponse(http.StatusBadRequest, "Quantity must be greater than zero", nil)
		return
	}
	
	// Set default values if not provided
	if itemBatch.DateIn.IsZero() {
		itemBatch.DateIn = time.Now()
	}
	
	// Save the item batch to database
	newItemBatch, err := c.repo.CreateItemBatch(&itemBatch)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to create item batch: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusCreated, "Item batch created successfully", newItemBatch)
}

// Get retrieves an item batch by ID
func (c *ItemBatchController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	itemBatch, err := c.repo.GetItemBatch(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Item batch not found", nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Item batch retrieved successfully", itemBatch)
}

// GetAll retrieves all item batches
func (c *ItemBatchController) GetAll() {
	// Check for optional item filter
	itemIDStr := c.GetString("item_id")
	var itemBatches []model.ItemBatch
	var err error
	
	if itemIDStr != "" {
		var itemID int
		itemID, err = strconv.Atoi(itemIDStr)
		if err != nil {
			c.JSONResponse(http.StatusBadRequest, "Invalid item ID format", nil)
			return
		}
		
		itemBatches, err = c.repo.GetItemBatchesByItem(itemID)
	} else {
		itemBatches, err = c.repo.GetAllItemBatches()
	}
	
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve item batches: "+err.Error(), nil)
		return
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
	
	// Check if item batch exists
	_, err = c.repo.GetItemBatch(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Item batch not found", nil)
		return
	}
	
	// Update item batch
	updatedItemBatch, err := c.repo.UpdateItemBatch(&itemBatch)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to update item batch: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Item batch updated successfully", updatedItemBatch)
}

// Delete deletes an item batch
func (c *ItemBatchController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Check if item batch exists
	_, err = c.repo.GetItemBatch(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Item batch not found", nil)
		return
	}
	
	// Delete item batch
	err = c.repo.DeleteItemBatch(id)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to delete item batch: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Item batch deleted successfully", nil)
}