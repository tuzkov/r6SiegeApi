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
		log.Println(err)
		return
	}
	stats, err := pl.PlayerStats(true)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Ranked")
	log.Println("MatchWon", stats.Ranked.MatchWon)
	log.Println("MatchLost", stats.Ranked.MatchLost)
	log.Println("MatchPlayed", stats.Ranked.MatchPlayed)
	log.Println("Kills", stats.Ranked.Kills)
	log.Println("Death", stats.Ranked.Death)
	log.Println("TimePlayed", stats.Ranked.TimePlayed)

	log.Println("Casual")
	log.Println("MatchWon", stats.Casual.MatchWon)
	log.Println("MatchLost", stats.Casual.MatchLost)
	log.Println("MatchPlayed", stats.Casual.MatchPlayed)
	log.Println("Kills", stats.Casual.Kills)
	log.Println("Death", stats.Casual.Death)
	log.Println("TimePlayed", stats.Casual.TimePlayed)

	log.Println("Kills", stats.General.Kills)
	log.Println("Deaths", stats.General.Deaths)
	log.Println("BulletHit", stats.General.BulletHit)
	log.Println("BulletFired", stats.General.BulletFired)
	log.Println("Assists", stats.General.Assists)
	log.Println("Revive", stats.General.Revive)
	log.Println("Headshots", stats.General.Headshots)
	log.Println("PenetrationKills", stats.General.PenetrationKills)
	log.Println("MeleeKills", stats.General.MeleeKills)
	log.Println("Suicide", stats.General.Suicide)
	log.Println("Barricade", stats.General.Barricade)
	log.Println("Reinforcement", stats.General.Reinforcement)
	log.Println("DBNO", stats.General.DBNO)
	log.Println("GadgetDestroy", stats.General.GadgetDestroy)
	log.Println("DBNOAssists", stats.General.DBNOAssists)

	log.Println("done")
}
