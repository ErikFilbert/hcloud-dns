# hclouddns
--
    import "git.blindage.org/21h/hcloud-dns"


## Usage

#### type HCloudAnswerCreateRecords

```go
type HCloudAnswerCreateRecords struct {
	Records        []HCloudRecord `json:"records,omitempty"`
	ValidRecords   []HCloudRecord `json:"valid_records,omitempty"`
	InvalidRecords []HCloudRecord `json:"invalid_records,omitempty"`
	Error          HCloudError    `json:"error,omitempty"`
	HTTPCode       int
}
```


#### type HCloudAnswerGetRecord

```go
type HCloudAnswerGetRecord struct {
	Record   HCloudRecord `json:"record,omitempty"`
	Error    HCloudError  `json:"error,omitempty"`
	HTTPCode int
}
```


#### type HCloudAnswerGetRecords

```go
type HCloudAnswerGetRecords struct {
	Records  []HCloudRecord `json:"records,omitempty"`
	Meta     HCloudMeta     `json:"meta,omitempty"`
	Error    HCloudError    `json:"error,omitempty"`
	HTTPCode int
}
```


#### type HCloudAnswerUpdateRecords

```go
type HCloudAnswerUpdateRecords struct {
	Records        []HCloudRecord `json:"records,omitempty"`
	InvalidRecords []HCloudRecord `json:"failed_records,omitempty"`
	Error          HCloudError    `json:"error,omitempty"`
	HTTPCode       int
}
```


#### type HCloudDNS

```go
type HCloudDNS struct {
}
```

Base types

#### func  New

```go
func New(t string) *HCloudDNS
```
New instance

#### func (*HCloudDNS) CreateRecord

```go
func (d *HCloudDNS) CreateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error)
```
CreateRecord creates new single record Accepts HCloudRecord with record to
create, of cource no ID. Returns HCloudAnswerGetRecord with HTTPCode,
HCloudRecord and error

#### func (*HCloudDNS) CreateRecordBulk

```go
func (d *HCloudDNS) CreateRecordBulk(record []HCloudRecord) (HCloudAnswerCreateRecords, error)
```
CreateRecordBulk creates many records at once Accepts array of HCloudRecord,
converting to json and makes POST to Hetzner Returns HCloudAnswerCreateRecords
with HTTPCode, arrays of HCloudRecord with whole list, valid and invalid, error

#### func (*HCloudDNS) DeleteRecord

```go
func (d *HCloudDNS) DeleteRecord(ID string) (int, error)
```
DeleteRecord remove record by ID Accepts single ID string Returns HTTP code and
error. HTTP code 200 also raise error.

#### func (*HCloudDNS) GetRecord

```go
func (d *HCloudDNS) GetRecord(ID string) (HCloudAnswerGetRecord, error)
```
GetRecord retrieve one single record by ID Accepts single ID of record Returns
HCloudAnswerGetRecord with HCloudRecord, HTTPCode and error

#### func (*HCloudDNS) GetRecords

```go
func (d *HCloudDNS) GetRecords(zoneID string) (HCloudAnswerGetRecords, error)
```
GetRecords retrieve all records of user Accepts zone_id as string Returns
HCloudAnswerGetRecords with array of HCloudRecord, Meta, HTTPCode and error

#### func (*HCloudDNS) UpdateRecord

```go
func (d *HCloudDNS) UpdateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error)
```
UpdateRecord makes update of single record by ID Accepts HCloudRecord with
fullfilled fields Returns HCloudAnswerGetRecord with HTTP code, HCloudRecord and
error

#### func (*HCloudDNS) UpdateRecordBulk

```go
func (d *HCloudDNS) UpdateRecordBulk(record []HCloudRecord) (HCloudAnswerUpdateRecords, error)
```
UpdateRecordBulk updates many records at once Accepts array of HCloudRecord,
converting to json and makes PUT to Hetzner Returns HCloudAnswerUpdateRecords
with HTTPCode, arrays of HCloudRecord with updated and failed, error.

#### type HCloudError

```go
type HCloudError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
```


#### type HCloudMeta

```go
type HCloudMeta struct {
	Pagination struct {
		Page         int `json:"page"`
		PerPage      int `json:"per_page"`
		LastPage     int `json:"last_page"`
		TotalEntries int `json:"total_entries"`
	} `json:"pagination,omitempty"`
}
```


#### type HCloudRecord

```go
type HCloudRecord struct {
	RecordType RecordType `json:"type"`
	ID         string     `json:"id"`
	Created    string     `json:"created"`
	Modified   string     `json:"modified"`
	ZoneID     string     `json:"zone_id"`
	Name       string     `json:"name"`
	Value      string     `json:"value"`
	TTL        int        `json:"ttl"`
}
```


#### type RecordType

```go
type RecordType string
```

RecordType supported by Hetzner

```go
const (
	A     RecordType = "A"
	AAAA  RecordType = "AAAA"
	CNAME RecordType = "CNAME"
	MX    RecordType = "MX"
	NS    RecordType = "NS"
	TXT   RecordType = "TXT"
	RP    RecordType = "RP"
	SOA   RecordType = "SOA"
	HINFO RecordType = "HINFO"
	SRV   RecordType = "SRV"
	DANE  RecordType = "DANE"
	TLSA  RecordType = "TLSA"
	DS    RecordType = "DS"
	CAA   RecordType = "CAA"
)
```
