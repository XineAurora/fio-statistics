package database

import (
	"github.com/XineAurora/fio-statistics/intrernal/database/models"
	"gorm.io/gorm"
)

type PGFIORepository struct {
	db *gorm.DB
}

func (r *PGFIORepository) CreateFIO(fio models.FIO) (models.FIO, error) {
	if err := r.db.Create(&fio).Error; err != nil {
		return models.FIO{}, err
	}
	return fio, nil
}

func (r *PGFIORepository) UpdateFIO(fio models.FIO) error {
	if err := r.db.Save(&fio).Error; err != nil {
		return err
	}
	return nil
}

func (r *PGFIORepository) DeleteFIO(id uint) error {
	if err := r.db.Delete(&models.FIO{}, id).Error; err != nil {
		return err
	}
	return nil
}
