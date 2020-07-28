package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/buntdb"
	"marcel.works/bingo-backend/app"
)

func main() {
	r := gin.New()
	r.Use(cors.Default())
	db, _ := buntdb.Open("annabingo.db")
	a := app.App{
		Router: r,
		Db:     db,
	}
	a.Run(":8000")
}
