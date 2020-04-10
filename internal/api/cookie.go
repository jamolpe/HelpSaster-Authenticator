package api

import (
	"go-sessioner/pkg/models"
	b64 "encoding/base64"
	"encoding/json"
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
