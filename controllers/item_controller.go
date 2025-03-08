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
	repo repository.ItemRepository
}

// Create adds a new item
func (c *ItemController) Create() {
	var item model.Item
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &item); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// TODO: Implement repository call to save the item
	// For now, mock the response
	item.ID = 1 // Mocked ID
	
	c.JSONResponse(http.StatusCreated, "Item created successfully", item)
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
	// TODO: Implement repository call to fetch all items
	// For now, mock the response
	items := []model.Item{
		{ID: 1, CategoryID: 1, Name: "Item 1", Price: 1000},
		{ID: 2, CategoryID: 2, Name: "Item 2", Price: 2000},
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
	
	// TODO: Implement repository call to update the item
	
	c.JSONResponse(http.StatusOK, "Item updated successfully", item)
}

// Delete deletes an item
func (c *ItemController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	//err = c.repo.Delete(id)
	//if err != nil {
	//	c.JSONResponse(http.StatusInternalServerError, "Failed to delete item", nil)
	//	return
	//}

	_ = id //temporary fix
	
	c.JSONResponse(http.StatusOK, "Item deleted successfully", nil)
}