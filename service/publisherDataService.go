package service

import (
	"ad-crawler/model"
)

type PublisherDataService interface {
	GetPublisherData() ([]model.PublisherData, error)
}
