package consts

const (
	IsStart       = "INTERACTIVE_SIGNAL_START"
	IsStop        = "INTERACTIVE_SIGNAL_STOP"
	MaxPacketSize = 65536
)

const (
	CodeWelcome              = 1001
	CodeHomeOptions          = 1002
	CodeRoomList             = 1003
	CodeRoomEventJoin        = 1004
	CodeGameTypeOptions      = 1005
	CodeRoomEventCreate      = 1006
	CodeRoomEventExit        = 1007
	CodeRoomEventOffline     = 1008
	CodeRoomEventOwnerChange = 1009
)

type FacesType int

const (
	_                   FacesType = iota
	FacesBomb                     = 1 //炸彈
	FacesSingle                   = 2 //單牌
	FacesDouble                   = 3 //對子
	FacesTriple                   = 4 //三張
	FacesUnion3                   = 5 //三帶一
	FacesUnion4                   = 6 //四帶二
	FacesStraight                 = 7 //順子 or 連隊
	FacesUnion3Straight           = 8 //飛機 在跑得快下飛機可以多帶幾張 如 333444 帶 5679
	FacesInvalid                  = 9
	FacesUnion3c2                 = 10 //三帶二
	FacesUnion3c2s                = 11 //標準三帶二
)
