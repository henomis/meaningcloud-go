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

	request := &request.LanguageIdentification{}
	request.Text = `Robert Downey Jr has topped Forbes magazine's annual list of the highest paid actors for the second 
	year in a row. The 49-year-old star of the Iron Man and Avengers films made an estimated $75m over the past year, 
	beating rivals Dwayne Johnson, Bradley Cooper, Chris Hemsworth and Leonardo DiCaprio.`

	// Text analysis
	response, err := meaningCloudClient.LanguageIdentification(request)
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
	response, err = meaningCloudClient.LanguageIdentification(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Status.IsSuccess() {
		log.Fatalf("error: %s", response.Status.Error())
	}

	bytes, _ = json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

}
