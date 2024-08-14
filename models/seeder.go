package models

import (
	"time"
)

type Seeder struct {
	VersionID int  `json:"version_id" gorm:"type:integer;"`
	IsApplied bool `json:"is_applied" gorm:"type:boolean;"`
	Tstamp    time.Time
}

func (Seeder) TableName() string {
	return "seeder"
}
