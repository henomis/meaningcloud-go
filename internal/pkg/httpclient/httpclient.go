package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/henomis/meaningcloud-go/internal/pkg/multipartform"
)

type HttpClient struct {
	httpClient *http.Client
	baseURL    string
}

type RequestData interface {
	ToMultipartForm() (multipartform.MultipartForm, error)
}

func New(baseURL string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

func (h *HttpClient) Request(path string, requestData RequestData) (io.ReadCloser, error) {

	multipartForm, err := requestData.ToMultipartForm()
	if err != nil {
		return nil, fmt.Errorf("invalid query: %w", err)
	}

	body := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(body)

	for multipartField, value := range multipartForm {
		switch multipartField.Type {
		case multipartform.MultipartFormField:
			err := multipartWriter.WriteField(multipartField.Name, value)
			if err != nil {
				return nil, err
			}
		case multipartform.MultipartFormFile:
			part, err := multipartWriter.CreateFormFile(multipartField.Name, "file")
			if err != nil {
				return nil, err
			}
			file, err := os.Open(value)
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				file.Close()
				return nil, err
			}
			file.Close()
		}
	}

	multipartWriter.Close()

	request, err := http.NewRequest("POST", h.baseURL+path, bytes.NewReader(body.Bytes()))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", response.StatusCode)
	}

	return response.Body, nil
}
