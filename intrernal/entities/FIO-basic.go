package entities

import (
	"encoding/json"
	"errors"

	"github.com/XineAurora/fio-statistics/intrernal/enricher"
)

type BasicFIO struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

func NewFIO(rawJson []byte) (BasicFIO, error) {
	var fio BasicFIO
	err := json.Unmarshal(rawJson, &fio)
	if err != nil {
		return fio, err
	}
	return fio, nil
}

func (fio BasicFIO) Validate() error {
	if fio.Name == "" {
		return errors.New("require a name")
	}
	if fio.Surname == "" {
		return errors.New("require a surname")
	}

	return nil
}

func (fio BasicFIO) EnrichFIO(enricher enricher.Enricher) (FIO, error) {
	if err := fio.Validate(); err != nil {
		return FIO{}, err
	}
	age, err := enricher.GetAge(fio.Name)
	if err != nil {
		return FIO{}, err
	}
	gender, err := enricher.GetGender(fio.Name)
	if err != nil {
		return FIO{}, err
	}
	nation, err := enricher.GetNationality(fio.Name)
	if err != nil {
		return FIO{}, err
	}
	return FIO{
		Name:       fio.Name,
		Surname:    fio.Surname,
		Patronymic: fio.Patronymic,
		Age:        age, Gender: gender,
		Nationality: nation,
	}, nil
}
