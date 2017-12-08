package r6

import (
	"log"
	"net/http"
	"time"
)

// R6 интерфейс для работы с game-rainbow6.ubi.com
type R6 interface {
	Test()
}

type r6api struct {
	client   *http.Client
	maxRetry int

	token string
	appID string

	ticket     string
	sessionID  string
	spaceID    string
	expiration time.Time

	loginCooldown time.Time
}

// NewByEmail создает имплементацию R6 по email/password uplay
func NewByEmail(email, password string) (R6, error) {
	return NewByToken(createTokenByEmail(email, password))
}

// NewByToken создает имплементацию R6 по uplay token
func NewByToken(token string) (R6, error) {
	r6 := &r6api{
		client:   &http.Client{},
		maxRetry: 3,

		token: token,
		appID: AppID,
	}

	return r6, nil
}

func (r6 *r6api) Test() {
	_, err := r6.getPlayers("AlexanderTzk", PlatformUplay)
	if err != nil {
		log.Println(err)
	}
	log.Println("done")
}
