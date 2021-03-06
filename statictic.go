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

var (
	primaryStats = []string{
		"matchwon", "matchlost", "timeplayed",
		"matchplayed", "kills", "death",
	}
	generalStats = []string{
		"generalpvp_kills",
		"generalpvp_death",
		"generalpvp_bullethit",
		"generalpvp_bulletfired",
		"generalpvp_killassists",
		"generalpvp_revive",
		"generalpvp_headshot",
		"generalpvp_penetrationkills",
		"generalpvp_meleekills",
		"generalpvp_suicide",
		"generalpvp_barricadedeployed",
		"generalpvp_reinforcementdeploy",
		"generalpvp_dbno",
		"generalpvp_gadgetdestroy",
		"generalpvp_dbnoassists",
		"generalpvp_blindkills",
	}
)

// PlayerStats текущая статистика игрока
type PlayerStats struct {
	Casual GameStats
	Ranked GameStats
	// General GameStatsGeneral
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

func (st *GameStats) fromMap(m map[string]json.Number, game string, stats []string) {
	st.MatchWon = intFromJSONNumber(m[game+stats[0]])
	st.MatchLost = intFromJSONNumber(m[game+stats[1]])
	st.TimePlayed = intFromJSONNumber(m[game+stats[2]])
	st.MatchPlayed = intFromJSONNumber(m[game+stats[3]])
	st.Kills = intFromJSONNumber(m[game+stats[4]])
	st.Death = intFromJSONNumber(m[game+stats[5]])
}

// GameStatsGeneral статистика общая
type GameStatsGeneral struct {
	Kills            int // generalpvp_kills
	Deaths           int // generalpvp_death
	BulletHit        int // generalpvp_bullethit
	BulletFired      int // generalpvp_bulletfired
	Assists          int // generalpvp_killassists
	Revive           int // generalpvp_revive
	Headshots        int // generalpvp_headshot
	PenetrationKills int // generalpvp_penetrationkills
	MeleeKills       int // generalpvp_meleekills
	Suicide          int // generalpvp_suicide
	Barricade        int // generalpvp_barricadedeployed
	Reinforcement    int // generalpvp_reinforcementdeploy
	DBNO             int // generalpvp_dbno
	GadgetDestroy    int // generalpvp_gadgetdestroy
	DBNOAssists      int // generalpvp_dbnoassists
	BlindKills       int // generalpvp_blindkills
}

func (st *GameStatsGeneral) fromMap(m map[string]json.Number) {
	st.Kills = intFromJSONNumber(m["generalpvp_kills"])
	st.Deaths = intFromJSONNumber(m["generalpvp_death"])
	st.BulletHit = intFromJSONNumber(m["generalpvp_bullethit"])
	st.BulletFired = intFromJSONNumber(m["generalpvp_bulletfired"])
	st.Assists = intFromJSONNumber(m["generalpvp_killassists"])
	st.Revive = intFromJSONNumber(m["generalpvp_revive"])
	st.Headshots = intFromJSONNumber(m["generalpvp_headshot"])
	st.PenetrationKills = intFromJSONNumber(m["generalpvp_penetrationkills"])
	st.MeleeKills = intFromJSONNumber(m["generalpvp_meleekills"])
	st.Suicide = intFromJSONNumber(m["generalpvp_suicide"])
	st.Barricade = intFromJSONNumber(m["generalpvp_barricadedeployed"])
	st.Reinforcement = intFromJSONNumber(m["generalpvp_reinforcementdeploy"])
	st.DBNO = intFromJSONNumber(m["generalpvp_dbno"])
	st.GadgetDestroy = intFromJSONNumber(m["generalpvp_gadgetdestroy"])
	st.DBNOAssists = intFromJSONNumber(m["generalpvp_dbnoassists"])
	st.BlindKills = intFromJSONNumber(m["generalpvp_blindkills"])
}

// PlayerStats получает текущую статистику игрока
func (pl *Player) PlayerStats(ranked bool) (*GameStats, error) {
	statsF := make([]string, 0, len(primaryStats)*2)
	gameMode := "casualpvp_"
	if ranked {
		gameMode = "rankedpvp_"
	}
	for _, s := range primaryStats {
		statsF = append(statsF, gameMode+s)
	}

	m, err := pl.fetchStats(statsF...)
	if err != nil {
		return nil, errors.Wrap(err, "fetchStats")
	}
	ps := &GameStats{}
	ps.fromMap(m, gameMode, primaryStats)

	return ps, nil
}

// PlayerStatsExt возвращает расширенные статы
func (pl *Player) PlayerStatsExt() (*GameStatsGeneral, error) {
	m, err := pl.fetchStats(generalStats...)
	if err != nil {
		return nil, errors.Wrap(err, "fetchStats")
	}
	ps := &GameStatsGeneral{}
	ps.fromMap(m)
	return ps, nil
}

type statisticReply struct {
	Results map[string]map[string]json.Number `json:"results"`
}

const (
	operatorPVPStatsPrefix = "operatorpvp"
	operatorStatsTemplate  = "%s_%s" // $prefix_$stat like operatorpvp_kills to get kills for all operators
)

var (
	operatorStats = []string{
		"death",
		"kills",
		"roundlost",
		"roundwon",
		"timeplayed",
	}
)

// OperatorsStats get all operators stats
func (pl *Player) OperatorsStats() (map[string]int, error) {
	stats := make([]string, 0)
	for _, s := range operatorStats {
		stats = append(stats, fmt.Sprintf(operatorStatsTemplate, operatorPVPStatsPrefix, s))
	}
	stats = append(stats, Operators.UniqueStatsPVP()...)

	resp, err := pl.fetchStats(stats...)
	if err != nil {
		return nil, err
	}
	result := make(map[string]int, len(resp))
	for k, v := range resp {
		result[k] = intFromJSONNumber(v)
	}
	return result, nil
}

func (pl *Player) fetchStats(stats ...string) (map[string]json.Number, error) {
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

	result := make(map[string]json.Number, len(s))
	for k, v := range s {
		idx := strings.LastIndex(k, ":")
		var name string
		if idx != -1 {
			name = k[:idx]
		}
		result[name] = v
	}

	return result, nil
}
