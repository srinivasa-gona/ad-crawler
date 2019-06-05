package main

import (
	"ad-crawler/controller"
	"ad-crawler/model"
	"ad-crawler/repository"
	"ad-crawler/service"
	"ad-crawler/util"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config, _ := getDefaultConfig()
	database, _ := repository.GetConnection(config)
	adRepo := repository.NewAdRepositoryImpl(database)
	pdService := service.NewPublisherDataServiceImpl(config, util.NewUtilImpl())
	adService := service.NewAdServiceImpl(adRepo, pdService, util.NewUtilImpl())
	intializeApp(adService, config)

}

func intializeApp(adService service.AdService, config model.Configuration) {

	router := mux.NewRouter()
	controllerConfig := controller.NewcontrollerConfig(adService)
	controllerConfig.InitializeRouter(router)
	go adService.PopulateAllAds()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.ApplicationPort), router))

}
