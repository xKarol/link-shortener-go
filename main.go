package main

import (
	"app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Init(r)

	r.Run()
}
