package main

import (
	"fmt"
	"os"
	"pismo-assignment/routes"
	"pismo-assignment/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	utils.InitLogger()
	log := utils.GetLogger("main")
	port := os.Getenv("PORT")
	r := gin.Default()
	routes.RegisterRoutes(r)
	log.Info(fmt.Sprintf("Starting server on port %s", port))
	r.Run(":" + port)
}
