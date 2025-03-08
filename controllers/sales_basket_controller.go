package controllers

import (
	"encoding/json"
	"go-pos/model"
	"net/http"
	"strconv"
	"time"
)

// SalesBasketController handles SalesBasket CRUD operations
type SalesBasketController struct {
	BaseController
}

// Create adds a new sales basket
func (c *SalesBasketController) Create() {
	var salesBasket model.SalesBasket
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &salesBasket); err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	// Set current time for sales date if not provided
	if salesBasket.SalesDate == 0 {
		// Convert current timestamp to int
		salesBasket.SalesDate = int(time.Now().Unix())
	}
	
	// TODO: Implement repository call to save the sales basket
	// For now, mock the response
	salesBasket.ID = 1 // Mocked ID
	
	c.JSONResponse(http.StatusCreated, "Sales basket created successfully", salesBasket)
}

// Get retrieves a sales basket by ID
func (c *SalesBasketController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to fetch the sales basket
	// For now, mock the response
	salesBasket := &model.SalesBasket{
		ID:            id,
		SalesDate:     int(time.Now().Unix()),
		UserID:        1,
		MemberID:      1,
		PaymentMethod: model.PaymentMethodCash,
		Total:         5000,
	}
	
	// Mock some sales items
	salesBasket.Items = []model.SalesItem{
		{
			ID:          1,
			SalesID:     id,
			ItemID:      1,
			Qty:         2,
			TotalAmount: 2000,
		},
		{
			ID:          2,
			SalesID:     id,
			ItemID:      2,
			Qty:         3,
			TotalAmount: 3000,
		},
	}
	
	c.JSONResponse(http.StatusOK, "Sales basket retrieved successfully", salesBasket)
}

// GetAll retrieves all sales baskets
func (c *SalesBasketController) GetAll() {
	// TODO: Implement repository call to fetch all sales baskets
	// For now, mock the response
	salesBaskets := []model.SalesBasket{
		{
			ID:            1,
			SalesDate:     int(time.Now().Unix()),
			UserID:        1,
			MemberID:      1,
			PaymentMethod: model.PaymentMethodCash,
			Total:         5000,
		},
		{
			ID:            2,
			SalesDate:     int(time.Now().Unix() - 86400), // yesterday
			UserID:        2,
			MemberID:      2,
			PaymentMethod: model.PaymentMethodCredit,
			Total:         7500,
		},
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
	
	// TODO: Implement repository call to update the sales basket
	
	c.JSONResponse(http.StatusOK, "Sales basket updated successfully", salesBasket)
}

// Delete deletes a sales basket
func (c *SalesBasketController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSONResponse(http.StatusBadRequest, "Invalid ID format", nil)
		return
	}
	
	// TODO: Implement repository call to delete the sales basket
	_  = id //temporary fix
	c.JSONResponse(http.StatusOK, "Sales basket deleted successfully", nil)
}
