package hclouddns

// Records answers

type HCloudAnswerGetRecord struct {
	Record   HCloudRecord `json:"record,omitempty"`
	Error    HCloudError  `json:"error,omitempty"`
	HTTPCode int
}

type HCloudAnswerGetRecords struct {
	Records  []HCloudRecord `json:"records,omitempty"`
	Meta     HCloudMeta     `json:"meta,omitempty"`
	Error    HCloudError    `json:"error,omitempty"`
	HTTPCode int
}

type HCloudAnswerCreateRecords struct {
	Records        []HCloudRecord `json:"records,omitempty"`
	ValidRecords   []HCloudRecord `json:"valid_records,omitempty"`
	InvalidRecords []HCloudRecord `json:"invalid_records,omitempty"`
	Error          HCloudError    `json:"error,omitempty"`
	HTTPCode       int
}

type HCloudAnswerUpdateRecords struct {
	Records        []HCloudRecord `json:"records,omitempty"`
	InvalidRecords []HCloudRecord `json:"failed_records,omitempty"`
	Error          HCloudError    `json:"error,omitempty"`
	HTTPCode       int
}

// Zones answers

type HCloudAnswerGetZone struct {
	Zone     HCloudZone  `json:"zone,omitempty"`
	Error    HCloudError `json:"error,omitempty"`
	HTTPCode int
}
type HCloudAnswerGetZones struct {
	Zones    []HCloudZone `json:"zones,omitempty"`
	Meta     HCloudMeta   `json:"meta,omitempty"`
	Error    HCloudError  `json:"error,omitempty"`
	HTTPCode int
}

// Params

type HCloudGetZonesParams struct {
	Name       string
	SearchName string
	Page       int
	PerPage    int
}
