package entities

import "errors"

type FIO struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"not null"`
	Surname     string `json:"surname" gorm:"not null"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int    `json:"age" gorm:"not null"`
	Gender      string `json:"gender" gorm:"not null"`
	Nationality string `json:"nationality" gorm:"not null"`
}

func (fio FIO) Validate() error {
	if fio.Name == "" {
		return errors.New("require a name")
	}
	if fio.Surname == "" {
		return errors.New("require a surname")
	}
	if fio.Age <= 0 {
		return errors.New("age must be higher than 0")
	}
	if fio.Gender != "male" && fio.Gender != "female" {
		return errors.New("gender must be male or female")
	}
	if fio.Nationality == "" {
		return errors.New("require a nationality")
	}
	return nil
}
