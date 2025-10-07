package handler

import (
	"net/http"
	"strconv"

	"go-currency/models"
	"go-currency/service"

	"github.com/gin-gonic/gin"
)

func ConvertHandler(c *gin.Context) {
	var req models.ConvertRequest

	// Для GET запроса - парсим из query параметров
	if c.Request.Method == "GET" {
		amountStr := c.Query("amount")
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid amount parameter",
			})
			return
		}

		req = models.ConvertRequest{
			From:   c.Query("from"),
			To:     c.Query("to"),
			Amount: amount,
		}
	} else {
		// Для POST запроса - парсим из JSON body
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "error parsing JSON request",
			})
			return
		}
	}

	// Валидация параметров
	if req.From == "" || req.To == "" || req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing or invalid parameters: from, to, amount",
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
