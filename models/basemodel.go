package models

import "time"

type BaseModel struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey;not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime:nano"`
}
