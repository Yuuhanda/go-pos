package controllers

import (
	"encoding/json"
	"go-pos/model"
	"net/http"
	"strconv"
)

// CategoryController handles Category CRUD operations
type CategoryController struct {
	BaseController
}

// Create adds a new category
func (c *CategoryController) Create() {
	var category model.Category
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &category); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// TODO: Implement repository call to save the category
	// For now, mock the response
	category.ID = 1 // Mocked ID
	
	c.JSONResponse(http.StatusCreated, "Category created successfully", category)
}

// Get retrieves a category by ID
func (c *CategoryController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to fetch the category
	// For now, mock the response
	category := &model.Category{
		ID:   id,
		Name: "Sample Category",
	}
	
	c.JSONResponse(http.StatusOK, "Category retrieved successfully", category)
}

// GetAll retrieves all categories
func (c *CategoryController) GetAll() {
	// TODO: Implement repository call to fetch all categories
	// For now, mock the response
	categories := []model.Category{
		{ID: 1, Name: "Category 1"},
		{ID: 2, Name: "Category 2"},
	}
	
	c.JSONResponse(http.StatusOK, "Categories retrieved successfully", categories)
}

// Update updates a category
func (c *CategoryController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	var category model.Category
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &category); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	category.ID = id
	
	// TODO: Implement repository call to update the category
	
	c.JSONResponse(http.StatusOK, "Category updated successfully", category)
}

// Delete deletes a category
func (c *CategoryController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to delete the category
	_ = id // Temporary fix: use the id variable to avoid unused variable error
	
	c.JSONResponse(http.StatusOK, "Category deleted successfully", nil)
}