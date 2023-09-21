package database

import (
	"errors"

	"github.com/XineAurora/fio-statistics/intrernal/entities"
	"gorm.io/gorm"
)

type DBFIORepository struct {
	db *gorm.DB
}

func NewDBFIORepository(db *gorm.DB) *DBFIORepository {
	return &DBFIORepository{db: db}
}

func (r *DBFIORepository) CreateFIO(fio entities.FIO) (entities.FIO, error) {
	if err := r.db.Create(&fio).Error; err != nil {
		return entities.FIO{}, err
	}
	return fio, nil
}

func (r *DBFIORepository) GetFIOs(filter FIOFilter, page Pagination) ([]entities.FIO, error) {
	var fios []entities.FIO
	db := r.db
	if filter.Name != "" {
		db = db.Where("name LIKE ?", filter.Name+"%")
	}
	if filter.Surname != "" {
		db = db.Where("surname LIKE ?", filter.Surname+"%")
	}
	if filter.Patronymic != "" {
		db = db.Where("patronymic LIKE ?", filter.Patronymic+"%")
	} else if filter.WithoutPatronymic {
		db = db.Where("patronymic IS NULL")
	}
	if filter.LowerAge > 0 {
		db = db.Where("age >= ?", filter.LowerAge)
	}
	if filter.UpperAge > 0 && filter.UpperAge >= filter.LowerAge {
		db = db.Where("age <= ?", filter.UpperAge)
	}
	if filter.Gender != "" {
		db = db.Where("gender = ?", filter.Gender)
	}
	if filter.Nationality != "" {
		db = db.Where("nationality = ?", filter.Nationality)
	}
	if page.OnPage <= 0 {
		page.OnPage = 15
	}
	if page.Page <= 0 {
		page.Page = 0
	}
	db = db.Offset(page.Page * page.OnPage).Limit(page.OnPage)
	if err := db.Find(&fios).Error; err != nil {
		return nil, err
	}

	return fios, nil
}

func (r *DBFIORepository) UpdateFIO(fio entities.FIO) error {
	if fio.ID == 0 {
		return errors.New("id is required")
	}
	if err := r.db.Save(&fio).Error; err != nil {
		return err
	}
	return nil
}

func (r *DBFIORepository) DeleteFIO(id uint) error {
	if err := r.db.Delete(&entities.FIO{}, id).Error; err != nil {
		return err
	}
	return nil
}
