package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	meaningcloudgo "github.com/henomis/meaningcloud-go"
	"github.com/henomis/meaningcloud-go/pkg/request"
)

const Key = "YOUR_API_KEY"

func main() {

	meaningCloudClient := meaningcloudgo.New(
		meaningcloudgo.MeaningCloudEndpoint,
		Key,
		10*time.Second,
	)

	// Text analysis
	request := &request.Sentiment{}
	request.Key = Key
	request.InputLanguage = "en"
	request.Text = "Main dishes were quite good, but desserts were too sweet for me."
	response, err := meaningCloudClient.SentimentAnalysis(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Status.IsSuccess() {
		log.Fatalf("error: %s", response.Status.Error())
	}

	bytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

	// URL analysis
	request.Text = ""
	request.URL = "https://en.wikipedia.org/wiki/Star_Trek"
	response, err = meaningCloudClient.SentimentAnalysis(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Status.IsSuccess() {
		log.Fatalf("error: %s", response.Status.Error())
	}

	bytes, _ = json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

	// File analysis
	request.URL = ""
	request.Document = "./examples/resources/file.txt"
	response, err = meaningCloudClient.SentimentAnalysis(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Status.IsSuccess() {
		log.Fatalf("error: %s", response.Status.Error())
	}

	bytes, _ = json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

}
