package hclouddns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// GetZone retrieve one single record by ID.
// Accepts zone ID string.
// Returns HCloudAnswerGetZone with HCloudZone, HTTPCode and error
func (d *HCloudDNS) GetZone(ID string) (HCloudAnswerGetZone, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones/%v", ID), nil)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	req.Header.Add("Auth-API-Token", d.token)

	resp, err := client.Do(req)

	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	answer := HCloudAnswerGetZone{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetZone{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetZone{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil

}

// GetZones retrieve all zones of user.
// Accepts exact name as string, search name with partial name.
// Returns HCloudAnswerGetZones with array of HCloudZone, Meta, HTTPCode and error.
func (d *HCloudDNS) GetZones(params HCloudGetZonesParams) (HCloudAnswerGetZones, error) {

	v := url.Values{}
	if params.Name != "" {
		v.Add("name", params.Name)
	}
	if params.SearchName != "" {
		v.Add("search_name", params.SearchName)
	}
	if params.Page != "" {
		v.Add("page", params.Page)
	}
	if params.PerPage != "" {
		v.Add("per_page", params.PerPage)
	}

	client := &http.Client{}
	log.Println(fmt.Sprintf("https://dns.hetzner.com/api/v1/zones?%v", v.Encode()))
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones?%v", v.Encode()), nil)
	if err != nil {
		return HCloudAnswerGetZones{}, err
	}

	req.Header.Add("Auth-API-Token", d.token)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		return HCloudAnswerGetZones{}, parseFormErr
	}

	resp, err := client.Do(req)

	if err != nil {
		return HCloudAnswerGetZones{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetZones{}, err
	}

	answer := HCloudAnswerGetZones{}
	log.Println(string(respBody))
	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetZones{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetZones{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}
