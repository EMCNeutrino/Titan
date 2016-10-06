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
  router.GET("/hero/:id/activate", api.heroActivate)
  router.GET("/exit", api.exit)
  router.Run(":8080")
}

func (api *API) heroList(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"heros": api.game.heros, "count": len(api.game.heros)})
}

func (api *API) heroPost(c *gin.Context) {
  token := c.Request.Header.Get("X-Auth-Token")
  if len(token) == 0 {
    c.String(http.StatusUnauthorized, "No auth token is present")
    return
  }
  api.game.joinChan <- token
  c.String(http.StatusOK, "Hero Post")
}

func (api *API) heroGet(c *gin.Context) {
  id := c.Param("id")
  c.String(http.StatusOK, "Hero Get ID: %s", id)
}

func (api *API) heroActivate(c *gin.Context) {
  id := c.Param("id")
  c.String(http.StatusOK, "Hello Activate ID: %s", id)
}

func (api *API) exit(c *gin.Context) {
  close(api.game.exitChan)
  c.String(http.StatusOK, "Exit")
}
