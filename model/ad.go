package model

type Record struct {
	DomainName     string `json:"domainName,omitempty"`
	PublisherActId string `json:"publisherActId,omitempty"`
	ActType        string `json:"actType,omitempty"`
	CertAuthId     string `json:"certAuthId,omitempty"`
}

type PublisherData struct {
	PublisherName string `json:"publisherName,omitempty"`
	Url           string `json:"url,omitempty"`
}
