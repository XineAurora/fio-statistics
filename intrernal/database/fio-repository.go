package database

import (
	"github.com/XineAurora/fio-statistics/intrernal/entities"
)

type FIORepository interface {
	CreateFIO(fio entities.FIO) (entities.FIO, error)
	GetFIOs(filter FIOFilter, page Pagination) ([]entities.FIO, error)
	UpdateFIO(fio entities.FIO) error
	DeleteFIO(id uint) error
}

type FIOFilter struct {
	Name              string
	Surname           string
	Patronymic        string
	WithoutPatronymic bool
	LowerAge          int
	UpperAge          int
	Gender            string
	Nationality       string
}

type Pagination struct {
	Page   int
	OnPage int
}
