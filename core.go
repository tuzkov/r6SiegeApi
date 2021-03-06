package r6

import (
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// R6 интерфейс для работы с game-rainbow6.ubi.com
type R6 interface {
	GetPlayer(username, platform string) (*Player, error)
	GetPlayerByID(id, platform string) (*Player, error)
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

	players   map[string]*cache // platform/term:profile
	playerIDs map[string]*cache // platform/id:profile
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

		players: map[string]*cache{
			PlatformUplay: newLru(100),
			PlatformPSN:   newLru(100),
			PlatformXbox:  newLru(100),
		},
		playerIDs: map[string]*cache{
			PlatformUplay: newLru(100),
			PlatformPSN:   newLru(100),
			PlatformXbox:  newLru(100),
		},
	}

	err := r6.tryConnect()
	if err != nil {
		return nil, errors.Wrap(err, "connect")
	}

	return r6, nil
}

func (r6 *r6api) Test() {
	pl, err := r6.GetPlayer("AlexanderTzk", PlatformUplay)
	if err != nil {
		log.Fatalln(err)
	}
	s, err := pl.PlayerStats(true)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(s.MatchWon, s.MatchLost)
}
