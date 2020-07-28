package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tidwall/buntdb"
	"log"
)

type App struct {
	Router *gin.Engine
	Db     *buntdb.DB
}

func (a *App) init() {
	a.Router.GET("/api", a.HandleGetTestBingo)
	a.Router.GET("/api/view/:id", a.HandleGetBingoById)
	a.Router.POST("/api/create", a.HandlePostBingo)
}

func (a *App) Run(host string) {
	a.init()
	_ = a.Router.Run(host)
}

func (a *App) HandleGetBingo(c *gin.Context) {
	c.JSON(200, TestData())
}

func (a *App) HandleGetTestBingo(c *gin.Context) {
	c.JSON(200, ApiData())
}

func (a *App) HandleGetBingoById(c *gin.Context) {
	id := c.Param("id")
	var bingo [4][4]string

	err := a.Db.View(func(tx *buntdb.Tx) error {
		value, err := tx.Get(id)
		if err != nil {
			return err
		}
		_ = json.Unmarshal([]byte(value), &bingo)
		return nil
	})
	if err != nil {
		c.Status(500)
	}

	c.JSON(200, bingo)
}

func (a *App) HandlePostBingo(c *gin.Context) {
	var bingo [4][4]string
	_ = c.BindJSON(&bingo)
	payload, _ := json.Marshal(bingo)
	key := uuid.New()

	err := a.Db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key.String(), string(payload), nil)
		return err
	})
	if err != nil {
		c.Status(500)
	}

	log.Println("Received POST data:", bingo)
	c.String(201, key.String())
}
