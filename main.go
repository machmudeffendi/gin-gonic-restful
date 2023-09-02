package main

import (
	"os"

	"fenx.dev/restfull-gin-gonic/app/router"
	"fenx.dev/restfull-gin-gonic/config"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
