package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type Join struct {
  FirstName string `json:"firstName" binding:"required"`
  LastName  string `json:"lastName" binding:"required"`
  Email     string `json:"email" binding:"required"`
  HeroName  string `json:"heroName" binding:"required"`
  HeroClass string `json:"heroClass" binding:"required"`
  Twitter   string `json:"twitter" binding:"required"`
}

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
  router.GET("/hero/:name", api.heroGet)
  router.GET("/hero/:name/activate", api.heroActivate)
  router.GET("/hero/:name/events", api.heroEvents)
  router.GET("/exit", api.exit)
  router.Run(":8080")
}

func (api *API) heroList(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"heroes": api.game.heroes, "count": len(api.game.heroes)})
}

func (api *API) heroPost(c *gin.Context) {
  token := c.Request.Header.Get("X-Auth-Token")
  if len(token) == 0 {
    c.String(http.StatusUnauthorized, "No auth token is present")
    return
  }

  var json Join
  if err := c.BindJSON(&json); err != nil {
    c.String(http.StatusBadRequest, "Invalid request body")
    return
  }

  req := JoinRequest{TokenRequest: TokenRequest{GameRequest: GameRequest{Response: make(chan GameResponse)}, token: token},
    firstName: json.FirstName,
    lastName:  json.LastName,
    email:     json.Email,
    twitter:   json.Twitter,
    heroName:  json.HeroName,
    heroClass: json.HeroClass,
  }
  api.game.joinChan <- req
  res := <-req.Response
  if res.success {
    c.String(http.StatusOK, res.message)
  } else {
    c.String(http.StatusBadRequest, res.message)
  }
}

func (api *API) heroGet(c *gin.Context) {
  name := c.Param("name")

  hero, err := api.game.getHero(name)
  if err != nil {
    c.String(http.StatusNotFound, "Hero not found")
    return
  }

  hero.TTL = hero.getTTL()
  hero.Equipment.Total = hero.getTotalItems()
  c.JSON(http.StatusOK, hero)
}

func (api *API) heroActivate(c *gin.Context) {
  name := c.Param("name")
  token := c.Request.Header.Get("X-Auth-Token")
  if len(token) == 0 {
    c.String(http.StatusUnauthorized, "No auth token is present")
    return
  }

  req := ActivateHeroRequest{TokenRequest: TokenRequest{GameRequest: GameRequest{Response: make(chan GameResponse)}, token: token}, name: name}
  api.game.activateHeroChan <- req
  res := <-req.Response
  if res.success {
    c.String(http.StatusOK, res.message)
  } else {
    c.String(http.StatusBadRequest, res.message)
  }
}

func (api *API) heroEvents(c *gin.Context) {
  name := c.Param("name")
  token := c.Request.Header.Get("X-Auth-Token")
  if len(token) == 0 {
    c.String(http.StatusUnauthorized, "No auth token is present")
    return
  }

  hero, err := api.game.getHero(name)
  if err != nil {
    c.String(http.StatusNotFound, "Hero not found")
    return
  }

  events, err := api.game.GetEventsForHeroFromDB(hero.id)
  if err != nil {
    c.String(http.StatusNotFound, "Could not get Hero events")
    return
  }

  c.JSON(http.StatusOK, events)
}

func (api *API) exit(c *gin.Context) {
  close(api.game.exitChan)
  c.String(http.StatusOK, "Exit")
}
