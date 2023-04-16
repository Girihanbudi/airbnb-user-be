package auth

import "time"

func Set(key string, value interface{}, expiration time.Duration) error {
	exp := expiration / 60 * time.Minute
	_, err := Cache.Set(key, value, exp).Result()
	return err
}

func Get(key string) (string, error) {
	return Cache.Get(key).Result()
}

func Del(keys ...string) error {
	_, err := Cache.Del(keys...).Result()
	return err
}
