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

	log.Println("Get zone vhSHpH5mjcB2UywP9XtZGh")
	zone, err := hdns.GetZones(hclouddns.HCloudGetZonesParams{Name: "blindage.org"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(zone.Zones)
	log.Println(zone.Error)

}
