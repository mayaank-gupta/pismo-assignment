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
	router := gin.Default()
	routes.RegisterRoutes(router)
	log.Info(fmt.Sprintf("Starting server on port %s", port))
	router.Run(":" + port)
}
