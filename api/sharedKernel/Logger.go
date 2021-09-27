package sharedKernel

import (
	"bytes"
	"fmt"
)

type Logger struct{}

func (logger *Logger) Log(message string) {
	fmt.Println(message)
}

type HttpClient struct {
	logger *Logger
}

func (client *HttpClient) GetUrl(url string) string {
	client.logger.Log("Getting " + url)

	// make an HTTP request
	return "You have been requested: =>" + url
}

func NewHttpClient(logger *Logger) *HttpClient {
	return &HttpClient{logger}
}

type ConcatService struct {
	logger *Logger
	client *HttpClient
}

func (service *ConcatService) GetAll(urls ...string) string {
	service.logger.Log("Running dataFetchService from")

	var result bytes.Buffer

	for _, url := range urls {
		result.WriteString(service.client.GetUrl(url))
	}

	return result.String()
}

func NewConcatService(logger *Logger, client *HttpClient) *ConcatService {
	return &ConcatService{logger, client}
}
