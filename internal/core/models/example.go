package models

import "time"

type Example struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}
