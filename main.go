package syringe

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
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

type stringingest struct {
	ID	            string
	Availability	string
	Default		    string
	Name			string
	URLTemplate		string
	URLTemplateSecure string
}

func Resolve()(values stringingest) {
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

	err = json.Unmarshal(body, &Result); if err != nil {  
    panic(err)
}
		parts := strings.Split(string(body), "},")
		reming := strings.Split(parts[1], "{")
		sliced := strings.Split(reming[1], ",")
		
		var returnvalue = new(stringingest)
		returnvalue.ID = sliced[0]
		returnvalue.Availability = sliced[1]
		returnvalue.Default = sliced[2]
		returnvalue.Name = sliced[3] + sliced[4]
		returnvalue.URLTemplate = sliced[5] + strings.Split(reming[2], ",")[0]
		returnvalue.URLTemplateSecure = strings.Split(reming[2], ",")[1] + `{stream_key}"`
		return *returnvalue
		}

