package main

import (
	"io/ioutil"
	"log"

	hclouddns "git.blindage.org/21h/hcloud-dns"
)

func main() {
	// Get your own token on Hetzner DNS and save in plain text file
	token, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Create new instance")
	hdns := hclouddns.New(string(token))

	log.Println("Get zones")
	zone, err := hdns.GetZones(hclouddns.HCloudGetZonesParams{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zone.Zones)
	log.Println(zone.Error)

	log.Println("Export zone to plain text")
	zonePlainText, err := hdns.ExportZoneToString("vhSHpH5mjcB2UywP9XtZGh")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zonePlainText.ZonePlainText)

	log.Println("Import zone to plain text")
	zoneImport, err := hdns.ImportZoneString("vhSHpH5mjcB2UywP9XtZGh", zonePlainText.ZonePlainText)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zoneImport.Zone)

}
