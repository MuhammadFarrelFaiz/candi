package candishared

import (
	"bytes"
	"context"
)

// EventContext worker context in handler
type EventContext struct {
	ctx                      context.Context
	workerType, handlerRoute string
	header                   map[string]string
	key                      string
	err                      error

	messageBuff *bytes.Buffer
}

// NewEventContext event context constructor
func NewEventContext(buff *bytes.Buffer) *EventContext {
	buff.Reset()
	return &EventContext{
		messageBuff: buff,
	}
}

// SetContext setter
func (e *EventContext) SetContext(ctx context.Context) {
	e.ctx = ctx
}

// SetWorkerType setter
func (e *EventContext) SetWorkerType(w string) {
	e.workerType = w
}

// SetHandlerRoute setter can contains unique topic name, key, or task name
func (e *EventContext) SetHandlerRoute(h string) {
	e.handlerRoute = h
}

// SetHeader setter
func (e *EventContext) SetHeader(header map[string]string) {
	e.header = header
}

// SetKey setter
func (e *EventContext) SetKey(key string) {
	e.key = key
}

// SetError setter
func (e *EventContext) SetError(err error) {
	e.err = err
}

// Context get context
func (e *EventContext) Context() context.Context {
	return e.ctx
}

// WorkerType get worker type
func (e *EventContext) WorkerType() string {
	return e.workerType
}

// Header get context
func (e *EventContext) Header() map[string]string {
	return e.header
}

// HandlerRoute get handler name, contains unique topic name, or task name
func (e *EventContext) HandlerRoute() string {
	return e.handlerRoute
}

// Key get key
func (e *EventContext) Key() string {
	return e.key
}

// Message get context
func (e *EventContext) Message() []byte {
	return e.messageBuff.Bytes()
}

// Err get error
func (e *EventContext) Err() error {
	return e.err
}

// Read implement io.Reader
func (e *EventContext) Read(p []byte) (n int, err error) {
	return e.messageBuff.Read(p)
}

// Write implement io.Writer
func (e *EventContext) Write(p []byte) (n int, err error) {
	return e.messageBuff.Write(p)
}
