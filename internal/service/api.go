package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type AgeResponse struct {
	Age *int `json:"age"`
}

func (s *Service) fetchAgeFromAPI(name string) (int, error) {
	s.log.Debugf("Fetching age for name: %s", name)

	urlAPI := fmt.Sprintf("%s/?name=%s", s.apiUrl.ApiGetAge, name)

	resp, err := fetchExternalData(urlAPI, s.log)
	if err != nil {
		s.log.Errorf("Error fetching age from API: %s", err)
		return 0, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var ageResp AgeResponse
		err = json.NewDecoder(resp.Body).Decode(&ageResp)
		if err != nil {
			s.log.Errorf("Error decoding age response: %s", err)
			return 0, err
		}

		if ageResp.Age == nil {
			s.log.Debug("No age found in the response")
			return 0, fmt.Errorf("No age found")
		}

		s.log.Debugf("Successfully fetched age %d for name: %s", *ageResp.Age, name)
		return *ageResp.Age, nil

	case http.StatusTooManyRequests:
		s.log.Debug("Received 429 (Too Many Requests) status from the API. Rate limit exceeded.")
		return 0, fmt.Errorf("Rate limit exceeded. Try again later")

	default:
		s.log.Errorf("Received unexpected status code: %d", resp.StatusCode)
		return 0, fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
}

type GenderResponse struct {
	Gender *string `json:"gender"`
}

func (s *Service) fetchGenderFromAPI(name string) (string, error) {
	s.log.Debugf("Fetching gender for name: %s", name)

	urlAPI := fmt.Sprintf("%s/?name=%s", s.apiUrl.ApiGetGender, name)

	resp, err := fetchExternalData(urlAPI, s.log)
	if err != nil {
		s.log.Errorf("Error fetching gender from API: %s", err)
		return "", err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var genderResp GenderResponse
		err = json.NewDecoder(resp.Body).Decode(&genderResp)
		if err != nil {
			s.log.Errorf("Error decoding gender response: %s", err)
			return "", err
		}

		if genderResp.Gender == nil {
			s.log.Debug("No gender found in the response")
			return "", fmt.Errorf("No gender found")
		}

		s.log.Debugf("Successfully fetched gender %s for name: %s", *genderResp.Gender, name)
		return *genderResp.Gender, nil

	case http.StatusTooManyRequests:
		s.log.Debug("Received 429 (Too Many Requests) status from the API. Rate limit exceeded.")
		return "", fmt.Errorf("Rate limit exceeded. Try again later")

	default:
		s.log.Errorf("Received unexpected status code: %d", resp.StatusCode)
		return "", fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

}

type CountryIDResponse struct {
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

func (s *Service) fetchCountryIDFromAPI(name string) (string, error) {
	s.log.Debugf("Fetching country id for name: %s", name)

	urlAPI := fmt.Sprintf("%s/?name=%s", s.apiUrl.ApiGetCountry, name)

	resp, err := fetchExternalData(urlAPI, s.log)
	if err != nil {
		s.log.Errorf("Error fetching country id from API: %s", err)
		return "", err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var countryIDResp CountryIDResponse
		err = json.NewDecoder(resp.Body).Decode(&countryIDResp)
		if err != nil {
			s.log.Errorf("Error decoding country id response: %s", err)
			return "", err
		}

		var countryID string
		if len(countryIDResp.Country) > 0 {
			countryID = countryIDResp.Country[0].CountryID
		}
		if countryID == "" {
			s.log.Debug("No country found in the response")
			return "", fmt.Errorf("No country found")
		}

		s.log.Debugf("Successfully fetched country id %s for name: %s", countryID, name)
		return countryID, nil

	case http.StatusTooManyRequests:
		s.log.Debug("Received 429 (Too Many Requests) status from the API. Rate limit exceeded.")
		return "", fmt.Errorf("Rate limit exceeded. Try again later")

	default:
		s.log.Errorf("Received unexpected status code: %d", resp.StatusCode)
		return "", fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
}

func fetchExternalData(urlAPI string, log *logrus.Logger) (*http.Response, error) {
	log.Debugf("Fetching data from %s", urlAPI)

	req, err := http.NewRequest("GET", urlAPI, nil)
	if err != nil {
		log.Errorf("Error creating HTTP request: %s", err)
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error making HTTP request: %s", err)
		return nil, err
	}

	log.Debugf("Data fetched successfully from %s", urlAPI)
	return resp, nil
}
