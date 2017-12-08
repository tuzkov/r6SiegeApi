package r6

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pkg/errors"
)

const (
	getPlayersURL = "https://public-ubiservices.ubi.com/v2/profiles?nameOnPlatform=%s&platformType=%s"
)

var (
	// ErrNoProfiles ...
	ErrNoProfiles = errors.New("не найдено профилей игроков")
)

// Player данные о игроке
type Player struct {
	parent *r6api

	ID           string
	UserID       string
	Platform     string
	PlatformURL  string
	SpaceID      string
	IDOnPlatform string
	Name         string

	URL     string
	IconURL string
}

func (r6 *r6api) newPlayer(profile getPlayerProfile) *Player {
	return &Player{
		parent: r6,

		ID:           profile.ProfileID,
		UserID:       profile.UserID,
		Platform:     profile.PlatformType,
		PlatformURL:  PlatformURL[profile.PlatformType],
		SpaceID:      SpaceID[profile.PlatformType],
		IDOnPlatform: profile.IDOnPlatform,
		Name:         profile.NameOnPlatform,

		URL: fmt.Sprintf("https://game-rainbow6.ubi.com/en-us/%s/player-statistics/%s/multiplayer",
			profile.PlatformType, profile.ProfileID),
		IconURL: fmt.Sprintf("https://ubisoft-avatars.akamaized.net/%s/default_146_146.png", profile.ProfileID),
	}
}

// getPlayersResponse ответ API getPlayers
type getPlayersResponse struct {
	Profiles []getPlayerProfile `json:"profiles"`
}

// getPlayerProfile профиль игрока от API getPlayers
type getPlayerProfile struct {
	ProfileID      string `json:"profileId"`
	UserID         string `json:"userId"`
	PlatformType   string `json:"platformType"`
	IDOnPlatform   string `json:"idOnPlatform"`
	NameOnPlatform string `json:"nameOnPlatform"`
}

// getPlayers получает список игроков по имени
func (r6 *r6api) getPlayers(term, platform string) (result []getPlayerProfile, err error) {
	data, err := r6.get(fmt.Sprintf(getPlayersURL, url.QueryEscape(term), url.QueryEscape(platform)), "", true)
	if err != nil {
		return nil, errors.Wrap(err, "r6.get")
	}
	var resp getPlayersResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}

	result = make([]getPlayerProfile, 0, len(resp.Profiles))
	for _, profile := range resp.Profiles {
		if profile.PlatformType == platform {
			result = append(result, profile)
		}
	}
	if len(result) == 0 {
		return nil, ErrNoProfiles
	}

	return result, nil
}

// GetPlayer получает первого игрока из списка с таким именем
func (r6 *r6api) GetPlayer(username, platform string) (*Player, error) {
	if c, ok := r6.players[platform]; ok {
		if tt, ok := c.Get(username); ok {
			if player, ok := tt.(*Player); ok {
				return player, nil
			}
		}
	}
	result, err := r6.getPlayers(username, platform)
	if err != nil {
		return nil, err
	}
	player := r6.newPlayer(result[0])
	if c, ok := r6.players[platform]; ok {
		c.AddWithExpiresInSecs(username, player, 24*60*60)
	}
	return player, nil
}
