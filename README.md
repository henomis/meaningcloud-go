# MeaningCloud SDK for Go


[![Build Status](https://github.com/henomis/meaningcloud-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/henomis/meaningcloud-go/actions/workflows/test.yml?query=branch%3Amain) [![GoDoc](https://godoc.org/github.com/henomis/meaningcloud-go?status.svg)](https://godoc.org/github.com/henomis/meaningcloud-go) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/meaningcloud-go)](https://goreportcard.com/report/github.com/henomis/meaningcloud-go) [![GitHub release](https://img.shields.io/github/release/henomis/meaningcloud-go.svg)](https://github.com/henomis/meaningcloud-go/releases)

This is MeaningCloud's **unofficial** Go client, designed to enable you to use MeaningCloud's services easily from your own applications.

## MeaningCloud

MeaningCloud is a cloud-based text analytics service that through APIs allows you extract meaning from all kind of unstructured content: social conversation, articles, documents.

## SDK versions


|                          | v1.1                 |
|--------------------------|------------------------|
| Deep Categorization      | deepcategorization-1.0 |
| Topics Extraction        | topics-2.0             |
| Text Classification      | class-2.0              |
| Sentiment Analysis       | sentiment-2.1          |
| Language Identification  | lang-4.0               |
| Parser                   | parser-2.0             |
| Corporate Reputation     | reputation-2.0         |
| Text Clustering          | clustering-1.1		    |
| Summarization            | summarization-1.0      |
| Document Structure       | documentstructure-1.0  |



## Getting started

### Installation

You can load meaningcloud-go into your project by using:
```
go get github.com/henomis/meaningcloud-go
```


### Configuration

The only thing you need to start using MeaningCloud's APIs is the developer license key. Copy it and paste it in the corresponding place in the code, select the API you want to use and the parameters you want to use, and that's it.


### Usage

Please refer to the [examples folder](examples/) to see how to use the SDK.

Here below a simple usage example:

```go
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
		10 * time.Second,
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

}
```
