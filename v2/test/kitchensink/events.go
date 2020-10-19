package main

import (
	wails "github.com/wailsapp/wails/v2"
)

// Events struct
type Events struct {
	runtime *wails.Runtime
}

// WailsInit is called at application startup
func (e *Events) WailsInit(runtime *wails.Runtime) error {
	// Perform your setup here
	e.runtime = runtime
	return nil
}

// On will subscribe to the given event name
func (e *Events) On(eventName string) {
	e.runtime.Events.On(eventName, func(args ...interface{}) {
		type callbackData struct {
			Name string
			Data []interface{}
		}
		result := callbackData{Name: eventName, Data: args}
		e.runtime.Events.Emit("event fired by go subscriber", result)
	})
}

// Once will subscribe to the given event name
func (e *Events) Once(eventName string) {
	e.runtime.Events.Once(eventName, func(args ...interface{}) {
		type callbackData struct {
			Name string
			Data []interface{}
		}
		result := callbackData{Name: eventName, Data: args}
		e.runtime.Events.Emit("once event fired by go subscriber", result)
	})
}

// Emit will emit
func (e *Events) Emit(eventName string, data []interface{}) {
	e.runtime.Events.Emit(eventName, data...)
}