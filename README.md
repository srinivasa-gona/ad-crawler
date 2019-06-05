### PreRequisities

- Go installaton

### Rest Endpoints

| Protocol | End Point | Description
| ------ | ------ | ----- |
| GET | get-ads/{publisher} | Returns the Ads for specific publisher.
| POST | populate-ads | Updates Ads for all publishers specified in publisher_data.csv

### Instruction 
Extract the zip file to GOPATH 
To Build and Run
    ``` go run main.go config.go ``` or ```go install```
For Testing
    ```go test```
 
### Configuration

Configuration is customized using config.json

| Parameter | Default | Description
| ------ | ------ | ------ |
| DBFileLocation | ads_txt.db | Location of SQLLite file|
| PublishersFileLocation | publisher_data.csv | file for loading publisher data|
| ApplicationPort | 8080 | Application Port|
