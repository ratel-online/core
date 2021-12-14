package consts

const (
	IsStart       = "INTERACTIVE_SIGNAL_START"
	IsStop        = "INTERACTIVE_SIGNAL_STOP"
	MaxPacketSize = 65536
)

const (
	Service   = 1
	Broadcast = 2
	Instruct  = 3

	ServiceGetRooms           = 1001
	ServiceGetRoom            = 1002
	ServiceGetRoomPlayers     = 1003
	ServiceGetGame            = 1004
	ServiceGetGameTypeOptions = 1005

	BroadcastCodeWelcome              = 1001
	BroadcastCodeHomeOptions          = 1002
	BroadcastCodeRoomList             = 1003
	BroadcastCodeRoomEventJoin        = 1004
	BroadcastCodeGameTypeOptions      = 1005
	BroadcastCodeRoomEventCreate      = 1006
	BroadcastCodeRoomEventExit        = 1007
	BroadcastCodeRoomEventOffline     = 1008
	BroadcastCodeRoomEventOwnerChange = 1009
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
