package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    prom "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    r := gin.Default()
    r.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
    r.GET("/readyz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ready": true}) })
    r.GET("/metrics", gin.WrapH(prom.Handler()))
    r.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "hello from go-api"}) })
    _ = r.Run(":8080")
}
