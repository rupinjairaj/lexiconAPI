package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"lexiconAPI/constants"
)

// OxfordAPIClient is the generic http
// client that is used to communicate
// with the Oxford Dictionary API
type OxfordAPIClient struct {
}

// Get :
func (oxford OxfordAPIClient) Get(wordID string) []byte {
	url := constants.BaseURL + constants.SourceLang + wordID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set(constants.AppIDHeader, constants.AppIDValue)
	req.Header.Set(constants.AppKeyHeader, constants.AppKeyValue)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}
