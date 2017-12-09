package r6

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	// StaticticURL ...
	StaticticURL = "https://public-ubiservices.ubi.com/v1/spaces/%s/sandboxes/%s/playerstats2/statistics?populations=%s&statistics=%s"
)

// PlayerStats текущая статистика игрока
type PlayerStats struct {
	Casual GameStats
	Ranked GameStats
}

// GameStats статистика по типу игры (казуал/ранг)
type GameStats struct {
	MatchWon    int
	MatchLost   int
	MatchPlayed int
	Kills       int
	Death       int
	TimePlayed  int
}

func (st *GameStats) fromMap(m map[string]interface{}, game string, stats []string) {
	st.MatchWon = intFromInterface(m[game+stats[0]])
	st.MatchLost = intFromInterface(m[game+stats[1]])
	st.TimePlayed = intFromInterface(m[game+stats[2]])
	st.MatchPlayed = intFromInterface(m[game+stats[3]])
	st.Kills = intFromInterface(m[game+stats[4]])
	st.Death = intFromInterface(m[game+stats[5]])
}

// PlayerStats получает текущую статистику игрока
func (pl *Player) PlayerStats() (*PlayerStats, error) {
	stats := []string{
		"matchwon", "matchlost", "timeplayed",
		"matchplayed", "kills", "death",
	}
	statsF := make([]string, 0, len(stats)*2)
	for _, game := range []string{"casualpvp_", "rankedpvp_"} {
		for _, s := range stats {
			statsF = append(statsF, game+s)
		}
	}
	m, err := pl.fetchStats(statsF...)
	if err != nil {
		return nil, errors.Wrap(err, "fetchStats")
	}
	ps := &PlayerStats{}
	ps.Casual.fromMap(m, "casualpvp_", stats)
	ps.Ranked.fromMap(m, "rankedpvp_", stats)

	return ps, nil
}

type statisticReply struct {
	Results map[string]map[string]interface{} `json:"results"`
}

func (pl *Player) fetchStats(stats ...string) (map[string]interface{}, error) {
	if len(stats) == 0 {
		return nil, nil
	}
	data, err := pl.parent.get(fmt.Sprintf(StaticticURL, pl.SpaceID, pl.PlatformURL, pl.ID, strings.Join(stats, ",")), "", true)
	if err != nil {
		return nil, errors.Wrap(err, "r6.get")
	}

	var reply statisticReply
	err = json.Unmarshal(data, &reply)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}

	s, ok := reply.Results[pl.ID]
	if !ok {
		return nil, errors.New("в ответе нет ID игрока")
	}

	result := make(map[string]interface{}, len(s))
	for k, v := range s {
		result[strings.Split(k, ":")[0]] = v
	}

	return result, nil
}
