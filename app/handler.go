package app

import (
	"github.com/gin-gonic/gin"
)

type BingoHandler struct {
	service *BingoService
}

func NewBingoHandler(service *BingoService) *BingoHandler {
	return &BingoHandler{service: service}
}

func (h *BingoHandler) HandleGetBingo(c *gin.Context) {
	c.JSON(200, TestData())
}

func (h *BingoHandler) HandleGetTestBingo(c *gin.Context) {
	bingo := ApiData()
	bingo.Fields = *h.service.Shuffle(bingo.Fields)
	c.JSON(200, bingo)
}

func (h *BingoHandler) HandleGetBingoById(c *gin.Context) {
	id := c.Param("id")
	bingo, err := h.service.GetBingoById(id)
	if err != nil {
		c.Status(500)
	}
	bingo.Fields = *h.service.Shuffle(bingo.Fields)
	c.JSON(200, bingo)
}

func (h *BingoHandler) HandlePostBingo(c *gin.Context) {
	var matrix Bingo
	_ = c.BindJSON(&matrix)
	id, err := h.service.SaveBingo(matrix)
	if err != nil {
		c.Status(500)
	}
	c.String(201, id)
}

func (h *BingoHandler) HandleGetStatistics(c *gin.Context) {
	count, err := h.service.Count()
	if err != nil {
		c.Status(500)
	}
	stats := Stats{Count: count}
	c.JSON(200, stats)
}
