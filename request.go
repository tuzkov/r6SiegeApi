package r6

import (
	"bytes"
	"fmt"	
	"net/http"

	"github.com/pkg/errors"
)

const (
	// ContentTypeJSON ...
	ContentTypeJSON = "application/json"
)

const (
	// HeaderContentType ...
	HeaderContentType = "Content-Type"
	// HeaderUbiAppID ...
	HeaderUbiAppID = "Ubi-AppId"
	// HeaderAuth ...
	HeaderAuth = "Authorization"
	// HeaderSessionID ...
	HeaderSessionID = "Ubi-SessionId"
	// HeaderConnection ...
	HeaderConnection = "Connection"
	// HeaderReferer ...
	HeaderReferer = "Referer"
)

// newRequest создает запрос с требуемыми заголовками
func (r6 *r6api) newRequest(method, url string, body []byte, referer string) (*http.Request, error) {
	r, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest")
	}
	r.Header.Add(HeaderAuth, fmt.Sprintf("Ubi_v1 t=%s", r6.ticket))
	r.Header.Add(HeaderUbiAppID, r6.appID)
	r.Header.Add(HeaderSessionID, r6.sessionID)
	r.Header.Add(HeaderConnection, "keep-alive")
	if referer != "" {
		r.Header.Add(HeaderReferer, fmt.Sprintf("https://game-rainbow6.ubi.com/en-gb/uplay/player-statistics/%s/multiplayer", referer))
	}

	return r, nil
}
