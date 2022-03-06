package ws

import "g-server/server/model"

type Room struct {
	Uuid  string
	State string
	Count int
	Users []*model.User
	Hub   *Hub
}
