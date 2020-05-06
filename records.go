package hclouddns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// New instance
func New(t string) *HCloudDNS {
	return &HCloudDNS{
		token: t,
	}
}

// GetRecord retrieve one single record by ID
func (d *HCloudDNS) GetRecord(id string) (HCloudAnswerGetRecord, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/records/%v", id), nil)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	req.Header.Add("Auth-API-Token", d.token)

	resp, err := client.Do(req)

	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	answer := HCloudAnswerGetRecord{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	return answer, nil

}

// GetRecords retrieve all records of user
func (d *HCloudDNS) GetRecords(zoneID string) (HCloudAnswerGetRecords, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/records?zone_id=%v", zoneID), nil)
	if err != nil {
		return HCloudAnswerGetRecords{}, err
	}

	req.Header.Add("Auth-API-Token", d.token)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		return HCloudAnswerGetRecords{}, parseFormErr
	}

	resp, err := client.Do(req)

	if err != nil {
		return HCloudAnswerGetRecords{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetRecords{}, err
	}

	answer := HCloudAnswerGetRecords{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetRecords{}, err
	}

	return answer, nil
}

// UpdateRecord makes update of single record by ID
func (d *HCloudDNS) UpdateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error) {

	jsonRecordString, err := json.Marshal(record)
	body := bytes.NewBuffer(jsonRecordString)
	log.Println(string(jsonRecordString))

	client := &http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://dns.hetzner.com/api/v1/records/%v", record.ID), body)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", d.token)

	resp, err := client.Do(req)

	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}
	log.Println(string(respBody), resp.StatusCode)

	answer := HCloudAnswerGetRecord{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	return answer, nil

}

// DeleteRecord makes update of single record by ID
func (d *HCloudDNS) DeleteRecord(ID string) (int, error) {

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://dns.hetzner.com/api/v1/records/%v", ID), nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("Auth-API-Token", d.token)

	resp, err := client.Do(req)

	if err != nil {
		return resp.StatusCode, err
	}

	if resp.StatusCode != 200 {
		return resp.StatusCode, fmt.Errorf("Status code is not 200")
	}

	return resp.StatusCode, nil

}

// CreateRecord makes update of single record by ID
func (d *HCloudDNS) CreateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error) {

	jsonRecordString, err := json.Marshal(record)
	body := bytes.NewBuffer(jsonRecordString)
	log.Println(string(jsonRecordString))

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://dns.hetzner.com/api/v1/records/%v", record.ID), body)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", d.token)

	resp, err := client.Do(req)

	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}
	log.Println(string(respBody), resp.StatusCode)

	answer := HCloudAnswerGetRecord{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}

	return answer, nil

}

// CreateRecordBulk makes update of single record by ID
func (d *HCloudDNS) CreateRecordBulk(record []HCloudRecord) (HCloudAnswerCreateRecords, error) {

	jsonRecordString, err := json.Marshal(record)
	body := bytes.NewBuffer(jsonRecordString)
	log.Println(string(jsonRecordString))

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://dns.hetzner.com/api/v1/api/v1/records/bulk", body)
	if err != nil {
		return HCloudAnswerCreateRecords{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", d.token)

	resp, err := client.Do(req)

	if err != nil {
		return HCloudAnswerCreateRecords{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerCreateRecords{}, err
	}
	log.Println(string(respBody), resp.StatusCode)

	answer := HCloudAnswerCreateRecords{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerCreateRecords{}, err
	}

	return answer, nil

}
