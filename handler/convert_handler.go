package handler

import (
	"net/http"

	"go-currency/models"
	"go-currency/service"

	"github.com/gin-gonic/gin"
)

func ConvertHandler(c *gin.Context) {
	var req models.ConvertRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error params request",
		})
		return
	}

	result, err := service.ConvertCurrency(req.From, req.To, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ConvertResponse{
		Result: result,
	})
}