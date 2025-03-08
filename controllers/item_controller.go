package controllers

import (
	"encoding/json"
	"go-pos/model"
	"go-pos/repository"
	"net/http"
	"strconv"
)

// ItemController handles Item CRUD operations
type ItemController struct {
	BaseController
	repo *repository.ItemRepository
}

// Prepare initializes the controller
func (c *ItemController) Prepare() {
	// Initialize the repository
	c.repo = repository.NewItemRepository()
}

// Create adds a new item
func (c *ItemController) Create() {
	var item model.Item
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &item); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Save the item to database
	newItem, err := c.repo.CreateItem(&item)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to create item: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusCreated, "Item created successfully", newItem)
}

// Get retrieves an item by ID
func (c *ItemController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	item, err := c.repo.GetItem(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Item not found", nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Item retrieved successfully", item)
}

// GetAll retrieves all items
func (c *ItemController) GetAll() {
	// Check for optional category filter
	categoryIDStr := c.GetString("category_id")

	var items []model.Item
	var err error

	if categoryIDStr != "" {
		var categoryID int
		categoryID, err = strconv.Atoi(categoryIDStr)  // Use the existing err variable
		if err != nil {
			c.JSONResponse(http.StatusBadRequest, "Invalid category ID format", nil)
			return
		}
		
		items, err = c.repo.GetItemsByCategory(categoryID)
	} else {
		items, err = c.repo.GetAllItems()
	}
	
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve items: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Items retrieved successfully", items)
}

// Update updates an item
func (c *ItemController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	var item model.Item
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &item); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	item.ID = id
	
	// Check if item exists
	_, err = c.repo.GetItem(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Item not found", nil)
		return
	}
	
	// Update the item
	updatedItem, err := c.repo.UpdateItem(&item)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to update item: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Item updated successfully", updatedItem)
}

// Delete deletes an item
func (c *ItemController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Check if item exists
	_, err = c.repo.GetItem(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Item not found", nil)
		return
	}
	
	// Delete the item
	err = c.repo.DeleteItem(id)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to delete item: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Item deleted successfully", nil)
}