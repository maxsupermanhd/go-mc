package screen

import "github.com/Tnze/go-vmc/v767/chat"

type EventsListener struct {
	Open    func(id int, container_type int32, title chat.Message) error
	SetSlot func(id, index int) error
	Close   func(id int) error
}
