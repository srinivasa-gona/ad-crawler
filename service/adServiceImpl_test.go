package service_test

import (
	"ad-crawler/mocks"
	"ad-crawler/model"
	"ad-crawler/service"
	"ad-crawler/util"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestInsertAdsInDatabase_Happypath(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	MockAdRepo := mocks.NewMockAdRepository(mockCtrl)
	MockAdRepo.EXPECT().InsertRecords(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	MockAdRepo.EXPECT().CreateTable().Return(nil).Times(1)

	var publisherDataList []model.PublisherData
	publisherData := model.PublisherData{PublisherName: "cnn",
		Url: "http://cnn.com/ads.txt"}
	publisherDataList = append(publisherDataList, publisherData)

	MockPDService := mocks.NewMockPublisherDataService(mockCtrl)
	MockPDService.EXPECT().GetPublisherData().AnyTimes().Return(publisherDataList, nil)

	MockAdService := service.NewAdServiceImpl(MockAdRepo, MockPDService, util.NewUtilImpl())

	MockAdService.PopulateAllAds()

}

func TestInsertAdsInDatabase_SkipFailed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	MockAdRepo := mocks.NewMockAdRepository(mockCtrl)
	MockAdRepo.EXPECT().InsertRecords("cnn", gomock.Any()).Return(nil).Times(1)
	MockAdRepo.EXPECT().CreateTable().Return(nil).Times(1)

	publisherDataList := []model.PublisherData{{PublisherName: "dummy",
		Url: "http://dummyurl/ads.txt"}, {PublisherName: "cnn",
		Url: "http://cnn.com/ads.txt"}}

	MockPDService := mocks.NewMockPublisherDataService(mockCtrl)
	MockPDService.EXPECT().GetPublisherData().AnyTimes().Return(publisherDataList, nil)

	MockAdService := service.NewAdServiceImpl(MockAdRepo, MockPDService, util.NewUtilImpl())

	MockAdService.PopulateAllAds()

}
