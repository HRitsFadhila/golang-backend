package main

import (
	"github.com/HRitsFadhila/golang-backend/config"
	"github.com/HRitsFadhila/golang-backend/database"
	"github.com/HRitsFadhila/golang-backend/routes"
)

func main(){
	config.LoadEnv()
	database.InitDB()
	
	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}