# Hetzner DNS golang library

This library made to interact with Hetzner DNS API in most easy way. Hopefully in future it will be used for Hetzner external-dns provider. Check out [example](example) directory and [API_help.md](API_help.md).

Get your own token on Hetzner DNS and place it to `token` variable and run code

```go
token := "jcB2UywP9XtZGhvhSHpH5m"
zone := "vhSHpH5mjcB2UywP9XtZGh"

log.Println("Create new instance")
hdns := hclouddns.New(token)

log.Println("Get zone", zone)

allRecords, err := hdns.GetRecords(zone)
if err != nil {
	log.Fatalln(err)
}

log.Println(allRecords.Records)
log.Println(allRecords.Error)
```

At this moment library under development, be patient.