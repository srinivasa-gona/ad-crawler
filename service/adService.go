package service

import (
	"ad-crawler/model"
)

type AdService interface {
	InsertAdsInDatabase([]model.PublisherData) error
	PopulateAllAds() error
	GetAdsByPublisher(publisher string) ([]model.Record, error)
}
