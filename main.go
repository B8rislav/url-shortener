package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/B8rislav/url-shortener/database"
	"github.com/B8rislav/url-shortener/handler"
	"github.com/B8rislav/url-shortener/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	godotenv.Load()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Here is URL shortener",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitStore(ctx)
	database.NewPool(ctx)

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Didn't run an app due to error: %v", err.Error()))
	}
}
