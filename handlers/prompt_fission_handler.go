package handlers

import (
	"go-gin-app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PromptFissionRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

func PromptFission(c *gin.Context) {
	var req PromptFissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count := config.Cfg.PromptFissionCount
	if count <= 0 {
		count = 50 // Default to 50 if not configured or invalid
	}

	fissionedPrompts := make([]string, 0, count)
	for i := 0; i < count; i++ {
		fissionedPrompts = append(fissionedPrompts, req.Prompt)
	}

	c.JSON(http.StatusOK, gin.H{"prompts": fissionedPrompts})
}