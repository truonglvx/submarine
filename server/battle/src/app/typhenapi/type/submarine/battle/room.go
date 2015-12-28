// This file was generated by typhen-api

package battle

import (
	"errors"
)

var _ = errors.New

// Room is kind of a TyphenAPI type.
type Room struct {
	Id      int          `codec:"id"`
	Members []RoomMember `codec:"members"`
}

// Coerce the fields.
func (t *Room) Coerce() error {
	if t.Members == nil {
		return errors.New("Members should not be empty")
	}
	return nil
}
