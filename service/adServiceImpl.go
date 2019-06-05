package service

import (
	"ad-crawler/model"
	"ad-crawler/repository"
	"ad-crawler/util"
	"log"
)

type AdServiceImpl struct {
	adRepo               repository.AdRepository
	publisherDataService PublisherDataService
	util                 util.Util
}

func NewAdServiceImpl(repo repository.AdRepository, pdsi PublisherDataService, util util.Util) AdService {
	return AdServiceImpl{
		adRepo:               repo,
		publisherDataService: pdsi,
		util:                 util,
	}
}

func (asi AdServiceImpl) PopulateAllAds() error {
	asi.adRepo.CreateTable()
	list, _ := asi.publisherDataService.GetPublisherData()
	err := asi.InsertAdsInDatabase(list)
	if err != nil {
		log.Printf("Error in loading ads %v", err)
		return nil
	}
	return nil
}

func (asi AdServiceImpl) GetAdsByPublisher(publisher string) ([]model.Record, error) {
	return asi.adRepo.GetRecords(publisher)
}

func (asi AdServiceImpl) InsertAdsInDatabase(crawlUrlList []model.PublisherData) error {

	for _, crawlUrl := range crawlUrlList {
		recordsList := []model.Record{}

		httpResponse, err := asi.util.GetHttpResponse(crawlUrl.Url)

		if err != nil {
			log.Printf("Error in getting http response for url : %v is %v", crawlUrl.Url, err)
			continue
		}

		lines, err := asi.util.ParseCsvString(httpResponse)
		if err != nil {
			log.Printf("Error in reading CSV file %v ", err)
		}

		for _, line := range lines {

			if len(line) < 3 {
				log.Printf("malformed line %v", line)
				continue
			}

			data := model.Record{
				DomainName:     line[0],
				PublisherActId: line[1],
				ActType:        line[2],
			}

			if len(line) > 3 {
				data.CertAuthId = line[3]
			}
			recordsList = append(recordsList, data)
		}

		//TODO - It can be called asynchronously. Not working with sqllite
		asi.adRepo.InsertRecords(crawlUrl.PublisherName, recordsList)
	}
	log.Printf("Ads data load completed")
	return nil
}
