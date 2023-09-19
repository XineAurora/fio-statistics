package entities

import (
	"encoding/json"
	"errors"
)

type FIO struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

func NewFIO(rawJson []byte) (FIO, error) {
	var fio FIO
	err := json.Unmarshal(rawJson, &fio)
	if err != nil {
		return fio, err
	}
	return fio, nil
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
