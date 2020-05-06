package main

import (
	"io/ioutil"
	"log"

	hclouddns "git.blindage.org/21h/hcloud-dns"
)

func main() {
	// Get your own token on Hetzner DNS
	token, err := ioutil.ReadFile("token.txt") // just pass the file name
	if err != nil {
		log.Fatalln(err)
	}

	// Create new instance
	hdns := hclouddns.New(string(token))

	// Get zone vhSHpH5mjcB2UywP9XtZGh
	allRecords, err := hdns.GetRecords("vhSHpH5mjcB2UywP9XtZGh")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(allRecords.Records)
	log.Println(allRecords.Error)

	// Get record 1e960dc913f556d884bf01c241386103
	record, err := hdns.GetRecord("1e960dc913f556d884bf01c241386103")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(record.Record)
	log.Println(record.Error)

}
