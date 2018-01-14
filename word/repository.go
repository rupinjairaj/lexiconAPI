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
func (repo *Repository) GetWordData(w http.ResponseWriter, r *http.Request) []byte {

	// get the request body
	requestBody, requestBodyErr := ioutil.ReadAll(r.Body)
	if requestBodyErr != nil {
		log.Println(requestBodyErr)
		return
	}

	// unmarshal the request body
	// into WordRequest
	var requestedWord WordRequest
	requestJSONErr := json.Unmarshal(requestBody, &requestedWord)
	if requestJSONErr != nil {
		log.Println(requestJSONErr)
	}

	// get a mongodb session
	dbsession := context.Get(r, "database").(*mgo.Session)

	c := dbsession.DB(common.DBNAME).C(common.COLLECTION)
	// query the collections 'c' to check
	// if the word exists
	var findResult WordDataModel
	var wordDataModel WordDataModel
	mongoFindErr := c.Find(bson.M{"dictionaryentryresponse.results.id": requestedWord.Word}).One(&findResult)
	if mongoFindErr != nil {
		// get the word data from oxford dictionary
		wordDataModel.ID = bson.NewObjectId()
		var oxfordClient httpclient.OxfordAPIClient
		res := oxfordClient.Get(requestedWord.Word)
		dictionaryModelJSONErr := json.Unmarshal(res, &wordDataModel.DictionaryEntryResponse)
		if dictionaryModelJSONErr != nil {
			log.Println(dictionaryModelJSONErr)
		} else {
			// save the response to mongodb
			dbsession.DB(common.DBNAME).C(common.COLLECTION).Insert(wordDataModel)
			// build your custom response to the client and send that
			clientResponse, clientResponseErr := json.Marshal(buildClientResponse(wordDataModel))
			if clientResponseErr != nil {
				log.Println(clientResponseErr)
				return []byte(common.ClientErrMessage)
			} else {
				return clientResponse
			}
		}
	} else {
		// send the local mongodb data
		clientResponse, clientResponseErr := json.Marshal(buildClientResponse(findResult))
		if clientResponseErr != nil {
			log.Println(clientResponseErr)
			return []byte(common.ClientErrMessage)
		} else {
			return clientResponse
		}
	}
}

func buildClientResponse(model WordDataModel) []ClientResponse {
	var clientResponse []ClientResponse
	for _, result := range model.DictionaryEntryResponse.Results {
		var clientRes ClientResponse
		clientRes.word = result.ID

		lexicalEntries := result.LexicalEntries
		for _, lexicalEntry := range lexicalEntries {

			entries := lexicalEntry.Entries
			for _, entry := range entries {

				senses := entry.Senses
				for _, sense := range senses {

					definitions := sense.Definitions
					examples := sense.Examples
					for _, definition := range definitions {
						clientRes.definitions = append(clientRes.definitions, definition)
					}

					for _, example := range examples {
						clientRes.examples = append(clientRes.examples, example.Text)
					}
				}
			}
		}

		clientResponse = append(clientResponse, clientRes)
	}

	return clientResponse
}
