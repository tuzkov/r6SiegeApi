package r6

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type rankResponse struct {
	Players map[string]PlayerRank `json:"players"`
}

const (
	// RankURL ...
	RankURL = "https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/r6karma/players?board_id=pvp_ranked&profile_ids=%s&region_id=%s&season_id=%d"
)

// PlayerRank ранг игрока
type PlayerRank struct {
	Abandons          int     `json:"abandons"`
	BoardID           string  `json:"board_id"`
	Losses            int     `json:"losses"`
	MaxMMR            float32 `json:"max_mmr"`
	MaxRank           int     `json:"max_rank"`
	MMR               float32 `json:"mmr"`
	NextRankMMR       float32 `json:"next_rank_mmr"`
	PastSeasonAbadons int     `json:"past_seasons_abandons"`
	PastSeasonsLosses int     `json:"past_seasons_losses"`
	PastSeasonsWins   int     `json:"past_seasons_wins"`
	PreviousRankMMR   float32 `json:"previous_rank_mmr"`
	ProfileID         string  `json:"profile_id"`
	Rank              int     `json:"rank"`
	Region            string  `json:"region"`
	Season            int     `json:"season"`
	SkillMean         float32 `json:"skill_mean"`
	SkillStdev        float32 `json:"skill_stdev"`
	UpdateTime        string  `json:"update_time"`
	Wins              int     `json:"wins"`
}

// RankBracket получает ранг в human-like формате - "Золото 1", "Алмаз"
func (r *PlayerRank) RankBracket() string {
	if r.Season > 14 {
		return r.rankBracketNew()
	}
	return r.rankBracketOld()
}

func (r *PlayerRank) rankBracketOld() string {
	switch {
	case r.Rank == 0:
		return "Unranked"
	case r.Rank < 5:
		return fmt.Sprintf("%s %d", "Copper", 5-r.Rank)
	case r.Rank < 9:
		return fmt.Sprintf("%s %d", "Bronze", 9-r.Rank)
	case r.Rank < 13:
		return fmt.Sprintf("%s %d", "Silver", 13-r.Rank)
	case r.Rank < 17:
		return fmt.Sprintf("%s %d", "Gold", 17-r.Rank)
	case r.Rank < 20:
		return fmt.Sprintf("%s %d", "Platinum", 20-r.Rank)
	}
	return "Diamond"
}

func (r *PlayerRank) rankBracketNew() string {
	switch {
	case r.Rank == 0:
		return "Unranked"
	case r.Rank < 6:
		return fmt.Sprintf("%s %d", "Copper", 6-r.Rank)
	case r.Rank < 11:
		return fmt.Sprintf("%s %d", "Bronze", 11-r.Rank)
	case r.Rank < 16:
		return fmt.Sprintf("%s %d", "Silver", 16-r.Rank)
	case r.Rank < 19:
		return fmt.Sprintf("%s %d", "Gold", 19-r.Rank)
	case r.Rank < 22:
		return fmt.Sprintf("%s %d", "Platinum", 22-r.Rank)
	case r.Rank == 22:
		return "Diamond"
	case r.Rank > 22:
		return "Champion"
	}
	return "Unknown"
}

// Rank получает ранг по указанныму региону и сезоне (-1 = текущий)
func (pl *Player) Rank(region string, season int) (*PlayerRank, error) {
	data, err := pl.parent.get(fmt.Sprintf(RankURL, pl.SpaceID, pl.PlatformURL, pl.ID, region, season), "", true)
	if err != nil {
		return nil, errors.Wrap(err, "r6.get")
	}
	var rankResp rankResponse
	err = json.Unmarshal(data, &rankResp)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}

	rank := rankResp.Players[pl.ID]
	return &rank, nil
}

// SeasonNameByNum returns name of season like "Grim Sky"
func SeasonNameByNum(season int) string {
	switch season {
	case 6:
		return "Health"
	case 7:
		return "Blood Orchid"
	case 8:
		return "White Noise"
	case 9:
		return "Chimera"
	case 10:
		return "Para Bellum"
	case 11:
		return "Grim Sky"
	case 12:
		return "Wind Bastion"
	case 13:
		return "Burnt Horizon"
	case 14:
		return "Phantom Sight"
	case 15:
		return "Ember Rise"
	}
	return "Unknown"
}

// RankFromMMR returns rank from MMR
func RankFromMMR(mmr float32, season int) int {
	if season > 14 {
		return RankFromMMRNew(mmr)
	}
	return RankFromMMROld(mmr)
}

// RankFromMMRNew высчитывает ранг из MMR для сезонов 15 и дальше
func RankFromMMRNew(mmr float32) int {
	switch {
	case mmr < 1100:
		return 1
	case mmr >= 1100 && mmr < 2600:
		return int(mmr-1100)/100 + 1
	case mmr >= 2600 && mmr < 3200:
		return int(mmr-2600)/200 + 16
	case mmr >= 3200 && mmr < 4400:
		return int(mmr-3200)/400 + 19
	case mmr >= 4400 && mmr < 5000:
		return 22
	case mmr >= 5000:
		return 23
	}
	return 0
}

// RankFromMMROld высчитывает ранг из MMR для сезонов 6-14
func RankFromMMROld(mmr float32) int {
	if mmr < 1400 {
		return 1
	}
	if mmr > 4499 {
		return 20
	}
	if mmr < 2500 {
		return int((mmr-1400)/100) + 2
	}
	if mmr < 3300 {
		return int((mmr-2500)/200) + 13
	}
	return int((mmr-3300)/400) + 17
}
