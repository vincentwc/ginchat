package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
