package main

import (
	"log"
	"os"
	"sync"

	"github.com/aliirsyaadn/neobank-technical-test/app"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	if os.Getenv("ENV") == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	apiServer := app.InitServer()

	var wg sync.WaitGroup
	wg.Add(1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		defer wg.Done()
		err := apiServer.Run(":" + port)
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
