package context

import (
	"app/typhenapi/type/submarine/battle"
)

type container struct {
	actors             []Actor
	actorsByID         map[int64]Actor
	submarinesByUserID map[int64]Actor
}

func newContainer() *container {
	c := &container{
		actors:             make([]Actor, 0),
		actorsByID:         make(map[int64]Actor),
		submarinesByUserID: make(map[int64]Actor),
	}
	return c
}

func (c *container) addActor(actor Actor) {
	c.actorsByID[actor.ID()] = actor
	c.actors = append(c.actors, actor)
	if actor.ActorType() == battle.ActorType_Submarine {
		c.submarinesByUserID[actor.UserID()] = actor
	}
}

func (c *container) removeActor(rawActor Actor) Actor {
	actor := c.actorsByID[rawActor.ID()]
	if actor == nil {
		return nil
	}

	delete(c.actorsByID, actor.ID())
	c.actors = make([]Actor, len(c.actorsByID))
	for _, a := range c.actors {
		if a != actor {
			c.actors = append(c.actors, a)
		}
	}
	if actor.ActorType() == battle.ActorType_Submarine {
		delete(c.submarinesByUserID, actor.UserID())
	}
	return actor
}
