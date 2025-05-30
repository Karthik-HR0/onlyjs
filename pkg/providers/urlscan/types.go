package urlscan

import (
	"strings"
)

var _BaseURL = "https://urlscan.io/"

type apiResponse struct {
	Status  int            `json:"status"`
	Results []searchResult `json:"results"`
	HasMore bool           `json:"has_more"`
}

type searchResult struct {
	Page archivedPage
	Sort []interface{} `json:"sort"`
}

type archivedPage struct {
	Domain   string `json:"domain"`
	MimeType string `json:"mimeType"`
	URL      string `json:"url"`
	Status   string `json:"status"`
}

func parseSort(sort []interface{}) string {
	var sortParam []string
	for _, t := range sort {
		if s, ok := t.(string); ok {
			sortParam = append(sortParam, s)
		}
	}
	return strings.Join(sortParam, ",")
}
