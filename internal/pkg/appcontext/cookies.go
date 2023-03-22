package appcontext

const day = 60 * 60 * 24

const (
	LocaleCode     = "locale" // locale application context code
	LocaleDuration = 7 * day  // expires in a week
	LocaleDefault  = "en-US"  // default language using English US
)

const (
	CurrencyCode     = "currency" // currency application context code
	CurrencyDuration = 7 * day    // expires in a week
	CurrencyDefault  = "USD"      // default currency using United State Dollar
)

const (
	OauthCode     = "oauth"   // oauth application context code
	OauthDuration = 365 * day // expires in a week
)
