package event

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

type HandlerFunc[K any] func(event K)

type Handler[K any] struct {
	handlerFunc HandlerFunc[K]
	priority    Priority
}

type EventHandler struct {
	mutex    sync.RWMutex
	handlers map[string][]any // TODO: change to Handler[K] ideally
}

func New() *EventHandler {
	return &EventHandler{
		handlers: map[string][]any{},
	}

}

var globalEvents = New()

func Listen[K any](priority Priority, handlerFunc func(event K)) {
	globalEvents.mutex.Lock()
	defer globalEvents.mutex.Unlock()

	handler := Handler[K]{
		handlerFunc: handlerFunc,
		priority:    priority,
	}

	var eventGeneric K
	eventType := reflect.TypeOf(eventGeneric).String()

	globalEvents.handlers[eventType] = append(globalEvents.handlers[eventType], handler)

	sort.SliceStable(globalEvents.handlers[eventType], func(i, j int) bool {
		iHandler := globalEvents.handlers[eventType][i].(Handler[K])
		jHandler := globalEvents.handlers[eventType][j].(Handler[K])

		return iHandler.priority < jHandler.priority
	})
}

func Call[K any](event K, successCallback func()) error {
	eventType := reflect.TypeOf(event).String()

	globalEvents.mutex.RLock()
	defer globalEvents.mutex.RUnlock()

	if _, ok := globalEvents.handlers[eventType]; !ok {
		return fmt.Errorf("handler %s not found", event)
	}

	for _, h := range globalEvents.handlers[eventType] {
		handlerFunc := h.(Handler[K]).handlerFunc
		handlerFunc(event)
	}

	successCallback()
	return nil
}
