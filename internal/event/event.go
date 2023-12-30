package event

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"sync"
)

type Message interface{}

type HandlerFunc func(message Message)

type Handler struct {
	handlerFunc HandlerFunc
	priority    Priority
}

type Events struct {
	mutex    sync.RWMutex
	wg       sync.WaitGroup
	handlers map[string][]Handler
}

func New() *Events {
	return &Events{
		handlers: map[string][]Handler{},
	}
}

func (e *Events) Listen(message Message, priority Priority, handlerFunc func(message Message)) {
	event := reflect.TypeOf(message).String()

	e.mutex.Lock()
	defer e.mutex.Unlock()

	handler := Handler{
		handlerFunc: handlerFunc,
		priority:    priority,
	}

	e.handlers[event] = append(e.handlers[event], handler)

	sort.SliceStable(e.handlers[event], func(i, j int) bool {
		return e.handlers[event][i].priority > e.handlers[event][j].priority
	})
}

func (e *Events) Call(message Message, successCallback func()) error {
	event := reflect.TypeOf(message).String()

	e.mutex.RLock()
	defer e.mutex.RUnlock()

	if _, ok := e.handlers[event]; !ok {
		return fmt.Errorf("handler %s not found", event)
	}

	log.Println(len(e.handlers[event]))

	for _, h := range e.handlers[event] {
		e.wg.Add(1)
		handlerFunc := h.handlerFunc
		go func() {
			defer e.wg.Done()
			handlerFunc(message)
		}()
	}

	e.wg.Wait()

	successCallback()
	return nil
}

var globalEvents = New()

func Listen(message Message, priority Priority, handlerFunc func(message Message)) {
	globalEvents.Listen(message, priority, handlerFunc)
}

func Call(message Message, successCallback func()) error {
	return globalEvents.Call(message, successCallback)
}
