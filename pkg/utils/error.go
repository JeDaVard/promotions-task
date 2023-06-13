package utils

import "github.com/gin-gonic/gin"

// NewError example
func NewError(c *gin.Context, code int, err error) {
	c.JSON(status(code), httpError(code, err))
}

func AbortWithError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(status(code), httpError(code, err))
}

func status(code int) int {
	status := code
	// Custom errors are greater than 1000, use 400 as generic HTTP status in that case
	if status > 1000 {
		status = 400
	}

	return status
}

func httpError(code int, err error) HTTPError {
	return HTTPError{
		Code:    code,
		Message: err.Error(),
	}
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
