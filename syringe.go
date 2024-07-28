package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

type ingest struct {
	Ingests []struct {
		ID                int     `json:"_id"`
		Availability      float64 `json:"availability"`
		Default           bool    `json:"default"`
		Name              string  `json:"name"`
		URLTemplate       string  `json:"url_template"`
		URLTemplateSecure string  `json:"url_template_secure"`
		Priority          int     `json:"priority"`
	} `json:"ingests"`
}

func main() {
	fmt.Println("Selecting your closest available ingest TTV server..")
	resp, err := http.Get("https://ingest.twitch.tv/ingests")
	if err != nil {
	panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) 
	if err != nil {
		panic(err)
	}
	var result ingest
	if err := json.Unmarshal(body, &result); err != nil {   // Parse []byte to go struct pointer
    fmt.Println("Can not unmarshal JSON")
}

for _, rec := range result.Ingests {
	if(rec.Availability == 1) {
	fmt.Println("ID: ", rec.ID, "\nAvailability: ", rec.Availability, "\nName: ", rec.Name, "\nURL Template: ", rec.URLTemplate, "\nSecure URL Template: ", rec.URLTemplateSecure) 
	break; }
}
}
