package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"./query"
)

type app struct {
	QueryCache *query.Cache
	engine     *gin.Engine
}

// newApp returns a new app with a given cache lifetime in seconds.
func newApp(d time.Duration) *app {
	a := &app{QueryCache: query.NewCache(d * time.Second)}
	a.initRoutes()
	return a
}

// run sets up the application's routes and begins listening
func (a *app) initRoutes() {
	a.engine = gin.New()

	a.engine.GET("/query/:address", a.QueryController)
	a.engine.Static("/static", "static/")
	a.engine.StaticFile("/", "static/index.html")
}

func (a *app) run(port string) {
	a.engine.Run(port)
}

func (a *app) QueryController(c *gin.Context) {
	server := c.Param("address")

	ss := strings.Split(server, ":")
	if len(ss) < 2 {
		server = server + ":27015"
	}

	if len(ss) > 2 {
		c.String(http.StatusBadRequest, "Too many :s in request.")
		return
	}

	result, err := a.QueryCache.Get(server)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error querying server: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
