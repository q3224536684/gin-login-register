package models

import "time"

type Token struct {
	Token   string    `json:"token"`
	Expired time.Time `json:"expired"`
}
