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

const (
	emojiCheckMark      = "\U00002714"
	emojiCross          = "\U0000274C"
	emojiGemStone       = "\U0001F48E"
	emojiDiamondWithDot = "\U0001F4A0"
	emojiTrophy         = "\U0001F3C6"
	emojiSportsMedal    = "\U0001F3C5"
	emoji1stPlaceMedal  = "\U0001F947"
	emoji2ndPlaceMedal  = "\U0001F948"
	emoji3rdPlaceMedal  = "\U0001F949"
	emojiPoo            = "\U0001F4A9"
	emojiSkull          = "\U0001F480"
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
	return RankBracket(r.Rank, r.Season < 15)
}

func (r *PlayerRank) RankBracketEmoji() string {
	return RankBracketEmoji(r.Rank, r.Season < 15)
}

func RankBracket(rank int, old bool) string {
	if old {
		return RankBracketOld(rank)
	}
	return RankBracketNew(rank)
}

func RankBracketEmoji(rank int, old bool) string {
	if old {
		return RankBracketOldEmoji(rank)
	}
	return RankBracketNewEmoji(rank)
}

func RankBracketOld(rank int) string {
	switch {
	case rank == 0:
		return "Unranked"
	case rank < 5:
		return fmt.Sprintf("Copper %d", 5-rank)
	case rank < 9:
		return fmt.Sprintf("Bronze %d", 9-rank)
	case rank < 13:
		return fmt.Sprintf("Silver %d", 13-rank)
	case rank < 17:
		return fmt.Sprintf("Gold %d", 17-rank)
	case rank < 20:
		return fmt.Sprintf("Platinum %d", 20-rank)
	}
	return "Diamond"
}

func RankBracketOldEmoji(rank int) string {
	switch {
	case rank == 0:
		return emojiCross + "Unranked"
	case rank < 5:
		return fmt.Sprintf("%s Copper %d", emojiPoo, 5-rank)
	case rank < 9:
		return fmt.Sprintf("%s Bronze %d", emoji3rdPlaceMedal, 9-rank)
	case rank < 13:
		return fmt.Sprintf("%s Silver %d", emoji2ndPlaceMedal, 13-rank)
	case rank < 17:
		return fmt.Sprintf("%s Gold %d", emoji1stPlaceMedal, 17-rank)
	case rank < 20:
		return fmt.Sprintf("%s Platinum %d", emojiTrophy, 20-rank)
	}
	return emojiGemStone + " Diamond"
}

func RankBracketNew(rank int) string {
	switch {
	case rank == 0:
		return "Unranked"
	case rank < 6:
		return fmt.Sprintf("Copper %d", 6-rank)
	case rank < 11:
		return fmt.Sprintf("Bronze %d", 11-rank)
	case rank < 16:
		return fmt.Sprintf("Silver %d", 16-rank)
	case rank < 19:
		return fmt.Sprintf("Gold %d", 19-rank)
	case rank < 22:
		return fmt.Sprintf("Platinum %d", 22-rank)
	case rank == 22:
		return "Diamond"
	case rank > 22:
		return "Champion"
	}
	return "Unknown"
}

func RankBracketNewEmoji(rank int) string {
	switch {
	case rank == 0:
		return emojiCross + " Unranked"
	case rank < 6:
		return fmt.Sprintf("%s Copper %d", emojiPoo, 6-rank)
	case rank < 11:
		return fmt.Sprintf("%s Bronze %d", emoji3rdPlaceMedal, 11-rank)
	case rank < 16:
		return fmt.Sprintf("%s Silver %d", emoji2ndPlaceMedal, 16-rank)
	case rank < 19:
		return fmt.Sprintf("%s Gold %d", emoji1stPlaceMedal, 19-rank)
	case rank < 22:
		return fmt.Sprintf("%s Platinum %d", emojiTrophy, 22-rank)
	case rank == 22:
		return emojiGemStone + " Diamond"
	case rank > 22:
		return emojiSkull + " Champion"
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
	case 16:
		return "Shifting tides"
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
