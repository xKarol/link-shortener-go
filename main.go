package main

import (
	"app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	routes.Init(r)

	r.Run()
}
