package main

import (
	"app/routes"
	"app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	utils.ConnectToDB()

	r := gin.Default()

	routes.Init(r)

	r.Run()
}
