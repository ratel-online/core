package model

import (
	"github.com/ratel-online/core/errors"
	"github.com/ratel-online/core/util/json"
)

type AuthInfo struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Score int64  `json:"score"`
}

type Req struct {
	Type int    `json:"type"`
	Data []byte `json:"data"`
}

type Resp struct {
	Type int    `json:"type"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []byte `json:"data"`
}

func ErrResp(t int, err error) Resp {
	resp := Resp{Type: t}
	if err1, ok := err.(errors.Error); ok {
		resp.Code = err1.Code
		resp.Msg = err1.Msg
	} else {
		resp.Code = errors.SystemError.Code
		resp.Msg = err.Error()
	}
	return resp
}

func SucResp(t int, data interface{}) Resp {
	return Resp{
		Type: t,
		Data: json.Marshal(data),
	}
}

type Option struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Player struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Score  int64  `json:"score"`
	Group  int    `json:"group"`
	Pokers int    `json:"pokers"`
}

type Room struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	TypeDesc  string `json:"typeDesc"`
	Players   int    `json:"players"`
	State     int    `json:"state"`
	StateDesc string `json:"stateDesc"`
	Creator   int64  `json:"creator"`
}

type Game struct {
	Players    []int64       `json:"players"`
	Pokers     map[int64]int `json:"pokers"`
	Groups     map[int64]int `json:"groups"`
	Mnemonic   map[int]int   `json:"mnemonic"`
	LastPokers []int         `json:"lastPokers"`
	LastPlayer int64         `json:"lastPlayer"`
	Universals []int         `json:"universals"`
	Additional []int         `json:"additional"`
}

type Options struct {
	Options []Option `json:"options"`
}

type RoomList struct {
	Rooms []Room `json:"rooms"`
}

type RoomInfo struct {
	Room    Room     `json:"room"`
	Players []Player `json:"players"`
}

type RoomEvent struct {
	Room   Room   `json:"room"`
	Player Player `json:"player"`
}

type Play struct {
	Player Player `json:"player"`
	Pokers Pokers `json:"pokers"`
}

type GameEvent struct {
	Room   Room   `json:"room"`
	Player Player `json:"player"`
	Pokers Pokers `json:"pokers"`
}
