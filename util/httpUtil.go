package util

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func (utilImpl UtilImpl) GetHttpResponse(url string) (string, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Error in getting http data %v", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error in getting http data %v", err)
			return "", err
		}
		return string(bodyBytes), nil
	} else {
		return "", errors.New("Error in getting data from URL " + url + " : " + resp.Status)
	}

}
