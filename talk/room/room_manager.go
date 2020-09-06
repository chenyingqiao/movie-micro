package room

import "errors"

//RoomManager 聊天室管理
type RoomManager interface {
	//Get 获取聊天室
	Get(roomID string) *Room
	//Close 关闭聊天室
	Close(roomID string) error
	GetAll() map[string]*GinRoom
}

//GinRoomManager 聊天室管理
type GinRoomManager struct {
	rooms map[string]*GinRoom
}

//NewGinRoomManager NewGinRoomManager
func NewGinRoomManager() GinRoomManager {
	manager := GinRoomManager{
		rooms: map[string]*GinRoom{},
	}
	return manager
}

//Create Get
func (r *GinRoomManager) Get(roomID string) *GinRoom {
	if room, ok := r.rooms[roomID]; ok {
		return room
	}
	room := NewGinRoom(roomID)
	r.rooms[roomID] = &room
	return r.rooms[roomID]
}

//Close Close
func (r *GinRoomManager) Close(roomID string) error {
	if room, ok := r.rooms[roomID]; ok {
		err := room.broadcast.Close()
		if err != nil {
			return err
		}
		delete(r.rooms, roomID)
		return nil
	} else {
		return errors.New("roomID不存在")
	}
}

//Heartbeat 所有链接维持一个心跳
func (r *GinRoomManager) GetAll() map[string]*GinRoom {
	return r.rooms
}
