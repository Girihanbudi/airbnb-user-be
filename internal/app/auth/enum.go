package auth

type Provider int

const (
	ProviderEmail Provider = iota
	ProviderPhone
	ProviderGoogle
)

var keys = []string{"email", "phone", "google"}
var providers = []Provider{ProviderEmail, ProviderPhone, ProviderGoogle}

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
