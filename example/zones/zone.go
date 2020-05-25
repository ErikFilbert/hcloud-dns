package main

import (
	"io/ioutil"
	"log"

	hclouddns "git.blindage.org/21h/hcloud-dns"
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
	Token  string `yaml:"token"`
	Client mockHCloudClientAdapter
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

func main() {
	// Get your own token on Hetzner DNS and save in plain text file
	token, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Create new instance", string(token))
	hdns := hclouddns.New(string(token), &hclouddns.HCloudClient{})
	// hdns := hclouddns.New("", &mockHCloudClient{})

	log.Println("Get zones")
	zone, err := hdns.Client.GetZones(hclouddns.HCloudGetZonesParams{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zone.Zones)
	log.Println(zone.Error)

	zoneID := zone.Zones[0].ID
	log.Println("Export zone to plain text")
	zonePlainText, err := hdns.Client.ExportZoneToString(zoneID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zonePlainText.ZonePlainText)

	log.Println("Validate plain text zone")
	zoneValid, err := hdns.Client.ValidateZoneString(zonePlainText.ZonePlainText)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zoneValid.ValidRecords)

	log.Println("Import zone to plain text")
	zoneImport, err := hdns.Client.ImportZoneString(zoneID, zonePlainText.ZonePlainText)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zoneImport.Zone)

}
