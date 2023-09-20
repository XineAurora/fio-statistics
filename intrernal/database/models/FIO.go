package models

import "github.com/XineAurora/fio-statistics/intrernal/entities"

type FIO struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	entities.FIOEnriched
}
