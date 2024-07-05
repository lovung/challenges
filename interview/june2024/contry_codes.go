package june2024

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type APIResponse struct {
	Data []struct {
		Name           string   `json:"name"`
		NativeName     string   `json:"nativeName"`
		TopLevelDomain []string `json:"topLevelDomain"`
		Alpha2Code     string   `json:"alpha2Code"`
		NumericCode    string   `json:"numericCode"`
		Alpha3Code     string   `json:"alpha3Code"`
		Currencies     []string `json:"currencies"`
		CallingCodes   []string `json:"callingCodes"`
		Capital        string   `json:"capital"`
	} `json:"data"`
}

type CountryCodeRepository struct {
	endpoint string
}

// https://jsonmock.hackerrank.com/api/countries?name=Puerto%20Rico
func (r *CountryCodeRepository) GetLastCountryCode(ctx context.Context, name string) (string, error) {
	params := url.Values{}
	params.Add("name", name)
	fullUrl := r.endpoint + "?" + params.Encode()
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Printf("Failed to make request: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return "", err
	}

	var countryResp APIResponse
	if err := json.Unmarshal(body, &countryResp); err != nil {
		log.Printf("Failed to unmarshal response: %v", err)
		return "", err
	}

	if len(countryResp.Data) != 1 {
		log.Print("can not find country code")
		return "", errors.New("can not find country code")
	}
	numOfCallingCodes := len(countryResp.Data[0].CallingCodes)
	if numOfCallingCodes < 1 {
		log.Print("can not find calling code")
		return "", errors.New("can not find calling code")
	}
	return countryResp.Data[0].CallingCodes[numOfCallingCodes-1], nil
}

const (
	errrorOutput    = "-1"
	outputFmt       = "+%s %s"
	defaultEndpoint = "https://jsonmock.hackerrank.com/api/countries"
)

func getPhoneNumbers(country string, phoneNumber string) string {
	repo := CountryCodeRepository{
		endpoint: defaultEndpoint,
	}
	ctx := context.Background()
	code, err := repo.GetLastCountryCode(ctx, country)
	if err != nil {
		return errrorOutput
	}
	return fmt.Sprintf(outputFmt, code, phoneNumber)
}
