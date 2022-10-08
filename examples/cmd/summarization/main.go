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

	request := &request.Summarization{}
	request.Text = `Robert Downey Jr has topped Forbes magazine's annual list of the highest paid actors for the second 
	year in a row. The 49-year-old star of the Iron Man and Avengers films made an estimated $75m over the past year, 
	beating rivals Dwayne Johnson, Bradley Cooper, Chris Hemsworth and Leonardo DiCaprio.`
	request.Sentences = "5"

	// Text analysis
	response, err := meaningCloudClient.Summarization(request)
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
	response, err = meaningCloudClient.Summarization(request)
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
	response, err = meaningCloudClient.Summarization(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Ok() {
		log.Fatalf("error: %s", response.Error())
	}

	bytes, _ = json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

}
