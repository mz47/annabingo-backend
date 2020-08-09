package app

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/tidwall/buntdb"
	"log"
	"math/rand"
)

type BingoService struct {
	db *buntdb.DB
}

func NewBingoService(db *buntdb.DB) *BingoService {
	return &BingoService{db}
}

func (s *BingoService) GetBingoById(id string) (*Bingo, error) {
	var bingo Bingo
	err := s.db.View(func(tx *buntdb.Tx) error {
		value, err := tx.Get(id)
		if err != nil {
			return err
		}
		_ = json.Unmarshal([]byte(value), &bingo)
		return nil
	})
	if err != nil {
		return nil, err
	}
	log.Println("Found Bingo", bingo, "for uuid", id)
	return &bingo, nil
}

func (s *BingoService) SaveBingo(matrix Bingo) (string, error) {
	payload, _ := json.Marshal(matrix)
	key := uuid.New()
	err := s.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key.String(), string(payload), nil)
		return err
	})
	if err != nil {
		return "", err
	}
	log.Println("Saved", matrix, "with uuid", key.String())
	return key.String(), nil
}

func (s *BingoService) Count() (int, error) {
	var count int
	err := s.db.View(func(tx *buntdb.Tx) error {
		length, err := tx.Len()
		count = length
		return err
	})
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (s *BingoService) Shuffle(bingo [4][4]string) *[4][4]string {
	for i := len(bingo) - 1; i > 0; i-- {
		for j := len(bingo[i]) - 1; j > 0; j-- {
			m := rand.Intn(i + 1)
			n := rand.Intn(j + 1)
			temp := bingo[i][j]
			bingo[i][j] = bingo[m][n]
			bingo[m][n] = temp
		}
	}
	return &bingo
}
