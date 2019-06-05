package controller

import (
	"ad-crawler/model"
	"ad-crawler/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ControllerConfig struct {
	adService service.AdService
}

func NewcontrollerConfig(as service.AdService) *ControllerConfig {
	return &ControllerConfig{
		adService: as,
	}
}

func (cc *ControllerConfig) InitializeRouter(router *mux.Router) {

	router.HandleFunc("/get-ads/{publisher}", cc.getAdsByPublisher).Methods("GET")
	router.HandleFunc("/populate-ads", cc.loadAllAds).Methods("POST")

}

func (cc *ControllerConfig) getAdsByPublisher(w http.ResponseWriter, r *http.Request) {
	var records []model.Record
	params := mux.Vars(r)
	records, err := cc.adService.GetAdsByPublisher(params["publisher"])
	if err != nil {
		log.Printf("Error in get all response  %v", err)
	}
	json.NewEncoder(w).Encode(records)
}

func (cc *ControllerConfig) loadAllAds(w http.ResponseWriter, r *http.Request) {

	err := cc.adService.PopulateAllAds()
	if err != nil {
		log.Printf("Error in get all response  %v", err)
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(nil)
}
