package database

import "github.com/XineAurora/fio-statistics/intrernal/database/models"

type FIORepository interface {
	CreateFIO(fio models.FIO) (models.FIO, error)
	// GetFIO(id uint) models.FIO
	UpdateFIO(fio models.FIO) error
	DeleteFIO(id uint) error
}
