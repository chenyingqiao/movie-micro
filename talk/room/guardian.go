package room

import "sync"

var (
	guardian *ChatGuardian
	lock     *sync.Mutex
)

//ChatGuardian 聊天室守护者
type ChatGuardian struct {
	manager    GinRoomManager
	openChan   chan interface{}
	closeChan  chan interface{}
	message    chan interface{}
	deleteRoom chan interface{}
}

//NewChatGuardianSingleton
func NewChatGuardianSingleton() *ChatGuardian {
	lock = &sync.Mutex{}
	if guardian == nil {
		lock.Lock()
		guardian = &ChatGuardian{
			manager:    NewGinRoomManager(),
			openChan:   make(chan interface{}, 100),
			closeChan:  make(chan interface{}, 100),
			message:    make(chan interface{}, 100),
			deleteRoom: make(chan interface{}, 100),
		}
		guardian.Run()
		lock.Unlock()
	}
	return guardian
}

//Open Open
func (rb *ChatGuardian) Join(op RoomOperator) {
	rb.openChan <- op
}

//Close Close
func (rb *ChatGuardian) Out(op RoomOperator) {
	rb.closeChan <- op
}

//Delete Delete
func (rb *ChatGuardian) Delete(op RoomOperator) {
	rb.deleteRoom <- op
}

//Talk Talk
func (rb *ChatGuardian) Talk(message Message) {
	rb.message <- message
}

//IsInRoom isInRoom
func (rb *ChatGuardian) IsInRoom(roomID string, people *People) bool {
	room := rb.manager.Get(roomID)
	return room.IsInRoom(people)
}

//Run Run
func (rb *ChatGuardian) Run() {
	go func() {
		for {
			select {
			case operator := <-rb.openChan:
				roomOp := operator.(RoomOperator)
				room := rb.manager.Get(roomOp.RoomID)
				room.Join(roomOp.People)
			case operator := <-rb.closeChan:
				roomOp := operator.(RoomOperator)
				room := rb.manager.Get(roomOp.RoomID)
				room.Out(roomOp.People)
			case operator := <-rb.deleteRoom:
				roomOp := operator.(RoomOperator)
				rb.manager.Close(roomOp.RoomID)
			case operator := <-rb.message:
				message := operator.(Message)
				room := rb.manager.Get(message.GetRoomId())
				room.Talk(message)
			}
		}
	}()
}
