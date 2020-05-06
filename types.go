package hclouddns

// Base types
type HCloudDNS struct {
	token string `yaml:"token"`
}

// RecordType supported by Hetzner
type RecordType string

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

type HCloudError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HCloudRecord struct {
	RecordType RecordType `json:"type"`
	ID         string     `json:"id"`
	Created    string     `json:"created"`
	Modified   string     `json:"modified"`
	ZoneID     string     `json:"zone_id"`
	Name       string     `json:"name"`
	Value      string     `json:"value"`
	TTL        string     `json:"ttl"`
}

type HCloudMeta struct {
	Pagination struct {
		Page         int `json:"page"`
		PerPage      int `json:"per_page"`
		LastPage     int `json:"last_page"`
		TotalEntries int `json:"total_entries"`
	} `json:"pagination"`
}
