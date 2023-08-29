package define

var MailPassword = "MailPassword"

type MessageStruct struct {
	Message      string `json:"message"`
	RoomIdentity string `json:"room_identity"`
}
