package service

import (
	"ad-crawler/model"
	"ad-crawler/util"
	"log"
)

type PublisherDataServiceImpl struct {
	config model.Configuration
	util   util.Util
}

func NewPublisherDataServiceImpl(config model.Configuration, util util.Util) PublisherDataService {
	return PublisherDataServiceImpl{
		config: config,
		util:   util,
	}
}

func (pdsi PublisherDataServiceImpl) GetPublisherData() ([]model.PublisherData, error) {
	publisherDataList := []model.PublisherData{}
	fileName := pdsi.config.PublishersFileLocation

	lines, err := pdsi.util.ParseCsvFile(fileName)

	if err != nil {
		log.Printf("Error in reading CSV file : %s  is : %v ", fileName, err)
		return publisherDataList, err
	}
	//TODO - Add validation for URL

	// Loop through lines & turn into object
	for _, line := range lines {
		data := model.PublisherData{
			PublisherName: line[0],
			Url:           line[1],
		}

		publisherDataList = append(publisherDataList, data)
	}
	return publisherDataList, nil
}
