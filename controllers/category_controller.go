package controllers

import (
	"encoding/json"
	"go-pos/model"
	"go-pos/repository"
	"net/http"
	"strconv"
)

// CategoryController handles Category CRUD operations
type CategoryController struct {
	BaseController
	repo *repository.CategoryRepository
}

// Prepare initializes the controller
func (c *CategoryController) Prepare() {
	// Initialize the repository
	c.repo = repository.NewCategoryRepository()
}

// Create adds a new category
func (c *CategoryController) Create() {
	var category model.Category
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &category); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Save the category to database
	newCategory, err := c.repo.CreateCategory(&category)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to create category: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusCreated, "Category created successfully", newCategory)
}

// Get retrieves a category by ID
func (c *CategoryController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	category, err := c.repo.GetCategory(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Category not found", nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Category retrieved successfully", category)
}

// GetAll retrieves all categories
func (c *CategoryController) GetAll() {
	categories, err := c.repo.GetAllCategories()
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to retrieve categories: "+err.Error(), nil)
		return
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
	
	// Check if category exists
	_, err = c.repo.GetCategory(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Category not found", nil)
		return
	}
	
	// Update category
	updatedCategory, err := c.repo.UpdateCategory(&category)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to update category: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Category updated successfully", updatedCategory)
}

// Delete deletes a category
func (c *CategoryController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// Check if category exists
	_, err = c.repo.GetCategory(id)
	if err != nil {
		c.JSONResponse(http.StatusNotFound, "Category not found", nil)
		return
	}
	
	// Check if category is in use by any items
	inUse, err := c.repo.IsCategoryInUse(id)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to check if category is in use: "+err.Error(), nil)
		return
	}
	
	if inUse {
		c.JSONResponse(http.StatusBadRequest, "Cannot delete category: it is being used by one or more items", nil)
		return
	}
	
	// Delete category
	err = c.repo.DeleteCategory(id)
	if err != nil {
		c.JSONResponse(http.StatusInternalServerError, "Failed to delete category: "+err.Error(), nil)
		return
	}
	
	c.JSONResponse(http.StatusOK, "Category deleted successfully", nil)
}