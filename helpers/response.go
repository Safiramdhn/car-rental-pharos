package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type DataPage struct {
	Status      bool        `json:"status"`
	Message     string      `json:"message"`
	Total       int64       `json:"total"`
	Pages       int         `json:"pages"`
	CurrentPage uint        `json:"current_page"`
	Limit       uint        `json:"per_page"`
	Data        interface{} `json:"data"`
}

func BadResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, Response{
		Status:  false,
		Message: message,
	})
}

func GoodResponseWithData(c *gin.Context, message string, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func GoodResponseWithPage(c *gin.Context, message string, statusCode, total, totalPages, page, Limit int, data interface{}) {
	c.JSON(statusCode, DataPage{
		Status:      true,
		Message:     message,
		Total:       int64(total),
		Pages:       totalPages,
		CurrentPage: uint(page),
		Limit:       uint(Limit),
		Data:        data,
	})
}
