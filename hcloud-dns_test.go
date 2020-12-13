package hclouddns

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var (
	mockZones = []HCloudZone{
		{
			ID:           "ID_BLINDAGE",
			Name:         "blindage.org",
			TTL:          300,
			RecordsCount: 2,
		},
		{
			ID:           "ID_CBRADIO",
			Name:         "cb-radio.ru",
			TTL:          600,
			RecordsCount: 333,
		},
	}
	mockRecordsBlindage = []HCloudRecord{
		{
			RecordType: A,
			ID:         "rec1",
			ZoneID:     "ID_BLINDAGE",
			Name:       "blindage.org",
			Value:      "127.0.0.1",
			TTL:        300,
		},
		{
			RecordType: A,
			ID:         "rec2",
			ZoneID:     "ID_BLINDAGE",
			Name:       "www",
			Value:      "127.0.0.1",
			TTL:        300,
		},
	}
	mockToken = "test_token"
)

func mockAPIserver(listener net.Listener, port string) {
	fmt.Println("Starting local API server at port", port)
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/zones", responseGetZones).Methods("GET")
	router.HandleFunc("/api/v1/zones/{id}", responseGetZone).Methods("GET")
	router.HandleFunc("/api/v1/zones/{id}", responseUpdateZone).Methods("PUT")

	router.HandleFunc("/api/v1/records", responseGetRecords).Methods("GET")
	router.HandleFunc("/api/v1/records/{id}", responseGetRecord).Methods("GET")
	router.HandleFunc("/api/v1/records/{id}", responseUpdateRecord).Methods("PUT")

	err := http.Serve(listener, router)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

// New instance
func TestNew(t *testing.T) {
	// Start local server with expected data
	port := os.Getenv("PORT")
	if port == "" {
		port = "19808"
	}

	listener, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		fmt.Println("Can't listen port:", err)
	}
	go mockAPIserver(listener, port)

	// Client asks only local server
	client := &HCloudClient{
		Token:     mockToken,
		APIserver: "http://localhost:" + port,
	}

	TestingGetZone(client, t)
	TestingGetZones(client, t)
	TestingUpdateZone(client, t)

	TestingGetRecord(client, t)
	TestingGetRecords(client, t)
	TestingUpdateRecord(client, t)

}
