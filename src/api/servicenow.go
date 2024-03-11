package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type AlertDataFrame struct {
	Version           string `json:"version"`
	GroupKey          string `json:"groupKey"`
	TruncatedAlerts   string `json:"truncatedAlerts"`
	Status            string `json:"status"`
	Receiver          string `json:"receiver"`
	GroupLabels       string `json:"groupLabels"`
	CommonLabels      string `json:"commonLabels"`
	CommonAnnotations string `json:"commonAnnotations"`
	ExternalURL       string `json:"externalURL"`
	Alerts            []struct {
		Status       string `json:"status"`
		Labels       string `json:"labels"`
		Annotations  string `json:"annotations"`
		StartsAt     string `json:"startsAt"`
		EndsAt       string `json:"endsAt"`
		GeneratorURL string `json:"generatorURL"`
		Fingerprint  string `json:"fingerprint"`
	} `json:"alerts"`
}

type ServiceNowDataFrame struct {
}

func SnowHandler(w http.ResponseWriter, r *http.Request) {
	var alertData AlertDataFrame
	var snowData ServiceNowDataFrame

	switch r.Method {
	case "GET":
		io.WriteString(w, "GET\n")

	case "POST":
		err := json.NewDecoder(r.Body).Decode(&alertData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(snowData).Encode(&alertData)
		print(w)
		//test
	case "PUT":
		io.WriteString(w, "PUT\n")
	case "DELETE":
		io.WriteString(w, "DELETE\n")
	default:
		io.WriteString(w, "FAILURE\n")
		return
	}
	return
}
