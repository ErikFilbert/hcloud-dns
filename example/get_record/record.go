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

	log.Println("Get zone ZP45XES2phNmS2KG8PxuaM")
	allRecords, err := hdns.GetRecords(hclouddns.HCloudGetRecordsParams{ZoneID: "ZP45XES2phNmS2KG8PxuaM"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(allRecords.Records)
	log.Println(allRecords.Error)

	log.Println("Get first record of this zone")
	record, err := hdns.GetRecord(allRecords.Records[0].ID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(record.Record)
	log.Println(record.Error)

	log.Println("Now you know how to work with library")

}
