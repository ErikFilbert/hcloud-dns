package hclouddns

type HCloudAnswerGetRecord struct {
	Record HCloudRecord `json:"record,omitempty"`
	Error  HCloudError  `json:"error,omitempty"`
}

type HCloudAnswerGetRecords struct {
	Records []HCloudRecord `json:"records,omitempty"`
	Meta    HCloudMeta     `json:"meta,omitempty"`
	Error   HCloudError    `json:"error,omitempty"`
}

type HCloudAnswerCreateRecords struct {
	Records        []HCloudRecord `json:"records,omitempty"`
	ValidRecords   []HCloudRecord `json:"valid_records,omitempty"`
	InvalidRecords []HCloudRecord `json:"invalid_records,omitempty"`
	Error          HCloudError    `json:"error,omitempty"`
}
