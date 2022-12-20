package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.POST("/add/:url", controllers.CreateShortLink)
}
