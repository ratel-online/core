package errors

type Error struct {
	Code int
	Msg  string
}

func (e Error) Error() string {
	return e.Msg
}

func NewErr(code int, msg string) Error {
	return Error{Code: code, Msg: msg}
}

var (
	SystemError            = NewErr(500, "System error. ")
	Exist                  = NewErr(1, "Exist. ")
	ChanClosed             = NewErr(1, "Chan closed. ")
	Timeout                = NewErr(1, "Timeout. ")
	InputInvalid           = NewErr(1, "Input invalid. ")
	AuthFail               = NewErr(1, "Auth fail. ")
	RoomInvalid            = NewErr(1, "Room invalid. ")
	RoomNotInPlay          = NewErr(1, "Room are not in play. ")
	GameTypeInvalid        = NewErr(1, "Game type invalid. ")
	RoomPlayersIsFull      = NewErr(1, "Room players is fill. ")
	JoinFailForRoomRunning = NewErr(1, "Join fail, room is running. ")
	GamePlayersInvalid     = NewErr(1, "Game players invalid. ")
	PokersFacesInvalid     = NewErr(1, "Pokers faces invalid. ")
	HaveToPlay             = NewErr(1, "Have to play. ")
)
