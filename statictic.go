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
