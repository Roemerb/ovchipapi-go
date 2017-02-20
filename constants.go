package ovchipapi

// Locale represents the locale parameter for the API. Can be EN or NL
type Locale string

const (
	clientID     = "nmOIiEJO5khvtLBK9xad3UkkS8Ua"
	clientSecret = "FE8ef6bVBiyN0NeyUJ5VOWdelvQa"

	loginURL = "https://login.ov-chipkaart.nl/oauth2/token"

	api2BaseURL     = "https://api2.ov-chipkaart.nl/femobilegateway/v1"
	authorizeURL    = api2BaseURL + "/api/authorize"
	cardsURL        = api2BaseURL + "/cards/list"
	cardURL         = api2BaseURL + "/card/"
	transactionsURL = api2BaseURL + "/transactions"

	nlNL Locale = "nl-NL"
	enUS Locale = "en-US"
)
