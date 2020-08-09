package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/buntdb"
	"marcel.works/bingo-backend/app"
)

func main() {
	router := gin.New()
	gin.SetMode(gin.DebugMode)
	router.Use(cors.Default())
	db, _ := buntdb.Open("./db/annabingo.db")
	service := app.NewBingoService(db)
	_ = service.CreateIndexOnTitle()
	handler := app.NewBingoHandler(service)
	a := app.NewBingoApp(router, handler)
	a.Run(":8000")
}
