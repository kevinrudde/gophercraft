package event

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

type HandlerFunc[T any] func(event T)

type Handler[T any] struct {
	handlerFunc HandlerFunc[T]
	priority    Priority
}

type EventHandler struct {
	mutex    sync.RWMutex
	handlers map[reflect.Type][]interface{}
}

func New() *EventHandler {
	return &EventHandler{
		handlers: make(map[reflect.Type][]interface{}),
	}

}

var globalEvents = New()

func AddListener[T any](priority Priority, handlerFunc func(event T)) {
	globalEvents.mutex.Lock()
	defer globalEvents.mutex.Unlock()

	handler := Handler[T]{
		handlerFunc: handlerFunc,
		priority:    priority,
	}

	var eventGeneric T
	eventType := reflect.TypeOf(eventGeneric)

	globalEvents.handlers[eventType] = append(globalEvents.handlers[eventType], handler)

	sort.SliceStable(globalEvents.handlers[eventType], func(i, j int) bool {
		iHandler := globalEvents.handlers[eventType][i].(Handler[T])
		jHandler := globalEvents.handlers[eventType][j].(Handler[T])

		return iHandler.priority < jHandler.priority
	})
}

func Call[T any](event T, successCallback func()) error {
	eventType := reflect.TypeOf(event)

	globalEvents.mutex.RLock()
	defer globalEvents.mutex.RUnlock()

	if _, ok := globalEvents.handlers[eventType]; !ok {
		return fmt.Errorf("handler %s not found", eventType.String())
	}

	for _, h := range globalEvents.handlers[eventType] {
		handlerFunc := h.(Handler[T]).handlerFunc
		handlerFunc(event)
	}

	successCallback()
	return nil
}
