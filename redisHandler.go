package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Hitcount struct {
	Count int64 `json:"count"`
}

func reddisClient() *redis.Client {
	redisAddress := getEnv("REDIS_HOST", "127.0.0.1")
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress + ":6379",
		Password: "",
		DB:       0,
	})

	return client
}

func requestsCounter(c *gin.Context) {
	reddis := reddisClient()

	count, err := reddis.Do("INCR", "count").Result()
	if err != nil {
		message := map[string]string{"message": "Can't connect to Redis!"}
		c.Header("Content-Type", "application/json")
		c.IndentedJSON(http.StatusInternalServerError, message)
	}

	hitcount := Hitcount{Count: count.(int64)}

	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, hitcount)
}
