package word

import "gopkg.in/mgo.v2/bson"

type (
	WordRequest struct {
		Word string `json:"newWord"`
	}

	ClientResponse struct {
		Word        string   `json:"word"`
		Definitions []string `json:"definitions"`
		Examples    []string `json:"examples"`
	}

	WordDataModel struct {
		ID                      bson.ObjectId           `json:"id" bson:"_id"`
		DictionaryEntryResponse dictionaryEntryResponse `json:"dictionaryEntryResponse" bson:"dictionaryEntryResponse"`
	}

	dictionaryEntryResponse struct {
		Metadata metadata  `json:"metadata" bson:"metadata"`
		Results  []results `json:"results" bson:"results"`
	}

	pronunciations struct {
		AudioFile        string   `json:"audioFile" bson:"audioFile"`
		Dialects         []string `json:"dialects" bson:"dialects"`
		PhoneticNotation string   `json:"phoneticNotation" bson:"phoneticNotation"`
		PhoneticSpelling string   `json:"phoneticSpelling" bson:"phoneticSpelling"`
	}

	examples struct {
		Text string `json:"text" bson:"text"`
	}

	senses struct {
		Definitions []string   `json:"definitions" bson:"definitions"`
		Examples    []examples `json:"examples" bson:"examples"`
		ID          string     `json:"id" bson:"id"`
	}

	grammaticalFeatures struct {
		Text string `json:"text" bson:"text"`
		Type string `json:"type" bson:"type"`
	}

	entries struct {
		Etymologies         []string              `json:"etymologies" bson:"etymologies"`
		GrammaticalFeatures []grammaticalFeatures `json:"grammaticalFeatures" bson:"grammaticalFeatures"`
		HomographNumber     string                `json:"homographNumber" bson:"homographNumber"`
		Senses              []senses              `json:"senses" bson:"senses"`
	}

	derivatives struct {
		ID   string `json:"id" bson:"id"`
		Text string `json:"text" bson:"text"`
	}

	lexicalEntries struct {
		Derivatives     []derivatives    `json:"derivatives" bson:"derivatives"`
		Entries         []entries        `json:"entries" bson:"entries"`
		Language        string           `json:"language" bson:"language"`
		LexicalCategory string           `json:"lexicalCategory" bson:"lexicalCategory"`
		Pronunciations  []pronunciations `json:"pronunciations" bson:"pronunciations"`
		Text            string           `json:"text" bson:"text"`
	}

	results struct {
		ID             string           `json:"id" bson:"id"`
		Language       string           `json:"language" bson:"language"`
		LexicalEntries []lexicalEntries `json:"lexicalEntries" bson:"lexicalEntries"`
		Type           string           `json:"type" bson:"type"`
		Word           string           `json:"word" bson:"word"`
	}

	metadata struct {
		Provider string `json:"provider" bson:"provider"`
	}
)
