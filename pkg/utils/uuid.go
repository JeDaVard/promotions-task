package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

const (
	paramId = "id"
)

func QueryParamId(c *gin.Context) (uuid.UUID, error) {
	id := strings.TrimSpace(c.Params.ByName(paramId))
	return uuid.Parse(id)
}
