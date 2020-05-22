package modelgorm

import (
	"time"
)

type Home struct {
	Title    string
	Birthday time.Time
	Email    string `gorm:"type:varchar(100);unique_index"`
	Num      int    `gorm:"AUTO_INCREMENT`
}
