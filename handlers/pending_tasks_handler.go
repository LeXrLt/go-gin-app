package handlers

import (
	"go-gin-app/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPendingTasks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	query := "SELECT st.prompt FROM sub_tasks st JOIN tasks t ON st.task_id = t.id WHERE t.user_id = ? AND st.status = 'pending' AND t.expires_at > ?"
	rows, err := db.DB.Query(query, userID, time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var prompts []string
	for rows.Next() {
		var prompt string
		if err := rows.Scan(&prompt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
			return
		}
		prompts = append(prompts, prompt)
	}

	c.JSON(http.StatusOK, gin.H{"prompts": prompts})
}
