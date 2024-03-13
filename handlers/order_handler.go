package handlers

import (
	"net/http"
	"time"

	"restapigo/database"
	"restapigo/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var request struct {
		OrderedAt    time.Time `json:"orderedAt"`
		CustomerName string    `json:"customerName"`
		Items        []struct {
			Code        string `json:"itemCode"`
			Description string `json:"description"`
			Quantity    int    `json:"quantity"`
		} `json:"items"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		OrderedAt:    request.OrderedAt,
		CustomerName: request.CustomerName,
		Items:        []models.Item{},
	}
	for _, item := range request.Items {
		order.Items = append(order.Items, models.Item{
			Code:        item.Code,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	db := database.DB
	db.Create(&order)

	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	db := database.DB
	db.Preload("Items").Find(&orders)

	c.JSON(http.StatusOK, orders)
}

func UpdateOrder(c *gin.Context) {
	var request struct {
		OrderedAt    time.Time `json:"orderedAt"`
		CustomerName string    `json:"customerName"`
		Items        []struct {
			Code        string `json:"itemCode"`
			Description string `json:"description"`
			Quantity    int    `json:"quantity"`
		} `json:"items"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var order models.Order
	db := database.DB
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.OrderedAt = request.OrderedAt
	order.CustomerName = request.CustomerName

	for i, item := range request.Items {
		if i < len(order.Items) {
			order.Items[i].Code = item.Code
			order.Items[i].Description = item.Description
			order.Items[i].Quantity = item.Quantity
		} else {
			order.Items = append(order.Items, models.Item{
				Code:        item.Code,
				Description: item.Description,
				Quantity:    item.Quantity,
			})
		}
	}

	db.Save(&order)

	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	db := database.DB
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	db.Delete(&order)

	c.JSON(http.StatusOK, gin.H{"message": "Success delete"})
}
