package appcontext

const minute = 60
const hour = minute * 60
const day = hour * 24

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

const (
	AccessTokenCode      = "at"        // access token application context code
	AccessTokenDuration  = 15 * minute // expires in 15 minutes
	RefreshTokenCode     = "rt"        // refresh token application context code
	RefreshTokenDuration = 7 * day     // expires in a week
)

const (
	UserCode = "userid" // user id application context code
)
