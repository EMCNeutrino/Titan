package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

// StartAPI starts the API
func StartAPI() {
  router := gin.Default()
  router.GET("/hero", heroList)
  router.POST("/hero", heroPost)
  router.GET("/hero/:id", heroGet)
  router.Run(":8080")
}

func heroList(c *gin.Context) {
  c.String(http.StatusOK, "Hello List")
}

func heroPost(c *gin.Context) {
  c.String(http.StatusOK, "Hello Post")
}

func heroGet(c *gin.Context) {
  id := c.Param("id")
  c.String(http.StatusOK, "Hero Get ID: %s", id)
}
