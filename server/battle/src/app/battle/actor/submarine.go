package actor

import (
	"app/battle/context"
	"app/battle/event"
	"app/typhenapi/type/submarine/battle"
)

type submarine struct {
	*actor
}

// NewSubmarine creates a submarine.
func NewSubmarine(battleContext *context.Context, userID int64) context.Actor {
	s := &submarine{
		actor: newActor(battleContext, userID, battle.ActorType_Submarine),
	}
	s.event.On(event.AccelerationRequest, s.onAccelerationRequest)
	s.event.On(event.BrakeRequest, s.onBrakeRequest)
	s.event.On(event.TurnRequest, s.onTurnRequest)
	s.context.Event.Emit(event.ActorCreate, s)
	return s
}

func (s *submarine) onAccelerationRequest(message *battle.AccelerationRequestObject) {
	s.accelerate()
}

func (s *submarine) onBrakeRequest(message *battle.BrakeRequestObject) {
	s.brake()
}

func (s *submarine) onTurnRequest(message *battle.TurnRequestObject) {
	s.turn(message.Direction)
}
