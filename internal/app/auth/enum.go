package auth

type Provider int

const (
	ProviderGoogle Provider = iota
	ProviderEmail
)

var keys = []string{"google", "email"}
var providers = []Provider{ProviderGoogle, ProviderEmail}

func (m Provider) String() string {
	return keys[m]
}

func ProviderFromKey(lookupKey string) *Provider {
	for i, key := range keys {
		if key == lookupKey {
			return &providers[i]
		}
	}

	return nil
}
