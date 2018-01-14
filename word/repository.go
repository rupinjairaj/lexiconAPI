package word

import (
	"encoding/json"
	"io/ioutil"
	"lexiconAPI/common"
	"lexiconAPI/httpclient"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/context"
)

// Repository :
type Repository struct {
}

// GetWordData is the handler that will
// use  a mongodb session to check if
// the word data exists in our data store.
// If not it will query the oxford
// dictionary api to get the data,
// store the required fields in our
// mongodb store and then send the response
// to the client
func (repo *Repository) GetWordData(w http.ResponseWriter, r *http.Request) {

	// get the request body
	requestBody, requestBodyErr := ioutil.ReadAll(r.Body)
	if requestBodyErr != nil {
		log.Println(requestBodyErr)
		w.Write([]byte(common.ClientErrMessage))
	}

	// unmarshal the request body
	// into WordRequest
	var requestedWord WordRequest
	requestJSONErr := json.Unmarshal(requestBody, &requestedWord)
	if requestJSONErr != nil {
		log.Println(requestJSONErr)
		w.Write([]byte(common.ClientErrMessage))
	}

	// get a mongodb session
	dbsession := context.Get(r, "database").(*mgo.Session)

	c := dbsession.DB(common.DBNAME).C(common.COLLECTION)
	// query the collections 'c' to check
	// if the word exists
	var findResult WordDataModel
	var wordDataModel WordDataModel
	mongoFindErr := c.Find(bson.M{"dictionaryEntryResponse.results.id": requestedWord.Word}).One(&findResult)
	if mongoFindErr != nil {
		// get the word data from oxford dictionary
		wordDataModel.ID = bson.NewObjectId()
		var oxfordClient httpclient.OxfordAPIClient
		res := oxfordClient.Get(requestedWord.Word)
		dictionaryModelJSONErr := json.Unmarshal(res, &wordDataModel.DictionaryEntryResponse)
		if dictionaryModelJSONErr != nil {
			log.Println(dictionaryModelJSONErr)
			w.Write([]byte(common.ClientErrMessage))
			return
		}
		// save the response to mongodb
		dbsession.DB(common.DBNAME).C(common.COLLECTION).Insert(wordDataModel)
		// build your custom response to the client and send that
		builtResponse := buildClientResponse(wordDataModel)
		clientResponse, clientResponseErr := json.Marshal(builtResponse)
		if clientResponseErr != nil {
			log.Println(clientResponseErr)
			w.Write([]byte(common.ClientErrMessage))
			return
		}
		w.Write(clientResponse)
	} else {
		// send the local mongodb data
		builtResponse := buildClientResponse(findResult)
		clientResponse, clientResponseErr := json.Marshal(builtResponse)
		if clientResponseErr != nil {
			log.Println(clientResponseErr)
			w.Write([]byte(common.ClientErrMessage))
			return
		}
		w.Write(clientResponse)
	}
}

func buildClientResponse(model WordDataModel) []ClientResponse {
	var clientResponse []ClientResponse
	for _, result := range model.DictionaryEntryResponse.Results {
		var clientRes ClientResponse
		clientRes.Word = result.ID

		lexicalEntries := result.LexicalEntries
		for _, lexicalEntry := range lexicalEntries {

			entries := lexicalEntry.Entries
			for _, entry := range entries {

				senses := entry.Senses
				for _, sense := range senses {

					definitions := sense.Definitions
					examples := sense.Examples
					for _, definition := range definitions {
						clientRes.Definitions = append(clientRes.Definitions, definition)
					}

					for _, example := range examples {
						clientRes.Examples = append(clientRes.Examples, example.Text)
					}
				}
			}
		}
		clientResponse = append(clientResponse, clientRes)
	}
	return clientResponse
}
