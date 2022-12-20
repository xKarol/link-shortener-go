package main

import (
	"app/models"
	"app/utils"
)

func init() {
	utils.LoadEnv()
	utils.ConnectToDB()
}

func main() {
	utils.DB.AutoMigrate(&models.ShortLink{})
}
