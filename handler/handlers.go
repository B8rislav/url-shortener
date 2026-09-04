package handler

import (
	"log"
	"net/http"

	"github.com/B8rislav/url-shortener/database"
	"github.com/B8rislav/url-shortener/shortener"
	"github.com/B8rislav/url-shortener/store"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type URLCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(ctx *gin.Context) {
	var creationRequest URLCreationRequest
	if err := ctx.ShouldBindJSON(&creationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	err := store.SaveUrl(ctx, shortUrl, creationRequest.LongUrl, creationRequest.UserId)
	if err != nil {
		return
	}
	database.SaveUrl(ctx, shortUrl, creationRequest.LongUrl)

	host := "http://localhost:8080/"
	ctx.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	var initialUrl string
	initialUrl, err := store.GetUrl(ctx, shortUrl)
	if err != nil {
		if err == redis.Nil {
			log.Println("Not found in redis. Scanning db...")
			initialUrl, err = database.GetUrl(ctx, shortUrl)
			if err != nil {
				if err == pgx.ErrNoRows {
					log.Println("Nothing in db either")
					return
				}
			}
		}
	}
	ctx.Redirect(302, initialUrl)
}
