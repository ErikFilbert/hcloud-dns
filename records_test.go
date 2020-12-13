package hclouddns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func searchRecordByID(id string) HCloudRecord {
	for _, v := range mockRecordsBlindage {
		if v.ID == id {
			return v
		}
	}
	return HCloudRecord{}
}

// API SERVER RESPONCES

func responseGetRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(HCloudAnswerGetRecord{Record: searchRecordByID(vars["id"])})
}

func responseGetRecords(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(HCloudAnswerGetRecords{Records: mockRecordsBlindage})
}

func responseUpdateRecord(w http.ResponseWriter, r *http.Request) {
	resultRecord := HCloudRecord{
		ID:    "ID_UPDATE",
		Name:  "blindage.org",
		Value: "1.1.1.1",
		TTL:   1,
	}

	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resultRecord = HCloudRecord{}
	}
	fmt.Println("RESP", respBody)

	answer := HCloudRecord{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		resultRecord = HCloudRecord{}
	}

	resultRecord.TTL = answer.TTL
	resultRecord.Value = answer.Value

	json.NewEncoder(w).Encode(HCloudAnswerGetRecord{Record: resultRecord})
}

// TESTING RESPONCES

func TestingGetRecord(client *HCloudClient, t *testing.T) {
	fmt.Println("Testing GetRecord")
	record, err := client.GetRecord("rec1")
	if err != nil {
		fmt.Print(err)
	}

	if !reflect.DeepEqual(record.Record, searchRecordByID("rec1")) {
		t.Errorf("should be equal, %s", err)
	}

	if reflect.DeepEqual(record.Record, searchRecordByID("rec2")) {
		t.Errorf("should not be equal, %s", err)
	}
}

func TestingGetRecords(client *HCloudClient, t *testing.T) {
	records, err := client.GetRecords(HCloudGetRecordsParams{})
	if err != nil {
		fmt.Print(err)
	}

	if !reflect.DeepEqual(records.Records, mockRecordsBlindage) {
		t.Errorf("should be equal, %s", err)
	}

	mockRecordsBlindageUnexpected := []HCloudRecord{
		{
			RecordType: A,
			ID:         "rec1",
			ZoneID:     "ID_BLINDAGE",
			Name:       "blindage.org",
			Value:      "127.0.0.1",
			TTL:        300,
		},
		{
			RecordType: A,
			ID:         "rec2",
			ZoneID:     "ID_BLINDAGE",
			Name:       "www",
			Value:      "127.0.0.1",
			TTL:        300,
		},
		{
			RecordType: A,
			ID:         "rec3",
			ZoneID:     "ID_BLINDAGE",
			Name:       "git",
			Value:      "127.0.0.1",
			TTL:        300,
		},
	}

	if reflect.DeepEqual(records.Records, mockRecordsBlindageUnexpected) {
		t.Errorf("should not be equal, %s", err)
	}
}

func TestingUpdateRecord(client *HCloudClient, t *testing.T) {
	fmt.Println("Testing UpdateRecord")
	updateData := HCloudRecord{
		ID:    "ID_UPDATE",
		Name:  "blindage.org",
		Value: "127.0.0.1",
		TTL:   600,
	}

	expectedRecord := HCloudRecord{
		ID:    "ID_UPDATE",
		Name:  "blindage.org",
		Value: "127.0.0.1",
		TTL:   600,
	}

	updatedRecord, err := client.UpdateRecord(updateData)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("UPD", updatedRecord.Record)

	if !reflect.DeepEqual(updatedRecord.Record, expectedRecord) {
		t.Errorf("should be equal, %s", err)
	}
}
