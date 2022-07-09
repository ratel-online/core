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
	FacesUnion3c2                 = 10 //跑得快三帶二
	FacesUnion3c2s                = 11 //跑得快標準三帶二
	FacesUnion4C3                 = 12 //跑得快四帶3
	FacesUnion4C3s                = 13 //跑得快標準四帶3
	FacesUnion3c2C                = 14 //跑得快二連三帶2
	FacesUnion3c2Cs               = 15 //跑得快標準二連三帶2
	FacesUnion3c2CM               = 16 //跑得快三連三帶2
	FacesUnion3c2CsM              = 17 //跑得快標準三連三帶2
	FacesDoubles                  = 18 //跑得快連隊
	FacesStraights                = 19 //跑得快順子
	FacesThreeStraights           = 20 //跑得快三連隊
)
