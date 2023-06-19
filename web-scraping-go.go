package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type WebsiteDetail struct {
	URL        string `json:"url"`
	Results    string `json:"results"`
	StatusCode int    `json:"status_code"`
	TotalChars int    `json:"total_chars"`
	Length     int    `json:"length"`
}

type ScrappingMessage struct {
	Message string `json:"message"`
}

func stripHTMLTags(htmlContent string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return ""
	}

	strippedText := doc.Text()
	return strippedText
}

func sanitizeHTML(htmlContent string) string {
	p := bluemonday.UGCPolicy()
	sanitizedHTML := p.Sanitize(htmlContent)
	return sanitizedHTML
}

func getWebsiteDetail(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	response, err := http.Get(url)
	if err != nil {
		errorResponse := ErrorResponse{Error: fmt.Sprintf("Error requesting URL: %s", err.Error())}
		jsonData, _ := json.Marshal(errorResponse)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonData)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorResponse := ErrorResponse{Error: fmt.Sprintf("Error retrieving website: %s", response.Status)}
		jsonData, _ := json.Marshal(errorResponse)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonData)
		return
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorResponse := ErrorResponse{Error: "Error reading response body"}
		jsonData, _ := json.Marshal(errorResponse)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonData)
		return
	}
	htmlContent := string(bodyBytes)

	strippedContent := stripHTMLTags(htmlContent)
	strippedContent = strings.Join(strings.Fields(strippedContent), " ")

	sanitizedContent := sanitizeHTML(strippedContent)

	if len(sanitizedContent) > 10000 {
		sanitizedContent = sanitizedContent[:10000]
	}

	websiteDetail := WebsiteDetail{
		URL:        url,
		Results:    sanitizedContent,
		StatusCode: response.StatusCode,
		TotalChars: len(sanitizedContent),
		Length:     len(strings.Fields(sanitizedContent)),
	}

	jsonData, err := json.Marshal(websiteDetail)
	if err != nil {
		errorResponse := ErrorResponse{Error: "Error marshaling JSON"}
		jsonData, _ := json.Marshal(errorResponse)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonData)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	message := ScrappingMessage{
		Message: "Web scraping with Go",
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		errorResponse := ErrorResponse{Error: "Error marshaling JSON"}
		jsonData, _ := json.Marshal(errorResponse)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonData)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	message := "Server is running"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/website-detail", getWebsiteDetail)
	http.HandleFunc("/", getIndex)
	http.HandleFunc("/ping", pingHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
