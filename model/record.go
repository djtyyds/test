package model

import "time"

type Record struct {
	Name       string    `json:"name"`
	Money      int       `json:"money"`
	RecordTime time.Time `json:"trade_time"`
	Txt        string    `json:"txt"`
}
