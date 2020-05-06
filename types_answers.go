package hclouddns

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
