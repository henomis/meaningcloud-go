package meaningcloudgo

import (
	"fmt"
	"time"

	"github.com/henomis/meaningcloud-go/internal/pkg/httpclient"
	"github.com/henomis/meaningcloud-go/pkg/request"
	"github.com/henomis/meaningcloud-go/pkg/response"
)

const (
	MeaningCloudEndpoint       = "https://api.meaningcloud.com"
	deepCategorizationPath     = "/deepcategorization-1.0"
	topicsExtractionPath       = "/topics-2.0"
	textClassificationPath     = "/class-2.0"
	sentimentPath              = "/sentiment-2.1"
	languageIdentificationPath = "/lang-4.0/identification"
	parserPath                 = "/parser-2.0"
	corporateReputationPath    = "/reputation-2.0"
	textClusteringPath         = "/clustering-1.1"
	summarizationPath          = "/summarization-1.0"
	documentStructurePath      = "/documentstructure-1.0"
)

type MeaningCloudClient struct {
	httpClient *httpclient.HttpClient
	key        string
}

func New(endpoint string, key string, timeout time.Duration) *MeaningCloudClient {
	return &MeaningCloudClient{
		httpClient: httpclient.New(endpoint, timeout),
		key:        key,
	}
}

func (m *MeaningCloudClient) SentimentAnalysis(
	sentimentRequest *request.Sentiment,
) (*response.Sentiment, error) {

	sentimentRequest.Key = m.key
	body, err := m.httpClient.Request(sentimentPath, sentimentRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	sentimentResponse := &response.Sentiment{}
	err = sentimentResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return sentimentResponse, nil
}

func (m *MeaningCloudClient) DeepCategorization(
	deepCategorizationRequest *request.DeepCategorization,
) (*response.DeepCategorization, error) {

	deepCategorizationRequest.Key = m.key
	body, err := m.httpClient.Request(deepCategorizationPath, deepCategorizationRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	deepCategorizationResponse := &response.DeepCategorization{}
	err = deepCategorizationResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return deepCategorizationResponse, nil
}

func (m *MeaningCloudClient) TopicsExtraction(
	topicsExtractionRequest *request.TopicsExtraction,
) (*response.TopicsExtraction, error) {

	topicsExtractionRequest.Key = m.key
	body, err := m.httpClient.Request(topicsExtractionPath, topicsExtractionRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	topicsExtractionResponse := &response.TopicsExtraction{}
	err = topicsExtractionResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return topicsExtractionResponse, nil
}

func (m *MeaningCloudClient) TextClassification(
	textClassificationRequest *request.TextClassification,
) (*response.TextClassification, error) {

	textClassificationRequest.Key = m.key
	body, err := m.httpClient.Request(textClassificationPath, textClassificationRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	textClassificationResponse := &response.TextClassification{}
	err = textClassificationResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return textClassificationResponse, nil
}

func (m *MeaningCloudClient) LanguageIdentification(
	languageIdentificationRequest *request.LanguageIdentification,
) (*response.LanguageIdentification, error) {

	languageIdentificationRequest.Key = m.key
	body, err := m.httpClient.Request(languageIdentificationPath, languageIdentificationRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	languageIdentificationResponse := &response.LanguageIdentification{}
	err = languageIdentificationResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return languageIdentificationResponse, nil
}

func (m *MeaningCloudClient) Parser(
	parserRequest *request.Parser,
) (*response.Parser, error) {

	parserRequest.Key = m.key
	body, err := m.httpClient.Request(parserPath, parserRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	parserResponse := &response.Parser{}
	err = parserResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return parserResponse, nil
}

func (m *MeaningCloudClient) CorporateReputation(
	corporateReputationRequest *request.CorporateReputation,
) (*response.CorporateReputation, error) {

	corporateReputationRequest.Key = m.key
	body, err := m.httpClient.Request(corporateReputationPath, corporateReputationRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	corporateReputationResponse := &response.CorporateReputation{}
	err = corporateReputationResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return corporateReputationResponse, nil
}

func (m *MeaningCloudClient) TextClustering(
	textClusteringRequest *request.TextClustering,
) (*response.TextClustering, error) {

	textClusteringRequest.Key = m.key
	body, err := m.httpClient.Request(textClusteringPath, textClusteringRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	textClusteringResponse := &response.TextClustering{}
	err = textClusteringResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return textClusteringResponse, nil
}

func (m *MeaningCloudClient) Summarization(
	summarizationRequest *request.Summarization,
) (*response.Summarization, error) {

	summarizationRequest.Key = m.key
	body, err := m.httpClient.Request(summarizationPath, summarizationRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	summarizationResponse := &response.Summarization{}
	err = summarizationResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return summarizationResponse, nil
}

func (m *MeaningCloudClient) DocumentStructure(
	documentStructureRequest *request.DocumentStructure,
) (*response.DocumentStructure, error) {

	documentStructureRequest.Key = m.key
	body, err := m.httpClient.Request(documentStructurePath, documentStructureRequest)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}

	documentStructureResponse := &response.DocumentStructure{}
	err = documentStructureResponse.Decode(body)
	if err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return documentStructureResponse, nil
}
