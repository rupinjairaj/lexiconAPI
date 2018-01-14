package common

const (
	//Oxford Dictionary API constants

	// BaseURL :
	BaseURL = "https://od-api.oxforddictionaries.com/api/v1/entries/"
	// SourceLang :
	SourceLang = "en/"
	// AppIDHeader :
	AppIDHeader = "app_id"
	// AppIDValue :
	AppIDValue = ""
	// AppKeyHeader :
	AppKeyHeader = "app_key"
	// AppKeyValue :
	AppKeyValue = ""

	// mongodb connection constant

	// SERVER : mongodb server connection
	SERVER = "localhost:27017"

	// DBNAME : the database we will use
	// is lexicon
	DBNAME = "lexicon"

	// COLLECTION : the name of the
	// collection we will store our
	// word documents in
	COLLECTION = "words"

	// Application error messages

	// ClientErrMessage is the generic error message
	// for the client. Devs should check the log messages.
	ClientErrMessage = "Oops! :( \nLooks like something went wrong."
)
