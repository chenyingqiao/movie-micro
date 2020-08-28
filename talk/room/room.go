package room

import (
	"errors"

	"github.com/dustin/go-broadcast"
)

//Room room
type Room interface {
	//Join join
	Join(people *People) chan<- interface{}
	//Out out
	Out(people *People) error
	//IsInRoom 是否在房间内
	IsInRoom(people *People) bool
	//Talk talk
	Talk(data interface{})
}

//GinRoom 聊天室
type GinRoom struct {
	roomID    string
	broadcast broadcast.Broadcaster
	peoples   map[string]*People
}

//RoomOperator RoomOperator
type RoomOperator struct {
	RoomID string
	People *People
}

//NewRoomOperator NewRoomOperator
func NewRoomOperator(roomID string, people *People) RoomOperator {
	return RoomOperator{
		RoomID: roomID,
		People: people,
	}
}

//NewGinRoom room
func NewGinRoom(roomID string) GinRoom {
	return GinRoom{
		roomID:    roomID,
		broadcast: broadcast.NewBroadcaster(10),
		peoples:   map[string]*People{},
	}
}

//Join join
func (r *GinRoom) Join(people *People) error {
	if r.broadcast == nil {
		return errors.New("没有广播载体")
	}
	listener := *(*people).GetChan()
	r.broadcast.Register(listener)
	r.peoples[(*people).GetUsername()] = people
	return nil
}

//Out 退出
func (r *GinRoom) Out(people *People) error {
	if r.broadcast == nil {
		return errors.New("没有广播载体")
	}
	r.broadcast.Unregister(*(*people).GetChan())
	// close(*(*people).GetChan())
	delete(r.peoples, (*people).GetUsername())
	return nil
}

//IsInRoom 判断是否在房间内
func (r *GinRoom) IsInRoom(people *People) bool {
	if _, ok := r.peoples[(*people).GetUsername()]; ok {
		return true
	}
	return false
}

//Talk 发言
func (r *GinRoom) Talk(data interface{}) {
	r.broadcast.Submit(data)
}
