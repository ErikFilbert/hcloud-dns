package main

import (
	"io/ioutil"
	"log"

	hclouddns "github.com/ErikFilbert/hcloud-dns"
)

type mockHCloudClientAdapter interface {
	GetZone(ID string) (hclouddns.HCloudAnswerGetZone, error)
	GetZones(params hclouddns.HCloudGetZonesParams) (hclouddns.HCloudAnswerGetZones, error)
	UpdateZone(zone hclouddns.HCloudZone) (hclouddns.HCloudAnswerGetZone, error)
	DeleteZone(ID string) (hclouddns.HCloudAnswerDeleteZone, error)
	CreateZone(zone hclouddns.HCloudZone) (hclouddns.HCloudAnswerGetZone, error)
	ImportZoneString(zoneID string, zonePlainText string) (hclouddns.HCloudAnswerGetZone, error)
	ExportZoneToString(zoneID string) (hclouddns.HCloudAnswerGetZonePlainText, error)
	ValidateZoneString(zonePlainText string) (hclouddns.HCloudAnswerZoneValidate, error)
	GetRecords(params hclouddns.HCloudGetRecordsParams) (hclouddns.HCloudAnswerGetRecords, error)
	UpdateRecord(record hclouddns.HCloudRecord) (hclouddns.HCloudAnswerGetRecord, error)
	DeleteRecord(ID string) (hclouddns.HCloudAnswerDeleteRecord, error)
	CreateRecord(record hclouddns.HCloudRecord) (hclouddns.HCloudAnswerGetRecord, error)
	CreateRecordBulk(record []hclouddns.HCloudRecord) (hclouddns.HCloudAnswerCreateRecords, error)
	UpdateRecordBulk(record []hclouddns.HCloudRecord) (hclouddns.HCloudAnswerUpdateRecords, error)
}

type mockHCloudClient struct {
	Token string `yaml:"token"`
}

// New instance
func New(t string) mockHCloudClientAdapter {
	return &mockHCloudClient{
		Token: t,
	}
}

// Mock all methods

func (m *mockHCloudClient) GetZone(ID string) (hclouddns.HCloudAnswerGetZone, error) {
	return hclouddns.HCloudAnswerGetZone{}, nil
}

func (m *mockHCloudClient) GetZones(params hclouddns.HCloudGetZonesParams) (hclouddns.HCloudAnswerGetZones, error) {
	return hclouddns.HCloudAnswerGetZones{
		Zones: []hclouddns.HCloudZone{
			{
				ID:           "HetznerZoneID",
				Name:         "blindage.org",
				TTL:          666,
				RecordsCount: 1,
			},
		},
	}, nil
}

// zones
func (m *mockHCloudClient) UpdateZone(zone hclouddns.HCloudZone) (hclouddns.HCloudAnswerGetZone, error) {
	return hclouddns.HCloudAnswerGetZone{}, nil
}
func (m *mockHCloudClient) DeleteZone(ID string) (hclouddns.HCloudAnswerDeleteZone, error) {
	return hclouddns.HCloudAnswerDeleteZone{}, nil
}
func (m *mockHCloudClient) CreateZone(zone hclouddns.HCloudZone) (hclouddns.HCloudAnswerGetZone, error) {
	return hclouddns.HCloudAnswerGetZone{}, nil
}
func (m *mockHCloudClient) ImportZoneString(zoneID string, zonePlainText string) (hclouddns.HCloudAnswerGetZone, error) {
	return hclouddns.HCloudAnswerGetZone{}, nil
}
func (m *mockHCloudClient) ExportZoneToString(zoneID string) (hclouddns.HCloudAnswerGetZonePlainText, error) {
	return hclouddns.HCloudAnswerGetZonePlainText{}, nil
}
func (m *mockHCloudClient) ValidateZoneString(zonePlainText string) (hclouddns.HCloudAnswerZoneValidate, error) {
	return hclouddns.HCloudAnswerZoneValidate{}, nil
}

// records
func (m *mockHCloudClient) GetRecords(params hclouddns.HCloudGetRecordsParams) (hclouddns.HCloudAnswerGetRecords, error) {
	return hclouddns.HCloudAnswerGetRecords{}, nil
}
func (m *mockHCloudClient) UpdateRecord(record hclouddns.HCloudRecord) (hclouddns.HCloudAnswerGetRecord, error) {
	return hclouddns.HCloudAnswerGetRecord{}, nil
}
func (m *mockHCloudClient) DeleteRecord(ID string) (hclouddns.HCloudAnswerDeleteRecord, error) {
	return hclouddns.HCloudAnswerDeleteRecord{}, nil
}
func (m *mockHCloudClient) CreateRecord(record hclouddns.HCloudRecord) (hclouddns.HCloudAnswerGetRecord, error) {
	return hclouddns.HCloudAnswerGetRecord{}, nil
}
func (m *mockHCloudClient) CreateRecordBulk(record []hclouddns.HCloudRecord) (hclouddns.HCloudAnswerCreateRecords, error) {
	return hclouddns.HCloudAnswerCreateRecords{}, nil
}
func (m *mockHCloudClient) UpdateRecordBulk(record []hclouddns.HCloudRecord) (hclouddns.HCloudAnswerUpdateRecords, error) {
	return hclouddns.HCloudAnswerUpdateRecords{}, nil
}

func main() {
	// Get your own token on Hetzner DNS and save in plain text file
	token, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Create new instance", string(token))
	// hdns := hclouddns.New(string(token))
	hdns := New("mytoken")

	log.Println("Get zones")
	zone, err := hdns.GetZones(hclouddns.HCloudGetZonesParams{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zone.Zones)
	log.Println(zone.Error)

	zoneID := zone.Zones[0].ID
	log.Println("Export zone to plain text:", zoneID)
	zonePlainText, err := hdns.ExportZoneToString(zoneID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zonePlainText.ZonePlainText)

	log.Println("Validate plain text zone")
	zoneValid, err := hdns.ValidateZoneString(zonePlainText.ZonePlainText)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zoneValid.ValidRecords)

	log.Println("Import zone to plain text")
	zoneImport, err := hdns.ImportZoneString(zoneID, zonePlainText.ZonePlainText)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zoneImport.Zone)

}
