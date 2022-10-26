package main

import (
	"io"
	"log"
	"os"
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

	request := &request.Parser{}
	request.Text = `Robert Downey Jr has topped Forbes magazine's annual list of the highest paid actors for the second 
	year in a row. The 49-year-old star of the Iron Man and Avengers films made an estimated $75m over the past year, 
	beating rivals Dwayne Johnson, Bradley Cooper, Chris Hemsworth and Leonardo DiCaprio.`
	request.TopicTypes = newString("e")
	request.UnknownWords = newString("y")
	request.InputLanguage = "en"

	// Text analysis
	reader, err := meaningCloudClient.ParserImage(request)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}
	defer reader.Close()

	// this will save the output to a file
	file, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		log.Fatal(err)
	}

}

// Support methods
func newString(s string) *string {
	return &s
}
