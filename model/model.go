package model

type AuthInfo struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Score int64  `json:"score"`
}

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Player struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Score  string `json:"score"`
	Group  int    `json:"group"`
	Pokers int    `json:"pokers"`
}

type Room struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Type      int64  `json:"type"`
	TypeDesc  string `json:"typeDesc"`
	Players   int    `json:"players"`
	State     int    `json:"state"`
	StateDesc string `json:"stateDesc"`
	Creator   int64  `json:"creator"`
}

type RoomList struct {
	Message
	Rooms []Room `json:"rooms"`
}

type RoomInfo struct {
	Message
	Room   Room     `json:"room"`
	Player []Player `json:"player"`
}

type RoomEvent struct {
	Message
	Room   Room   `json:"room"`
	Player Player `json:"player"`
}

type Play struct {
	Message
	Player Player `json:"player"`
	Pokers Pokers `json:"pokers"`
}

type GameEvent struct {
	Message
	Room   Room   `json:"room"`
	Player Player `json:"player"`
	Pokers Pokers `json:"pokers"`
}
