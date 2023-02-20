package Model

import (
	"time"
)
type UserLog struct {
    ID          uint     `gorm:primaryKey"`
    UserID    uint
    Action    string
	CreatedAt time.Time
    UpdatedAt time.Time
}
