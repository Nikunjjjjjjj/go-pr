package handlers

import (
	"context"
	"net/http"
	"time"
	"url-shortner/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type URLHandler struct {
	rdb *redis.Client
}

func NewURLHandler(rdb *redis.Client) *URLHandler {
	return &URLHandler{rdb}
}

func (h *URLHandler) ShortenUrl(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}
	ctx := context.Background()

	id, err := h.rdb.Incr(ctx, "url:counter").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "redis erroe",
		})
		return
	}

	shortCode := utils.Base62Encode(id)

	key := "url:" + shortCode
	err = h.rdb.Set(ctx, key, req.URL, 7*24*time.Hour).Err()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "redis store error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"short_url": "http://localhost:8080/" + shortCode,
	})
}

func (h *URLHandler) ResolveURL(c *gin.Context) {
	code := c.Param("code")
	ctx := context.Background()

	key := "url:" + code
	originalURL, err := h.rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		c.JSON(404, gin.H{"error": "URL not found or expired"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": "redis error"})
		return
	}

	h.rdb.Incr(ctx, "hits"+code)

	c.Redirect(302, originalURL)
}
