package database

import (
	"errors"

	"github.com/XineAurora/fio-statistics/intrernal/entities"
	"gorm.io/gorm"
)

type PGFIORepository struct {
	db *gorm.DB
}

func (r *PGFIORepository) CreateFIO(fio entities.FIO) (entities.FIO, error) {
	if err := r.db.Create(&fio).Error; err != nil {
		return entities.FIO{}, err
	}
	return fio, nil
}

func (r *PGFIORepository) GetFIOs(filter FIOFilter, page Pagination) []entities.FIO {

	return nil
}

func (r *PGFIORepository) UpdateFIO(fio entities.FIO) error {
	if fio.ID == 0 {
		return errors.New("id is required")
	}
	if err := r.db.Save(&fio).Error; err != nil {
		return err
	}
	return nil
}

func (r *PGFIORepository) DeleteFIO(id uint) error {
	if err := r.db.Delete(&entities.FIO{}, id).Error; err != nil {
		return err
	}
	return nil
}
