package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// currentUserID достаёт id, который положил AuthRequired.
func currentUserID(c *gin.Context) (uint, bool) {
	v, exists := c.Get("userID")
	id, ok := v.(uint)
	if !exists || !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return 0, false
	}
	return id, true
}

func internalError(c *gin.Context, err error) {
	log.Println("internal error:", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
}
