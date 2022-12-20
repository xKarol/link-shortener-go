package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/ping", controllers.Hello)
}
