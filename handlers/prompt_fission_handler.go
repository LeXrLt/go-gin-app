package handlers

import (
	"go-gin-app/utils"
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

	fissionedPrompts, err := utils.GetOpenAIPromptFission(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"prompts": fissionedPrompts})
}