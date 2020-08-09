package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	router  *gin.Engine
	handler *BingoHandler
}

func NewBingoApp(router *gin.Engine, handler *BingoHandler) *App {
	return &App{
		router:  router,
		handler: handler,
	}
}

func (a *App) init() {
	a.router.GET("/api", a.handler.HandleGetTestBingo)
	a.router.GET("/api/view/:id", a.handler.HandleGetBingoById)
	a.router.GET("/api/stats", a.handler.HandleGetStatistics)
	a.router.GET("/api/search/:query", a.handler.HandleSearch)
	a.router.GET("/api/index", a.handler.HandleCreateIndex)
	a.router.POST("/api/create", a.handler.HandlePostBingo)
}

func (a *App) Run(host string) {
	a.init()
	log.Fatal(a.router.Run(host))
}
