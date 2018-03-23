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
	return r.rankBracket(r.Rank)
}

func (r *PlayerRank) rankBracket(rank int) string {
	if rank == 0 {
		return "Unranked"
	} else if rank < 5 {
		return fmt.Sprintf("%s %d", "Copper", 5-rank)
	} else if rank < 9 {
		return fmt.Sprintf("%s %d", "Bronze", 9-rank)
	} else if rank < 13 {
		return fmt.Sprintf("%s %d", "Silver", 13-rank)
	} else if rank < 17 {
		return fmt.Sprintf("%s %d", "Gold", 17-rank)
	} else if rank < 20 {
		return fmt.Sprintf("%s %d", "Platinum", 20-rank)
	}
	return "Diamond"
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

// RankFromMMR высчитывает ранг из MMR
func RankFromMMR(mmr float32) int {
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
