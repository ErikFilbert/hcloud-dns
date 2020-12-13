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

func searchZoneByID(id string) HCloudZone {
	for _, v := range mockZones {
		if v.ID == id {
			return v
		}
	}
	return HCloudZone{}
}

// API SERVER RESPONCES

func responseGetZone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(HCloudAnswerGetZone{Zone: searchZoneByID(vars["id"])})
}

func responseGetZones(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(HCloudAnswerGetZones{Zones: mockZones})
}

func responseUpdateZone(w http.ResponseWriter, r *http.Request) {
	resultZone := HCloudZone{
		ID:           "ID_UPDATE",
		Name:         "unexpected.ru",
		TTL:          1,
		RecordsCount: 1,
	}

	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resultZone = HCloudZone{}
	}

	answer := HCloudZone{}

	err = json.Unmarshal([]byte(respBody), &answer)
	if err != nil {
		resultZone = HCloudZone{}
	}

	resultZone.Name = answer.Name
	resultZone.TTL = answer.TTL

	json.NewEncoder(w).Encode(HCloudAnswerGetZone{Zone: resultZone})
}

// TESTING RESPONCES

func TestingGetZone(client *HCloudClient, t *testing.T) {
	fmt.Println("Testing GetZone")
	zone, err := client.GetZone("ID_BLINDAGE")
	if err != nil {
		fmt.Print(err)
	}

	if !reflect.DeepEqual(zone.Zone, searchZoneByID("ID_BLINDAGE")) {
		t.Errorf("should be equal, %s", err)
	}

	if reflect.DeepEqual(zone.Zone, searchZoneByID("ID_CBRADIO")) {
		t.Errorf("should not be equal, %s", err)
	}
}

func TestingGetZones(client *HCloudClient, t *testing.T) {
	fmt.Println("Testing GetZones")
	zones, err := client.GetZones(HCloudGetZonesParams{})
	if err != nil {
		fmt.Print(err)
	}

	if !reflect.DeepEqual(zones.Zones, mockZones) {
		t.Errorf("should be equal, %s", err)
	}

	mockZonesUnexpected := []HCloudZone{
		{
			ID:           "ID_BLINDAGE",
			Name:         "blindage.org",
			TTL:          300,
			RecordsCount: 2,
		},
		{
			ID:           "ID_CBRADIO",
			Name:         "cb-radio.ru",
			TTL:          600,
			RecordsCount: 333,
		},
		{
			ID:           "ID_UNEXPECTED",
			Name:         "surprise.ru",
			TTL:          600,
			RecordsCount: 1,
		},
	}

	if reflect.DeepEqual(zones.Zones, mockZonesUnexpected) {
		t.Errorf("should not be equal, %s", err)
	}
}

func TestingUpdateZone(client *HCloudClient, t *testing.T) {
	fmt.Println("Testing UpdateZone")
	updateData := HCloudZone{
		ID:   "ID_UPDATE",
		Name: "blindage.org",
		TTL:  600,
	}

	expectedZone := HCloudZone{
		ID:           "ID_UPDATE",
		Name:         "blindage.org",
		TTL:          600,
		RecordsCount: 1,
	}

	updatedZone, err := client.UpdateZone(updateData)
	if err != nil {
		fmt.Print(err)
	}

	if !reflect.DeepEqual(updatedZone.Zone, expectedZone) {
		t.Errorf("should be equal, %s", err)
	}
}
