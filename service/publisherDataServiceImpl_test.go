package service_test

import (
	"ad-crawler/model"
	"ad-crawler/service"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockUtil struct {
	parseCsvStringMock  func() ([][]string, error)
	getHttpResponseMock func() (string, error)
	parseCsvFileMock    func() ([][]string, error)
}

func (m mockUtil) ParseCsvString(data string) ([][]string, error) {
	return m.parseCsvStringMock()
}

func (m mockUtil) ParseCsvFile(data string) ([][]string, error) {
	return m.parseCsvFileMock()
}

func (m mockUtil) GetHttpResponse(data string) (string, error) {
	return m.getHttpResponseMock()
}

func TestGetPublisherData_HappyPath(t *testing.T) {

	config := model.Configuration{PublishersFileLocation: "publisher_data.csv"}
	var csvResp [][]string
	csvResp = append(csvResp, strings.Split("cnn,http://cnn.com/ads.txt", ","))
	csvResp = append(csvResp, strings.Split("wordpress,http://wordpress.com/ads.txt", ","))

	utilMock := mockUtil{parseCsvFileMock: func() ([][]string, error) { return csvResp, nil }}
	mockPublisherDataServiceImpl := service.NewPublisherDataServiceImpl(config, utilMock)
	res, _ := mockPublisherDataServiceImpl.GetPublisherData()

	assert.Equal(t, len(res), 2, "2 publishers should be returned")
	assert.Equal(t, res[0].PublisherName, "cnn", "publisher name incorrect")
	assert.Equal(t, res[0].Url, "http://cnn.com/ads.txt", "publisher url incorrect")
	assert.Equal(t, res[1].PublisherName, "wordpress", "publisher name incorrect")
	assert.Equal(t, res[1].Url, "http://wordpress.com/ads.txt", "publisher url incorrect")

}

func TestGetPublisherData_ErrorReadingFile(t *testing.T) {

	config := model.Configuration{PublishersFileLocation: "publisher_data.csv"}

	utilMock := mockUtil{parseCsvFileMock: func() ([][]string, error) { return nil, errors.New("some error") }}
	mockPublisherDataServiceImpl := service.NewPublisherDataServiceImpl(config, utilMock)
	res, _ := mockPublisherDataServiceImpl.GetPublisherData()

	assert.Equal(t, len(res), 0, "In case of error, empty array should be returned")

}
