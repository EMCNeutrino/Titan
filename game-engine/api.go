package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type API struct {
  game *Game
}

// StartAPI starts the API
func (g *Game) StartAPI() {
  api := &API{
    game: g,
  }
  router := gin.Default()
  router.GET("/hero", api.heroList)
  router.POST("/hero", api.heroPost)
  router.GET("/hero/:id", api.heroGet)
  router.GET("/exit", api.exit)
  router.Run(":8080")
}

func (api *API) heroList(c *gin.Context) {
  c.String(http.StatusOK, "Hello List")
}

func (api *API) heroPost(c *gin.Context) {
  c.String(http.StatusOK, "Hello Post")
}

func (api *API) heroGet(c *gin.Context) {
  id := c.Param("id")
  c.String(http.StatusOK, "Hero Get ID: %s", id)
}

func (api *API) exit(c *gin.Context) {
  close(api.game.exitChan)
  c.String(http.StatusOK, "Exit")
}
