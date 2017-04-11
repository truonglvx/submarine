// Generated by: genem (https://github.com/shiwano/genem)

package scene

import (
	"reflect"
	"sync"

	"github.com/ungerik/go3d/float64/vec2"

	battleAPI "github.com/shiwano/submarine/server/battle/lib/typhenapi/type/submarine/battle"
)

// ActorEventEmitter represents an event emitter.
type ActorEventEmitter struct {
	collideWithStageEventMu            *sync.Mutex
	collideWithStageEventListeners     []func(point vec2.T)
	collideWithStageEventListenersOnce []func(point vec2.T)

	collideWithOtherActorEventMu            *sync.Mutex
	collideWithOtherActorEventListeners     []func(actor Actor, point vec2.T)
	collideWithOtherActorEventListenersOnce []func(actor Actor, point vec2.T)

	accelerationRequestEventMu            *sync.Mutex
	accelerationRequestEventListeners     []func(message *battleAPI.AccelerationRequest)
	accelerationRequestEventListenersOnce []func(message *battleAPI.AccelerationRequest)

	brakeRequestEventMu            *sync.Mutex
	brakeRequestEventListeners     []func(message *battleAPI.BrakeRequest)
	brakeRequestEventListenersOnce []func(message *battleAPI.BrakeRequest)

	turnRequestEventMu            *sync.Mutex
	turnRequestEventListeners     []func(message *battleAPI.TurnRequest)
	turnRequestEventListenersOnce []func(message *battleAPI.TurnRequest)

	torpedoRequestEventMu            *sync.Mutex
	torpedoRequestEventListeners     []func(message *battleAPI.TorpedoRequest)
	torpedoRequestEventListenersOnce []func(message *battleAPI.TorpedoRequest)

	pingerRequestEventMu            *sync.Mutex
	pingerRequestEventListeners     []func(message *battleAPI.PingerRequest)
	pingerRequestEventListenersOnce []func(message *battleAPI.PingerRequest)

	watcherRequestEventMu            *sync.Mutex
	watcherRequestEventListeners     []func(message *battleAPI.WatcherRequest)
	watcherRequestEventListenersOnce []func(message *battleAPI.WatcherRequest)

	userLeaveEventMu            *sync.Mutex
	userLeaveEventListeners     []func()
	userLeaveEventListenersOnce []func()
}

// NewActorEventEmitter creates an event emitter.
func NewActorEventEmitter() *ActorEventEmitter {
	return &ActorEventEmitter{

		collideWithStageEventMu: new(sync.Mutex),

		collideWithOtherActorEventMu: new(sync.Mutex),

		accelerationRequestEventMu: new(sync.Mutex),

		brakeRequestEventMu: new(sync.Mutex),

		turnRequestEventMu: new(sync.Mutex),

		torpedoRequestEventMu: new(sync.Mutex),

		pingerRequestEventMu: new(sync.Mutex),

		watcherRequestEventMu: new(sync.Mutex),

		userLeaveEventMu: new(sync.Mutex),
	}
}

// EmitCollideWithStageEvent emits the specified event.
func (_e *ActorEventEmitter) EmitCollideWithStageEvent(point vec2.T) {
	_e.collideWithStageEventMu.Lock()
	listeners := make([]func(point vec2.T), len(_e.collideWithStageEventListeners))
	copy(listeners, _e.collideWithStageEventListeners)
	listenersOnce := _e.collideWithStageEventListenersOnce
	_e.collideWithStageEventListenersOnce = make([]func(point vec2.T), 0)
	_e.collideWithStageEventMu.Unlock()
	for _, l := range listeners {
		l(point)
	}
	for _, l := range listenersOnce {
		l(point)
	}
}

// AddCollideWithStageEventListener registers the specified event listener.
func (_e *ActorEventEmitter) AddCollideWithStageEventListener(listener func(point vec2.T)) {
	_e.collideWithStageEventMu.Lock()
	_e.collideWithStageEventListeners = append(_e.collideWithStageEventListeners, listener)
	_e.collideWithStageEventMu.Unlock()
}

// AddCollideWithStageEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *ActorEventEmitter) AddCollideWithStageEventListenerOnce(listener func(point vec2.T)) {
	_e.collideWithStageEventMu.Lock()
	_e.collideWithStageEventListenersOnce = append(_e.collideWithStageEventListenersOnce, listener)
	_e.collideWithStageEventMu.Unlock()
}

// RemoveCollideWithStageEventListener removes the event listener previously registered.
func (_e *ActorEventEmitter) RemoveCollideWithStageEventListener(listener func(point vec2.T)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.collideWithStageEventMu.Lock()
	listeners := _e.collideWithStageEventListeners[:0]
	for _, l := range _e.collideWithStageEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.collideWithStageEventListeners = listeners
	listenersOnce := _e.collideWithStageEventListenersOnce[:0]
	for _, l := range _e.collideWithStageEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.collideWithStageEventListenersOnce = listenersOnce
	_e.collideWithStageEventMu.Unlock()
}

// EmitCollideWithOtherActorEvent emits the specified event.
func (_e *ActorEventEmitter) EmitCollideWithOtherActorEvent(actor Actor, point vec2.T) {
	_e.collideWithOtherActorEventMu.Lock()
	listeners := make([]func(actor Actor, point vec2.T), len(_e.collideWithOtherActorEventListeners))
	copy(listeners, _e.collideWithOtherActorEventListeners)
	listenersOnce := _e.collideWithOtherActorEventListenersOnce
	_e.collideWithOtherActorEventListenersOnce = make([]func(actor Actor, point vec2.T), 0)
	_e.collideWithOtherActorEventMu.Unlock()
	for _, l := range listeners {
		l(actor, point)
	}
	for _, l := range listenersOnce {
		l(actor, point)
	}
}

// AddCollideWithOtherActorEventListener registers the specified event listener.
func (_e *ActorEventEmitter) AddCollideWithOtherActorEventListener(listener func(actor Actor, point vec2.T)) {
	_e.collideWithOtherActorEventMu.Lock()
	_e.collideWithOtherActorEventListeners = append(_e.collideWithOtherActorEventListeners, listener)
	_e.collideWithOtherActorEventMu.Unlock()
}

// AddCollideWithOtherActorEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *ActorEventEmitter) AddCollideWithOtherActorEventListenerOnce(listener func(actor Actor, point vec2.T)) {
	_e.collideWithOtherActorEventMu.Lock()
	_e.collideWithOtherActorEventListenersOnce = append(_e.collideWithOtherActorEventListenersOnce, listener)
	_e.collideWithOtherActorEventMu.Unlock()
}

// RemoveCollideWithOtherActorEventListener removes the event listener previously registered.
func (_e *ActorEventEmitter) RemoveCollideWithOtherActorEventListener(listener func(actor Actor, point vec2.T)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.collideWithOtherActorEventMu.Lock()
	listeners := _e.collideWithOtherActorEventListeners[:0]
	for _, l := range _e.collideWithOtherActorEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.collideWithOtherActorEventListeners = listeners
	listenersOnce := _e.collideWithOtherActorEventListenersOnce[:0]
	for _, l := range _e.collideWithOtherActorEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.collideWithOtherActorEventListenersOnce = listenersOnce
	_e.collideWithOtherActorEventMu.Unlock()
}

// EmitAccelerationRequestEvent emits the specified event.
func (_e *ActorEventEmitter) EmitAccelerationRequestEvent(message *battleAPI.AccelerationRequest) {
	_e.accelerationRequestEventMu.Lock()
	listeners := make([]func(message *battleAPI.AccelerationRequest), len(_e.accelerationRequestEventListeners))
	copy(listeners, _e.accelerationRequestEventListeners)
	listenersOnce := _e.accelerationRequestEventListenersOnce
	_e.accelerationRequestEventListenersOnce = make([]func(message *battleAPI.AccelerationRequest), 0)
	_e.accelerationRequestEventMu.Unlock()
	for _, l := range listeners {
		l(message)
	}
	for _, l := range listenersOnce {
		l(message)
	}
}

// AddAccelerationRequestEventListener registers the specified event listener.
func (_e *ActorEventEmitter) AddAccelerationRequestEventListener(listener func(message *battleAPI.AccelerationRequest)) {
	_e.accelerationRequestEventMu.Lock()
	_e.accelerationRequestEventListeners = append(_e.accelerationRequestEventListeners, listener)
	_e.accelerationRequestEventMu.Unlock()
}

// AddAccelerationRequestEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *ActorEventEmitter) AddAccelerationRequestEventListenerOnce(listener func(message *battleAPI.AccelerationRequest)) {
	_e.accelerationRequestEventMu.Lock()
	_e.accelerationRequestEventListenersOnce = append(_e.accelerationRequestEventListenersOnce, listener)
	_e.accelerationRequestEventMu.Unlock()
}

// RemoveAccelerationRequestEventListener removes the event listener previously registered.
func (_e *ActorEventEmitter) RemoveAccelerationRequestEventListener(listener func(message *battleAPI.AccelerationRequest)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.accelerationRequestEventMu.Lock()
	listeners := _e.accelerationRequestEventListeners[:0]
	for _, l := range _e.accelerationRequestEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.accelerationRequestEventListeners = listeners
	listenersOnce := _e.accelerationRequestEventListenersOnce[:0]
	for _, l := range _e.accelerationRequestEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.accelerationRequestEventListenersOnce = listenersOnce
	_e.accelerationRequestEventMu.Unlock()
}

// EmitBrakeRequestEvent emits the specified event.
func (_e *ActorEventEmitter) EmitBrakeRequestEvent(message *battleAPI.BrakeRequest) {
	_e.brakeRequestEventMu.Lock()
	listeners := make([]func(message *battleAPI.BrakeRequest), len(_e.brakeRequestEventListeners))
	copy(listeners, _e.brakeRequestEventListeners)
	listenersOnce := _e.brakeRequestEventListenersOnce
	_e.brakeRequestEventListenersOnce = make([]func(message *battleAPI.BrakeRequest), 0)
	_e.brakeRequestEventMu.Unlock()
	for _, l := range listeners {
		l(message)
	}
	for _, l := range listenersOnce {
		l(message)
	}
}

// AddBrakeRequestEventListener registers the specified event listener.
func (_e *ActorEventEmitter) AddBrakeRequestEventListener(listener func(message *battleAPI.BrakeRequest)) {
	_e.brakeRequestEventMu.Lock()
	_e.brakeRequestEventListeners = append(_e.brakeRequestEventListeners, listener)
	_e.brakeRequestEventMu.Unlock()
}

// AddBrakeRequestEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *ActorEventEmitter) AddBrakeRequestEventListenerOnce(listener func(message *battleAPI.BrakeRequest)) {
	_e.brakeRequestEventMu.Lock()
	_e.brakeRequestEventListenersOnce = append(_e.brakeRequestEventListenersOnce, listener)
	_e.brakeRequestEventMu.Unlock()
}

// RemoveBrakeRequestEventListener removes the event listener previously registered.
func (_e *ActorEventEmitter) RemoveBrakeRequestEventListener(listener func(message *battleAPI.BrakeRequest)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.brakeRequestEventMu.Lock()
	listeners := _e.brakeRequestEventListeners[:0]
	for _, l := range _e.brakeRequestEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.brakeRequestEventListeners = listeners
	listenersOnce := _e.brakeRequestEventListenersOnce[:0]
	for _, l := range _e.brakeRequestEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.brakeRequestEventListenersOnce = listenersOnce
	_e.brakeRequestEventMu.Unlock()
}

// EmitTurnRequestEvent emits the specified event.
func (_e *ActorEventEmitter) EmitTurnRequestEvent(message *battleAPI.TurnRequest) {
	_e.turnRequestEventMu.Lock()
	listeners := make([]func(message *battleAPI.TurnRequest), len(_e.turnRequestEventListeners))
	copy(listeners, _e.turnRequestEventListeners)
	listenersOnce := _e.turnRequestEventListenersOnce
	_e.turnRequestEventListenersOnce = make([]func(message *battleAPI.TurnRequest), 0)
	_e.turnRequestEventMu.Unlock()
	for _, l := range listeners {
		l(message)
	}
	for _, l := range listenersOnce {
		l(message)
	}
}

// AddTurnRequestEventListener registers the specified event listener.
func (_e *ActorEventEmitter) AddTurnRequestEventListener(listener func(message *battleAPI.TurnRequest)) {
	_e.turnRequestEventMu.Lock()
	_e.turnRequestEventListeners = append(_e.turnRequestEventListeners, listener)
	_e.turnRequestEventMu.Unlock()
}

// AddTurnRequestEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *ActorEventEmitter) AddTurnRequestEventListenerOnce(listener func(message *battleAPI.TurnRequest)) {
	_e.turnRequestEventMu.Lock()
	_e.turnRequestEventListenersOnce = append(_e.turnRequestEventListenersOnce, listener)
	_e.turnRequestEventMu.Unlock()
}

// RemoveTurnRequestEventListener removes the event listener previously registered.
func (_e *ActorEventEmitter) RemoveTurnRequestEventListener(listener func(message *battleAPI.TurnRequest)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.turnRequestEventMu.Lock()
	listeners := _e.turnRequestEventListeners[:0]
	for _, l := range _e.turnRequestEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.turnRequestEventListeners = listeners
	listenersOnce := _e.turnRequestEventListenersOnce[:0]
	for _, l := range _e.turnRequestEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.turnRequestEventListenersOnce = listenersOnce
	_e.turnRequestEventMu.Unlock()
}

// EmitTorpedoRequestEvent emits the specified event.
func (_e *ActorEventEmitter) EmitTorpedoRequestEvent(message *battleAPI.TorpedoRequest) {
	_e.torpedoRequestEventMu.Lock()
	listeners := make([]func(message *battleAPI.TorpedoRequest), len(_e.torpedoRequestEventListeners))
	copy(listeners, _e.torpedoRequestEventListeners)
	listenersOnce := _e.torpedoRequestEventListenersOnce
	_e.torpedoRequestEventListenersOnce = make([]func(message *battleAPI.TorpedoRequest), 0)
	_e.torpedoRequestEventMu.Unlock()
	for _, l := range listeners {
		l(message)
	}
	for _, l := range listenersOnce {
		l(message)
	}
}

// AddTorpedoRequestEventListener registers the specified event listener.
func (_e *ActorEventEmitter) AddTorpedoRequestEventListener(listener func(message *battleAPI.TorpedoRequest)) {
	_e.torpedoRequestEventMu.Lock()
	_e.torpedoRequestEventListeners = append(_e.torpedoRequestEventListeners, listener)
	_e.torpedoRequestEventMu.Unlock()
}

// AddTorpedoRequestEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *ActorEventEmitter) AddTorpedoRequestEventListenerOnce(listener func(message *battleAPI.TorpedoRequest)) {
	_e.torpedoRequestEventMu.Lock()
	_e.torpedoRequestEventListenersOnce = append(_e.torpedoRequestEventListenersOnce, listener)
	_e.torpedoRequestEventMu.Unlock()
}

// RemoveTorpedoRequestEventListener removes the event listener previously registered.
func (_e *ActorEventEmitter) RemoveTorpedoRequestEventListener(listener func(message *battleAPI.TorpedoRequest)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.torpedoRequestEventMu.Lock()
	listeners := _e.torpedoRequestEventListeners[:0]
	for _, l := range _e.torpedoRequestEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.torpedoRequestEventListeners = listeners
	listenersOnce := _e.torpedoRequestEventListenersOnce[:0]
	for _, l := range _e.torpedoRequestEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.torpedoRequestEventListenersOnce = listenersOnce
	_e.torpedoRequestEventMu.Unlock()
}

// EmitPingerRequestEvent emits the specified event.
func (_e *ActorEventEmitter) EmitPingerRequestEvent(message *battleAPI.PingerRequest) {
	_e.pingerRequestEventMu.Lock()
	listeners := make([]func(message *battleAPI.PingerRequest), len(_e.pingerRequestEventListeners))
	copy(listeners, _e.pingerRequestEventListeners)
	listenersOnce := _e.pingerRequestEventListenersOnce
	_e.pingerRequestEventListenersOnce = make([]func(message *battleAPI.PingerRequest), 0)
	_e.pingerRequestEventMu.Unlock()
	for _, l := range listeners {
		l(message)
	}
	for _, l := range listenersOnce {
		l(message)
	}
}

// AddPingerRequestEventListener registers the specified event listener.
func (_e *ActorEventEmitter) AddPingerRequestEventListener(listener func(message *battleAPI.PingerRequest)) {
	_e.pingerRequestEventMu.Lock()
	_e.pingerRequestEventListeners = append(_e.pingerRequestEventListeners, listener)
	_e.pingerRequestEventMu.Unlock()
}

// AddPingerRequestEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *ActorEventEmitter) AddPingerRequestEventListenerOnce(listener func(message *battleAPI.PingerRequest)) {
	_e.pingerRequestEventMu.Lock()
	_e.pingerRequestEventListenersOnce = append(_e.pingerRequestEventListenersOnce, listener)
	_e.pingerRequestEventMu.Unlock()
}

// RemovePingerRequestEventListener removes the event listener previously registered.
func (_e *ActorEventEmitter) RemovePingerRequestEventListener(listener func(message *battleAPI.PingerRequest)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.pingerRequestEventMu.Lock()
	listeners := _e.pingerRequestEventListeners[:0]
	for _, l := range _e.pingerRequestEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.pingerRequestEventListeners = listeners
	listenersOnce := _e.pingerRequestEventListenersOnce[:0]
	for _, l := range _e.pingerRequestEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.pingerRequestEventListenersOnce = listenersOnce
	_e.pingerRequestEventMu.Unlock()
}

// EmitWatcherRequestEvent emits the specified event.
func (_e *ActorEventEmitter) EmitWatcherRequestEvent(message *battleAPI.WatcherRequest) {
	_e.watcherRequestEventMu.Lock()
	listeners := make([]func(message *battleAPI.WatcherRequest), len(_e.watcherRequestEventListeners))
	copy(listeners, _e.watcherRequestEventListeners)
	listenersOnce := _e.watcherRequestEventListenersOnce
	_e.watcherRequestEventListenersOnce = make([]func(message *battleAPI.WatcherRequest), 0)
	_e.watcherRequestEventMu.Unlock()
	for _, l := range listeners {
		l(message)
	}
	for _, l := range listenersOnce {
		l(message)
	}
}

// AddWatcherRequestEventListener registers the specified event listener.
func (_e *ActorEventEmitter) AddWatcherRequestEventListener(listener func(message *battleAPI.WatcherRequest)) {
	_e.watcherRequestEventMu.Lock()
	_e.watcherRequestEventListeners = append(_e.watcherRequestEventListeners, listener)
	_e.watcherRequestEventMu.Unlock()
}

// AddWatcherRequestEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *ActorEventEmitter) AddWatcherRequestEventListenerOnce(listener func(message *battleAPI.WatcherRequest)) {
	_e.watcherRequestEventMu.Lock()
	_e.watcherRequestEventListenersOnce = append(_e.watcherRequestEventListenersOnce, listener)
	_e.watcherRequestEventMu.Unlock()
}

// RemoveWatcherRequestEventListener removes the event listener previously registered.
func (_e *ActorEventEmitter) RemoveWatcherRequestEventListener(listener func(message *battleAPI.WatcherRequest)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.watcherRequestEventMu.Lock()
	listeners := _e.watcherRequestEventListeners[:0]
	for _, l := range _e.watcherRequestEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.watcherRequestEventListeners = listeners
	listenersOnce := _e.watcherRequestEventListenersOnce[:0]
	for _, l := range _e.watcherRequestEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.watcherRequestEventListenersOnce = listenersOnce
	_e.watcherRequestEventMu.Unlock()
}

// EmitUserLeaveEvent emits the specified event.
func (_e *ActorEventEmitter) EmitUserLeaveEvent() {
	_e.userLeaveEventMu.Lock()
	listeners := make([]func(), len(_e.userLeaveEventListeners))
	copy(listeners, _e.userLeaveEventListeners)
	listenersOnce := _e.userLeaveEventListenersOnce
	_e.userLeaveEventListenersOnce = make([]func(), 0)
	_e.userLeaveEventMu.Unlock()
	for _, l := range listeners {
		l()
	}
	for _, l := range listenersOnce {
		l()
	}
}

// AddUserLeaveEventListener registers the specified event listener.
func (_e *ActorEventEmitter) AddUserLeaveEventListener(listener func()) {
	_e.userLeaveEventMu.Lock()
	_e.userLeaveEventListeners = append(_e.userLeaveEventListeners, listener)
	_e.userLeaveEventMu.Unlock()
}

// AddUserLeaveEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *ActorEventEmitter) AddUserLeaveEventListenerOnce(listener func()) {
	_e.userLeaveEventMu.Lock()
	_e.userLeaveEventListenersOnce = append(_e.userLeaveEventListenersOnce, listener)
	_e.userLeaveEventMu.Unlock()
}

// RemoveUserLeaveEventListener removes the event listener previously registered.
func (_e *ActorEventEmitter) RemoveUserLeaveEventListener(listener func()) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.userLeaveEventMu.Lock()
	listeners := _e.userLeaveEventListeners[:0]
	for _, l := range _e.userLeaveEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.userLeaveEventListeners = listeners
	listenersOnce := _e.userLeaveEventListenersOnce[:0]
	for _, l := range _e.userLeaveEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.userLeaveEventListenersOnce = listenersOnce
	_e.userLeaveEventMu.Unlock()
}