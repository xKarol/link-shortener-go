package controllers

import (
	"app/models"
	"app/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortLink(c *gin.Context) {
	var body struct {
		Url       string `json:"url" binding:"required,url"`
		ShortCode string `json:"shortCode" binding:"required"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	newShortLink := models.ShortLink{OriginalUrl: body.Url, ShortCode: body.ShortCode}

	res := utils.DB.Create(&newShortLink)

	if res.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, newShortLink)
}

func GetShortLink(c *gin.Context) {
	shortCode := c.Param("shortCode")

	var ShortLink models.ShortLink

	err := utils.DB.Where("short_code = ?", shortCode).First(&ShortLink).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found."})
		return
	}

	c.Redirect(http.StatusFound, ShortLink.OriginalUrl)
}
