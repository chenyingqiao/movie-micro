package room

//People 发言人
type People interface {
	GetUsername() string
	//GetChan 获取发言通道
	GetChan() *chan interface{}
}

//TalkPeople 发言人员
type TalkPeople struct {
	username string
	talkChan chan interface{}
}

//NewTalkPeople 新发言人
func NewTalkPeople(username string) *TalkPeople {
	return &TalkPeople{
		username: username,
		talkChan: make(chan interface{}),
	}
}

//GetChan GetChan
func (p TalkPeople) GetChan() *chan interface{} {
	return &p.talkChan
}

//GetUsername GetUsername
func (p TalkPeople) GetUsername() string {
	return p.username
}

type Message struct {
	roomID   string
	username string
	message  interface{}
}

func NewMessage(username string, roomID string, message string) Message {
	return Message{
		username: username,
		roomID:   roomID,
		message:  message,
	}
}

//GetMessage GetMessage
func (p *Message) GetMessage() interface{} {
	return p.message
}

//GetUsername GetMessage
func (p *Message) GetUsername() string {
	return p.username
}

//GetRoomId GetMessage
func (p *Message) GetRoomId() string {
	return p.roomID
}
