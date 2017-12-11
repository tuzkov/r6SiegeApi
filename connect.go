package r6

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	// ConnectURL для подключения к серверу Ubi
	ConnectURL = "https://connect.ubi.com/ubiservices/v2/profiles/sessions"
)

var (
	// ErrConnect ...
	ErrConnect = errors.New("не удалось подключиться к Ubi")
	// ErrLoginCooldown ...
	ErrLoginCooldown = errors.New("login cooldown")
)

// connectRequest тело запроса к API connect
type connectRequest struct {
	RememberMe bool `json:"rememberMe"`
}

// connectResponse ответ API connect
type connectResponse struct {
	Token      string `json:"token"`
	Ticket     string `json:"ticket"`
	Expiration string `json:"expiration"`
	SpaceID    string `json:"spaceId"`
	SessionID  string `json:"sessionId"`
}

func (r6 *r6api) needConnect() bool {
	return r6.ticket == "" || r6.expiration.Before(time.Now())
}

// connect производит подключение/авторизацию в Uplay
func (r6 *r6api) connect() error {
	if time.Now().Before(r6.loginCooldown) {
		return ErrLoginCooldown
	}

	body, err := json.Marshal(connectRequest{true})
	if err != nil {
		return errors.Wrap(err, "json.Marshal")
	}
	req, err := http.NewRequest(http.MethodPost, ConnectURL, bytes.NewReader(body))
	if err != nil {
		return errors.Wrap(err, "http.NewRequest")
	}

	req.Header.Add(HeaderContentType, ContentTypeJSON)
	req.Header.Add(HeaderUbiAppID, r6.appID)
	req.Header.Add(HeaderAuth, fmt.Sprintf("Basic %s", r6.token))

	resp, err := r6.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "client.Do")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return ErrLoginIncorrect
	}

	if 200 > resp.StatusCode || resp.StatusCode > 299 {
		return errors.Errorf("response status is not ok - %s", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "ReadAll body")
	}

	var cResp connectResponse
	err = json.Unmarshal(respBody, &cResp)
	if err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	if cResp.Ticket == "" {
		return ErrConnect
	}
	r6.ticket = cResp.Ticket
	r6.sessionID = cResp.SessionID
	r6.spaceID = cResp.SpaceID
	r6.expiration, err = time.Parse(time.RFC3339, cResp.Expiration)
	if err != nil {
		return errors.Wrap(err, "time.Parse")
	}
	r6.expiration = r6.expiration.Add(-time.Minute)

	return nil
}

// tryConnect пробует подключиться несколько раз подряд
func (r6 *r6api) tryConnect() (err error) {
	for i := 0; i < r6.maxRetry; i++ {
		err = r6.connect()
		if err != nil {
			log.Printf("r6: Ошибка при connect, попытка %d/%d: %s", i, r6.maxRetry, err.Error())
			if err == ErrLoginCooldown || err == ErrLoginIncorrect {
				return
			}
			continue
		}
		break
	}
	if err != nil {
		r6.loginCooldown = time.Now().Add(time.Minute)
	}
	return err
}

// get делает запрос в Uplay с авторизацией и нужными заголовками
func (r6 *r6api) get(url string, referer string, isJSON bool) (message []byte, err error) {
	if r6.needConnect() {
		err = r6.tryConnect()
		if err != nil {
			return
		}
	}

	req, err := r6.newRequest(http.MethodGet, url, nil, referer)
	if err != nil {
		return nil, errors.Wrap(err, "r6.newRequest")
	}

	resp, err := r6.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client.Do")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}

	if isJSON {
		var tempJ map[string]interface{}
		err = json.Unmarshal(body, &tempJ)
		if err != nil {
			return nil, errors.Wrap(err, "json.Unmarshal")
		}
		if httpCode, ok := tempJ["httpCode"]; ok {
			switch httpCode {
			case 401:
				err = r6.tryConnect()
				if err != nil {
					return
				}
			case 404:
				return nil, errors.New("missing resourse")
			}
		}
		return body, nil
	}
	return body, nil
}
