package models

import (
	"encoding/json"
	"time"
)

type PrizesInput struct {
	IndexPrize json.Number `json:"index_prize" binding:"required,number"`
	PrizeName  string      `json:"prize_name" binding:"required"`
	Qouta      json.Number `json:"qouta" binding:"required,number"`
	DayLimit   json.Number `json:"day_limit" binding:"required,number"`
}

type Prize struct {
	ID         int
	IndexPrize int
	PrizeName  string
	Qouta      int
	DayLimit   int
	CreatedAt  time.Time
}

type PrizeToJson struct {
	ID         int    `json:"id"`
	IndexPrize int    `json:"index_prize"`
	PrizeName  string `json:"prize_name"`
	Qouta      int    `json:"quota"`
	DayLimit   int    `json:"day_limit"`
	// CreatedAt  time.Time `json:"created_at"`
}
