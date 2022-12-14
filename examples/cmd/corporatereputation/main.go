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

	request := &request.CorporateReputation{}
	request.Model = newString("CorporateReputation")
	request.Text = `Endesa is building the largest photovoltaic project for self-consumption in the Balearic Islands`
	request.InputLanguage = "en"

	// Text analysis
	response, err := meaningCloudClient.CorporateReputation(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Status.IsSuccess() {
		log.Fatalf("error: %s", response.Status.Error())
	}

	bytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

}

// Support methods
func newString(s string) *string {
	return &s
}
