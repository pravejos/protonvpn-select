package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type logicals struct {
	Code           int             `json:"Code"`
	LogicalServers []logicalServer `json:"LogicalServers"`
	// Config not needed as of now.
}
type logicalServer struct {
	Name         string   `json:"Name"`
	EntryCountry string   `json:"EntryCountry"`
	ExitCountry  string   `json:"ExitCountry"`
	Domain       string   `json:"Domain"`
	Tier         int      `json:"Tier"`
	Features     int      `json:"Features"`
	Region       string   `json:"Region"`
	City         string   `json:"City"`
	ID           string   `json:"ID"`
	Location     location `json:"Location"`
	Status       int      `json:"Status"`
	Servers      []server `json:"Servers"`
	Load         int      `json:"Load"`
	Score        float32  `json:"Score"`
}

type location struct {
	Latitude  float32 `json:"Latitude"`
	Longitude float32 `json:"Longitude"`
}

type server struct {
	EntryIP string `json:"EntryIP"`
	ExitIP  string `json:"ExitIP"`
	Domain  string `json:"Domain"`
	ID      string `json:"ID"`
	Status  int    `json:"Status"`
}

func getLogicals() (resp *logicals, err error) {
	apiResp, err := http.Get("https://api.protonmail.ch/vpn/logicals")
	if err != nil {
		return nil, err
	}

	defer apiResp.Body.Close()
	body, err := ioutil.ReadAll(apiResp.Body)
	logicals := logicals{}
	json.Unmarshal(body, &logicals)

	return &logicals, nil
}
