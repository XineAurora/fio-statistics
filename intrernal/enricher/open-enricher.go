package enricher

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

const (
	apikeyAddition = "&apikey="
	agify          = "https://api.agify.io/?name="
	genderize      = "https://api.genderize.io/?name="
	nationalize    = "https://api.nationalize.io/?name="
)

type OpenEnricher struct {
	apiKey string
	client http.Client
}

func NewOpenEnricher(client http.Client) *OpenEnricher {
	return &OpenEnricher{apiKey: "", client: client}
}

func NewOpenEnricherWithApiKey(client http.Client, apiKey string) OpenEnricher {
	return OpenEnricher{apiKey: apiKey, client: client}
}

func (enricher OpenEnricher) createRequest(api string, name string) (*http.Request, error) {
	if enricher.apiKey != "" {
		return http.NewRequest("GET", fmt.Sprintf("%s%s%s%s", api, name, apikeyAddition, enricher.apiKey), nil)
	} else {
		return http.NewRequest("GET", fmt.Sprintf("%s%s", api, name), nil)
	}
}

func (enricher OpenEnricher) makeRequest(api string, name string, typ reflect.Type) (interface{}, error) {
	request, err := enricher.createRequest(api, name)
	if err != nil {
		return nil, err
	}
	resp, err := enricher.client.Do(request)
	if err != nil {
		return nil, err
	}

	body := reflect.New(typ).Interface()
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (enricher OpenEnricher) GetAge(name string) (int, error) {
	type Body struct {
		Count int
		Name  string
		Age   int
		Error string
	}
	res, err := enricher.makeRequest(agify, name, reflect.TypeOf(Body{}))
	if err != nil {
		return -1, err
	}
	body := res.(*Body)
	if body.Error != "" {
		return -1, errors.New(body.Error)
	}
	return body.Age, nil
}

func (enricher OpenEnricher) GetGender(name string) (string, error) {
	type Body struct {
		Count       int
		Name        string
		Gender      string
		Probability float32
		Error       string
	}
	res, err := enricher.makeRequest(genderize, name, reflect.TypeOf(Body{}))
	if err != nil {
		return "", err
	}
	body := res.(*Body)
	if body.Error != "" {
		return "", errors.New(body.Error)
	}
	return body.Gender, nil
}

func (enricher OpenEnricher) GetNationality(name string) (string, error) {
	type Country struct {
		CountryId   string `json:"country_id"`
		Probability float32
	}
	type Body struct {
		Count   int
		Name    string
		Country []Country
		Error   string
	}
	res, err := enricher.makeRequest(nationalize, name, reflect.TypeOf(Body{}))
	if err != nil {
		return "", err
	}
	body := res.(*Body)
	if body.Error != "" {
		return "", errors.New(body.Error)
	}
	if len(body.Country) == 0 {
		return "", errors.New("nationalize api responded without countries")
	}
	return body.Country[0].CountryId, nil
}
