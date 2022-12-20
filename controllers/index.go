package controllers

import (
	"app/models"
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateShortLink(c *gin.Context) {
	var body struct {
		Url   string
		Short string
	}

	c.Bind(&body)

	newShortLink := models.ShortLink{OriginalUrl: body.Url, ShortLink: body.Short}

	res := utils.DB.Create(&newShortLink)

	if res.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, newShortLink)
}
