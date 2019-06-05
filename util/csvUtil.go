package util

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func (utilImpl UtilImpl) ParseCsvFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("error in reading csv file %v", err)
		return nil, err
	}
	defer file.Close()

	// Read File into a Variable
	csvReader := csv.NewReader(file)
	csvReader.Comment = '#'
	csvReader.TrimLeadingSpace = true
	lines, err := csvReader.ReadAll()
	if err != nil {
		log.Printf("error in reading csv file %v", err)
		return nil, err
	}
	return lines, nil
}

func (utilImpl UtilImpl) ParseCsvString(data string) ([][]string, error) {

	lines := strings.Split(data, "\n")
	var res [][]string
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		if strings.Contains(line, "#") {
			line = line[:strings.Index(line, "#")]
		}
		res = append(res, strings.Split(line, ","))
	}

	return res, nil
}
