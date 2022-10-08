package main

import (
	"encoding/json"
	"fmt"
	"log"

	meaningcloudgo "github.com/henomis/meaningcloud-go"
	"github.com/henomis/meaningcloud-go/pkg/request"
)

const Key = "YOUR_API_KEY"

func main() {

	meaningCloudClient := meaningcloudgo.New(
		meaningcloudgo.MeaningCloudEndpoint,
		Key,
	)

	request := &request.TextClassification{}
	request.Model = "IPTC_en"
	request.Text = "The 85th Academy Awards ceremony took place February 24, 2013."

	// Text analysis
	response, err := meaningCloudClient.TextClassification(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Ok() {
		log.Fatalf("error: %s", response.Error())
	}

	bytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

	// URL analysis
	request.Text = ""
	request.URL = "https://en.wikipedia.org/wiki/Star_Trek"
	response, err = meaningCloudClient.TextClassification(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Ok() {
		log.Fatalf("error: %s", response.Error())
	}

	bytes, _ = json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

	// File analysis
	request.URL = ""
	request.Document = "./examples/resources/file.txt"
	response, err = meaningCloudClient.TextClassification(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Ok() {
		log.Fatalf("error: %s", response.Error())
	}

	bytes, _ = json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

}
