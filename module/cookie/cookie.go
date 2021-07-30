package cookie

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var (
	hashKey  = securecookie.GenerateRandomKey(64)
	blockKey = securecookie.GenerateRandomKey(32)
	s        = securecookie.New(hashKey, blockKey)
)

// SetCookie -
func SetCookie(key string, data interface{}) (*http.Cookie, error) {
	//encoded, err := securecookie.EncodeMulti(key, data, s)
	encoded, err := s.Encode(key, data)
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:     key,
		Value:    encoded,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}, nil
}

// ReadCookie -
func ReadCookie(cookie *http.Cookie, key string, target map[string]string) error {
	if err := s.Decode(key, cookie.Value, &target); err != nil {
		return err
	}
	//if err := securecookie.DecodeMulti(key, cookie.Value, &target, s); err != nil {
	//return err
	//}

	return nil
}
