package util

type Util interface {
	ParseCsvString(string) ([][]string, error)
	ParseCsvFile(string) ([][]string, error)
	GetHttpResponse(string) (string, error)
}

type UtilImpl struct {
}

func NewUtilImpl() Util {
	return UtilImpl{}
}
