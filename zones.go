package hclouddns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	answer.HTTPCode = resp.StatusCode
	return answer, nil

}

// GetZones retrieve all zones of user.
// Accepts exact name as string, search name with partial name.
// Returns HCloudAnswerGetZones with array of HCloudZone, Meta, HTTPCode and error.
func (d *HCloudDNS) GetZones(params HCloudGetZonesParams) (HCloudAnswerGetZones, error) {

	jsonRecordString, err := json.Marshal(params)
	if err != nil {
		return HCloudAnswerGetZones{}, err
	}
	body := bytes.NewBuffer(jsonRecordString)
	log.Println(string(jsonRecordString))

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/zones"), body)
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

	answer.HTTPCode = resp.StatusCode
	return answer, nil
}
