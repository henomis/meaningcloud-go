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

	request := &request.TextClustering{}
	request.Text = "He earns $200,000/yr and still has a mortgage on his house :(\nZara clothes will be the death of my credit card\nMy bank insisted I destroyed my credit card before I could get a mortgage\nI'm not paying the mortgage or credit card bills\nTell them you've never had a loan, you have no mortgage"
	request.ID = newString("text01\ntext02\ntext03\ntext04\ntext05")
	request.InputLanguage = "en"

	// Text analysis
	response, err := meaningCloudClient.TextClustering(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Ok() {
		log.Fatalf("error: %s", response.Error())
	}

	bytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))

}

// Support methods
func newString(s string) *string {
	return &s
}
