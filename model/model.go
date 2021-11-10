package model

type Player struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Score int64  `json:"score"`
	Type  int    `json:"type"`
}

type Room struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Type    int       `json:"type"`
	Players []*Player `json:"players"`
	GameID  int64     `json:"game_id"`
	Creator string    `json:"creator"`
}

type Game struct {
	ID       int64             `json:"id"`
	State    int               `json:"state"`
	Pokers   map[int64][]Poker `json:"pokers"`
	Multiple int               `json:"multiple"`
}

type Poker struct {
	ID   int `json:"id"`
	Type int `json:"type"`
}
