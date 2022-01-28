package model

type AuthInfo struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Score         int64  `json:"score"`
	ClientVersion int    `json:"clientVersion"` // 客户端版本，如果服务端进行了不兼容升级，使用该字段禁止低版本客户端连接
}

type Data struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
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
	MaxPlayer int    `json:"maxPlayer"` // 该房间允许的最大人数 0无限制
	Password  string `json:"password"`  // 房间密码 默认空
}

type Options struct {
	Data
	Options []Option `json:"options"`
}

type RoomList struct {
	Data
	Rooms []Room `json:"rooms"`
}

type RoomInfo struct {
	Data
	Room    Room     `json:"room"`
	Players []Player `json:"players"`
}

type RoomEvent struct {
	Data
	Room   Room   `json:"room"`
	Player Player `json:"player"`
}

type Play struct {
	Data
	Player Player `json:"player"`
	Pokers Pokers `json:"pokers"`
}

type GameEvent struct {
	Data
	Room   Room   `json:"room"`
	Player Player `json:"player"`
	Pokers Pokers `json:"pokers"`
}
