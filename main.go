package syringe

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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
	ID	            int
	Availability	float64
	Default		    bool
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
		reming := strings.Split(parts[0], "{")
		sliced := strings.Split(reming[2], ",")

		var returnvalue = new(stringingest)
		returnvalue.ID, _ = strconv.Atoi(strings.Split(sliced[0], ": ")[1])
		returnvalue.Availability, _ = strconv.ParseFloat(strings.Split(sliced[1], ": ")[1], 64)
		returnvalue.Default, _ = strconv.ParseBool(strings.Split(sliced[2], ": ")[1])
		returnvalue.Name = strings.Trim(strings.Split(sliced[3], ": ")[1] + sliced[4], `"`)
		returnvalue.URLTemplate = strings.Trim(strings.Split(sliced[5], ": ")[1], `"`)
		returnvalue.URLTemplateSecure = "rtmps" + strings.Trim(returnvalue.URLTemplate, "rtmp")
		return *returnvalue
		}