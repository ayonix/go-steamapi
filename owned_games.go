package steamapi

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Game struct {
	Appid json.Number
}

type OwnedGameResponse struct {
	Response struct {
		Games []Game
	}
}

func GetOwnedGames(id uint64, apiKey string) (games []Game, err error) {
	getOwnedGames := NewSteamMethod("IPlayerService", "GetOwnedGames", 1)

	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("format", "json")
	vals.Add("SteamId", strconv.FormatUint(id, 10))

	var resp OwnedGameResponse
	err = getOwnedGames.Request(vals, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Response.Games, nil
}
