package hclouddns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func New(t string) *HCloudDNS {
	return &HCloudDNS{
		token: t,
	}
}

type HCloudAnswerGetRecord struct {
	Record HCloudRecord `json:"record"`
	Error  HCloudError  `json:"error"`
}

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

type HCloudAnswerGetRecords struct {
	Records []HCloudRecord `json:"records"`
	Meta    HCloudMeta     `json:"meta"`
	Error   HCloudError    `json:"error"`
}

func (d *HCloudDNS) GetRecords(zone_id string) (HCloudAnswerGetRecords, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://dns.hetzner.com/api/v1/records?zone_id=%v", zone_id), nil)
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
