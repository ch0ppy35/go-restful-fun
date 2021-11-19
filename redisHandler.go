package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Hitcount struct {
	Count int64 `json:"count"`
}

var redisAddress = getEnv("REDIS_HOST", "127.0.0.1")
var ctx = context.Background()

func requestsCounter(c *gin.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress + ":6379",
		Password: "",
		DB:       0,
	})

	count, err := rdb.Do(ctx, "INCR", "count").Result()
	if err != nil {
		message := map[string]string{"message": "Can't connect to Redis!"}
		c.Header("Content-Type", "application/json")
		c.IndentedJSON(http.StatusInternalServerError, message)
	}

	hitcount := Hitcount{Count: count.(int64)}

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, hitcount)
}
