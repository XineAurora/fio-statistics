package entities

import "github.com/XineAurora/fio-statistics/intrernal/enricher"

type FIOEnriched struct {
	FIO
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

func EnrichFIO(fio FIO, enricher enricher.Enricher) (FIOEnriched, error) {
	age, err := enricher.GetAge(fio.Name)
	if err != nil {
		return FIOEnriched{}, err
	}
	gender, err := enricher.GetGender(fio.Name)
	if err != nil {
		return FIOEnriched{}, err
	}
	nation, err := enricher.GetNationality(fio.Name)
	if err != nil {
		return FIOEnriched{}, err
	}
	return FIOEnriched{FIO: fio, Age: age, Gender: gender, Nationality: nation}, nil
}
