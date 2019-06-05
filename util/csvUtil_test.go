package util_test

import (
	"ad-crawler/util"
	"testing"
)

func TestParseCsvString(t *testing.T) {

	utilImpl := util.NewUtilImpl()
	lines, err := utilImpl.ParseCsvString(`
 # CNN.com/ads.txt
 # 
 # DOMESTIC
google.com, pub-7439281311086140, DIRECT, f08c47fec0942fa0 # banner, video, native`)

	if len(lines) != 1 {
		t.Errorf("Expecting 1 lines from csv but got %v", len(lines))
	}
	if err != nil {
		t.Errorf("Expecting no error in parsing. Actual error is : %v", err)
	}
	if len(lines[0]) != 4 {
		t.Errorf("Expecting 4 colums in csv but got : %d", len(lines[0]))
	}
}
