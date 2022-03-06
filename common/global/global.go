package global

import (
	"g-server/server/ws"
)

func init() {
	RoomMap = make(map[string]*ws.Room)
}

var RoomMap map[string]*ws.Room
