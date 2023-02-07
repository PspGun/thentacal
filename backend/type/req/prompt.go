package req

import (
	"gorm.io/gorm"
)

type Prompt struct {
	gorm.Model
	UserId string `json:"userId"`
	Prompt string `json:"prompt"`
}
