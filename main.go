package syringe

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strconv"
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

func Resolve()(values []string) {
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
	var Result ingest

	if err := json.Unmarshal(body, &Result); err != nil {  
    fmt.Println("Can not unmarshal JSON")
}

	var array []string

	args := os.Args[1:]
	if(len(args) != 0) {
	if(slices.Contains(args, "-r")) {
	for _, rec := range Result.Ingests {
		if(rec.Availability == 1) {
			for i := 0; i > 4; i++ {
				switch i {
				case 0: array = append(array, strconv.Itoa(rec.ID))
				case 1: array = append(array, strconv.Itoa(int(rec.Availability)))
				case 2: array = append(array, string(rec.Name))
				case 3: array = append(array, string(rec.URLTemplate))
				case 4: array = append(array, string(rec.URLTemplateSecure))
			}
		}
		}}}}

	if(len(args) == 0) {
		for _, rec := range Result.Ingests {
			if(rec.Availability == 1) {
				fmt.Println("ID: ", rec.ID, "\nAvailability: ", rec.Availability, "\nName: ", rec.Name, "\nURL Template: ", rec.URLTemplate, "\nSecure URL Template: ", rec.URLTemplateSecure)  
		break; }
	}
	
}
return array
}