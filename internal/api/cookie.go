package api

import (
	"encoding/json"
	"go-sessioner/pkg/models"
	"net/http"
	"time"

	guuid "github.com/google/uuid"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte(securecookie.GenerateRandomKey(32))
var blockKey = []byte(securecookie.GenerateRandomKey(32))
var s = securecookie.New(hashKey, blockKey)

func mapAuthUserToSession(authUser models.AuthUser, token string) models.Session {
	var session models.Session
	session.CreatedAt = time.Now()
	session.Token = token
	session.User = authUser.User
	session.ID = guuid.New()
	return session
}

func createCookie(session models.Session) (*http.Cookie, error) {
	var userSessionInfo, err = json.Marshal(session)
	if err != nil {
		return nil, err
	}
	value := map[string]string{
		"session": string(userSessionInfo),
	}
	if encoded, err := s.Encode("help_saster", value); err == nil {
		cookie := &http.Cookie{
			Name:     "help_saster",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}
		return cookie, nil
	} else {
		return nil, nil
	}
}

func decodeCookie(cookie *http.Cookie) (*models.Session, error) {
	value := make(map[string]string)
	session := new(models.Session)
	if err := s.Decode("help_saster", cookie.Value, &value); err == nil {
		jsonerror := json.Unmarshal([]byte(value["session"]), session)
		if jsonerror != nil {
			return nil, jsonerror
		}
		return session, nil
	} else {
		return nil, err
	}
}
