package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())

	healthz := r.Group("/healthz")
	{
		healthz.GET("ready", isReady)
		healthz.GET("alive", isAlive)
	}
	r.POST("/echo", echo)
	r.Run(":3000")
}

func echo(c *gin.Context) {
	j := json.NewDecoder(c.Request.Body)
	payload := make(map[string]interface{})
	err := j.Decode(&payload)
	if err != nil {
		c.String(http.StatusBadRequest, "Please provide valid json as payload.")
		return
	}
	c.JSON(200, payload)
}

func isReady(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Ready",
	})
}

func isAlive(c *gin.Context) {
	c.JSON(200, gin.H{
		"alive": true,
	})
}
