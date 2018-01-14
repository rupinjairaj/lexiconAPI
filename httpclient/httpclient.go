package httpclient

import (
	"io/ioutil"
	"log"
	"net/http"

	"lexiconAPI/common"
)

// OxfordAPIClient is the generic http
// client that is used to communicate
// with the Oxford Dictionary API
type OxfordAPIClient struct {
}

// Get :
func (oxford OxfordAPIClient) Get(wordID string) []byte {
	url := common.BaseURL + common.SourceLang + wordID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return []byte(common.ClientErrMessage)
	}

	req.Header.Set(common.AppIDHeader, common.AppIDValue)
	req.Header.Set(common.AppKeyHeader, common.AppKeyValue)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return []byte(common.ClientErrMessage)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return []byte(common.ClientErrMessage)
	}

	return body
}
