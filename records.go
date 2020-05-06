package hclouddns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// GetRecord retrieve one single record by ID.
// Accepts single ID of record.
// Returns HCloudAnswerGetRecord with HCloudRecord, HTTPCode and error.
func (d *HCloudDNS) GetRecord(ID string) (HCloudAnswerGetRecord, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/records/%v", ID), nil)
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

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetRecord{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil

}

// GetRecords retrieve all records of user.
// Accepts zone_id as string.
// Returns HCloudAnswerGetRecords with array of HCloudRecord, Meta, HTTPCode and error.
func (d *HCloudDNS) GetRecords(params HCloudGetRecordsParams) (HCloudAnswerGetRecords, error) {

	v := url.Values{}
	if params.Page != "" {
		v.Add("page", params.Page)
	}
	if params.PerPage != "" {
		v.Add("per_page", params.PerPage)
	}
	if params.ZoneID != "" {
		v.Add("zone_id", params.ZoneID)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/records?%v", v.Encode()), nil)
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

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetRecords{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil
}

// UpdateRecord makes update of single record by ID.
// Accepts HCloudRecord with fullfilled fields.
// Returns HCloudAnswerGetRecord with HTTP code, HCloudRecord and error.
func (d *HCloudDNS) UpdateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error) {

	jsonRecordString, err := json.Marshal(record)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}
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

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetRecord{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil

}

// DeleteRecord remove record by ID.
// Accepts single ID string.
// Returns HTTP code and error. HTTP code 200 also raise error.
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

// CreateRecord creates new single record.
// Accepts HCloudRecord with record to create, of cource no ID.
// Returns HCloudAnswerGetRecord with HTTPCode, HCloudRecord and error.
func (d *HCloudDNS) CreateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error) {

	jsonRecordString, err := json.Marshal(record)
	if err != nil {
		return HCloudAnswerGetRecord{}, err
	}
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

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerGetRecord{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil

}

// CreateRecordBulk creates many records at once.
// Accepts array of HCloudRecord, converting to json and makes POST to Hetzner.
// Returns HCloudAnswerCreateRecords with HTTPCode, arrays of HCloudRecord with whole list, valid and invalid, error.
func (d *HCloudDNS) CreateRecordBulk(record []HCloudRecord) (HCloudAnswerCreateRecords, error) {

	jsonRecordString, err := json.Marshal(record)
	if err != nil {
		return HCloudAnswerCreateRecords{}, err
	}
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

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerCreateRecords{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil

}

// UpdateRecordBulk updates many records at once.
// Accepts array of HCloudRecord, converting to json and makes PUT to Hetzner.
// Returns HCloudAnswerUpdateRecords with HTTPCode, arrays of HCloudRecord with updated and failed, error.
func (d *HCloudDNS) UpdateRecordBulk(record []HCloudRecord) (HCloudAnswerUpdateRecords, error) {

	jsonRecordString, err := json.Marshal(record)
	if err != nil {
		return HCloudAnswerUpdateRecords{}, err
	}
	body := bytes.NewBuffer(jsonRecordString)
	log.Println(string(jsonRecordString))

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "https://dns.hetzner.com/api/v1/api/v1/records/bulk", body)
	if err != nil {
		return HCloudAnswerUpdateRecords{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-API-Token", d.token)

	resp, err := client.Do(req)

	if err != nil {
		return HCloudAnswerUpdateRecords{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return HCloudAnswerUpdateRecords{}, err
	}
	log.Println(string(respBody), resp.StatusCode)

	answer := HCloudAnswerUpdateRecords{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		return HCloudAnswerUpdateRecords{}, err
	}

	// parse error
	errorResult := HCloudAnswerError{}
	err = json.Unmarshal([]byte(respBody), &errorResult)
	if err != nil {
		//ok, non-standard error, try another form
		errorResultString := HCloudAnswerErrorString{}
		err = json.Unmarshal([]byte(respBody), &errorResultString)
		if err != nil {
			return HCloudAnswerUpdateRecords{}, err
		}
		errorResult.Error.Message = errorResultString.Error
		errorResult.Error.Code = resp.StatusCode
	}
	answer.Error = errorResult.Error

	return answer, nil

}
