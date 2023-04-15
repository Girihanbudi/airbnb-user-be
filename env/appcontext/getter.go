package appcontext

import "context"

func GetLocale(ctx context.Context) string {
	value := ctx.Value(LocaleCode)
	if value == nil {
		return LocaleDefault
	}
	return value.(string)
}

func GetCurrency(ctx context.Context) string {
	value := ctx.Value(CurrencyCode)
	if value == nil {
		return CurrencyDefault
	}
	return value.(string)
}

func GetAccessToken(ctx context.Context) *string {
	value := ctx.Value(AccessTokenCode)
	if value == nil {
		return nil
	}
	token := value.(string)

	return &token
}

func GetUserId(ctx context.Context) *string {
	value := ctx.Value(UserCode)
	if value == nil {
		return nil
	}
	userId := value.(string)

	return &userId
}
