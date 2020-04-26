package api

import (
	b64 "encoding/base64"
	"encoding/json"
	"go-sessioner/pkg/models"
	"net/http"
)

func createCookie(userInfo models.AuthUser) (*http.Cookie, error) {
	cookie := new(http.Cookie)
	cookie.Name = "help_saster"
	var userSessionInfo, err = json.Marshal(userInfo)
	if err != nil {
		return nil, err
	}
	cookie.Domain = "localhost"
	cookie.Value = b64.StdEncoding.EncodeToString(userSessionInfo)
	cookie.Secure = false
	return cookie, nil
}

func decodeCookie(cookie *http.Cookie) (*models.AuthUser, error) {
	value := cookie.Value
	authUser := new(models.AuthUser)
	decodedValue, _ := b64.StdEncoding.DecodeString(value)
	err := json.Unmarshal([]byte(decodedValue), authUser)
	if err != nil {
		return nil, err
	}
	return authUser, nil
}
