package entities

import "errors"

type FIO struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

func (fio FIO) Validate() error {
	if fio.Name == "" {
		return errors.New("require a name")
	}
	if fio.Surname == "" {
		return errors.New("require a surname")
	}
	
	return nil
}
