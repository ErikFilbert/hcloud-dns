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
	hdns := *hclouddns.New(string(token), &hclouddns.HCloudDNS{})

	log.Println("Get zone vhSHpH5mjcB2UywP9XtZGh")
	allRecords, err := hdns.GetRecords(hclouddns.HCloudGetRecordsParams{ZoneID: "vhSHpH5mjcB2UywP9XtZGh"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(allRecords.Records)
	log.Println(allRecords.Error)

	log.Println("Get record 1e960dc913f556d884bf01c241386103 of this zone")
	record, err := hdns.GetRecord("1e960dc913f556d884bf01c241386103")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(record.Record)
	log.Println(record.Error)

	log.Println("Now update record 1e960dc913f556d884bf01c241386103")

	updateRecord := record.Record
	updateRecord.Value = "blindage.org."
	updateRecord.RecordType = hclouddns.CNAME
	updateRecord.TTL = 300

	record, err = hdns.UpdateRecord(updateRecord)
	if err != nil {
		log.Println(record.Error.Code, record.Error.Message)
		log.Fatalln(err)
	}
	log.Println("See result of update")
	log.Println(record.Record)
	log.Println(record.Error)

	log.Println("And get record 1e960dc913f556d884bf01c241386103 again to be completely ensure")
	record, err = hdns.GetRecord("1e960dc913f556d884bf01c241386103")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(record.Record)
	log.Println(record.Error)

	log.Println("I do not like this record, want to remove it")
	response, err := hdns.DeleteRecord("1e960dc913f556d884bf01c241386103")
	if err != nil {
		log.Fatalln(response.Error, err)
	}

	log.Println("Now you know how to work with library")

}
