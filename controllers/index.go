package controllers

import (
	"app/models"
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortLink(c *gin.Context) {
	url := c.Param("url")

	newShortLink := models.ShortLink{OriginalUrl: url, ShortLink: url}

	res := utils.DB.Create(&newShortLink)

	if res.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, newShortLink)
}
