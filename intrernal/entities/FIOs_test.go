package entities_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/XineAurora/fio-statistics/intrernal/entities"
	"github.com/stretchr/testify/assert"
)

type MockEnricher struct{}

func (e MockEnricher) GetAge(name string) (int, error) {
	if name == "" {
		return -1, errors.New("name required")
	}
	return 999, nil
}

func (e MockEnricher) GetGender(name string) (string, error) {
	if name == "" {
		return "", errors.New("name required")
	}
	return "combat helicopter", nil
}

func (e MockEnricher) GetNationality(name string) (string, error) {
	if name == "" {
		return "", errors.New("name required")
	}
	return "USSR", nil
}

func TestEnrichmentOk(t *testing.T) {
	enr := MockEnricher{}

	fio := entities.FIO{Name: "Test", Surname: "Test"}
	fioEnrichedExpected := entities.FIOEnriched{FIO: fio, Age: 999, Gender: "combat helicopter", Nationality: "USSR"}

	fioEnriched, err := entities.EnrichFIO(fio, enr)
	assert.Nil(t, err)
	assert.Equal(t, fioEnrichedExpected, fioEnriched)
}

func TestFIOConstructor(t *testing.T) {
	type test struct {
		Name     string
		RawJson  []byte
		Expected interface{} //FIO or error
	}
	tests := []test{
		{
			"ok",
			[]byte(`{"name":"Joseph","surname":"Jostar","patronymic":"George"}`),
			entities.FIO{Name: "Joseph", Surname: "Jostar", Patronymic: "George"},
		},
		{
			"ok without patronymic",
			[]byte(`{"name":"Joseph","surname":"Jostar"}`),
			entities.FIO{Name: "Joseph", Surname: "Jostar"},
		},
		{
			"missing surname",
			[]byte(`{"name":"Joseph","patronymic":"George"}`),
			entities.FIO{Name: "Joseph", Patronymic: "George"},
		},
		{
			"missing name",
			[]byte(`{"surname":"Jostar","patronymic":"George"}`),
			entities.FIO{Surname: "Jostar", Patronymic: "George"},
		},
		{
			"wrong json format error",
			[]byte(`{"name":"Joseph","surname":"Jostar","patronymic":George"}`),
			&json.SyntaxError{},
		},
	}

	for _, tst := range tests {
		actualFio, err := entities.NewFIO(tst.RawJson)
		if err != nil {
			assert.IsType(t, tst.Expected, err, "Test %s went wrong", tst.Name)
		} else {
			assert.Equal(t, tst.Expected, actualFio, "Test %s went wrong", tst.Name)
		}
	}
}

func TestFIOValid(t *testing.T) {
	type test struct {
		Name     string
		Fio      entities.FIO
		Expected interface{} //FIO or error
	}

	tests := []test{
		{
			Name:     "Valid FIO",
			Fio:      entities.FIO{"Test", "FIO", "with patronymic"},
			Expected: nil,
		},
		{
			Name:     "Valid FIO without patronymic",
			Fio:      entities.FIO{Name: "Test", Surname: "FIO"},
			Expected: nil,
		},
		{
			Name:     "Invalid FIO without Surname",
			Fio:      entities.FIO{Name: "Surnameless", Patronymic: "123"},
			Expected: errors.New("require a surname"),
		},
		{
			Name:     "Invalid FIO without Name",
			Fio:      entities.FIO{Surname: "Nameless", Patronymic: "123456"},
			Expected: errors.New("require a name"),
		},
	}

	for _, tst := range tests {
		actual := tst.Fio.Validate()
		assert.Equal(t, tst.Expected, actual, "Test %s went wrong", tst.Name)
	}
}
