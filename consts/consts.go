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
	FacesBomb                     = 1
	FacesSingle                   = 2
	FacesDouble                   = 3
	FacesTriple                   = 4
	FacesUnion3                   = 5
	FacesUnion4                   = 6
	FacesStraight                 = 7
	FacesUnion3Straight           = 8
	FacesInvalid                  = 9
)
