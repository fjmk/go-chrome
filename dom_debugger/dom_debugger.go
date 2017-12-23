package DOMDebugger

import (
	"fmt"

	"github.com/mkenney/go-chrome/dom"
	"github.com/mkenney/go-chrome/runtime"
)

/*
GetEventListenersParams represents DOMDebugger.getEventListeners parameters.
*/
type GetEventListenersParams struct {
	// ID of the object to return listeners for.
	ObjectID Runtime.RemoteObjectID `json:"objectId"`

	// Optional. The maximum depth at which Node children should be retrieved, defaults to 1. Use -1
	// for the entire subtree or provide an integer larger than 0.
	Depth int `json:"depth,omitempty"`

	// Optional. Whether or not iframes and shadow roots should be traversed when returning the
	// subtree (default is false). Reports listeners for all contexts if pierce is enabled.
	Pierce bool `json:"pierce,omitempty"`
}

/*
GetEventListenersResult represents the result of calls to DOMDebugger.getEventListeners.
*/
type GetEventListenersResult struct {
	// Array of relevant listeners.
	Listeners []EventListener `json:"listeners"`
}

/*
RemoveDOMBreakpointParams represents DOMDebugger.removeDOMBreakpoint parameters.
*/
type RemoveDOMBreakpointParams struct {
	// ID of the node to remove breakpoint from.
	NodeID DOM.NodeID `json:"nodeId"`

	// Type of the breakpoint to remove.
	Type DOMBreakpointType `json:"type"`
}

/*
RemoveEventListenerBreakpointParams represents DOMDebugger.removeEventListenerBreakpoint parameters.
*/
type RemoveEventListenerBreakpointParams struct {
	// Event name.
	EventName string `json:"eventName"`

	// Optional. EventTarget interface name. EXPERIMENTAL
	TargetName string `json:"targetName,omitempty"`
}

/*
RemoveInstrumentationBreakpointParams represents DOMDebugger.removeInstrumentationBreakpoint parameters.
*/
type RemoveInstrumentationBreakpointParams struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

/*
RemoveXHRBreakpointParams represents DOMDebugger.removeXHRBreakpoint parameters.
*/
type RemoveXHRBreakpointParams struct {
	// Resource URL substring.
	URL string `json:"url"`
}

/*
SetDOMBreakpointParams represents DOMDebugger.setDOMBreakpoint parameters.
*/
type SetDOMBreakpointParams struct {
	// ID of the node to set breakpoint on.
	NodeID DOM.NodeID `json:"nodeId"`

	// Type of the operation to stop upon.
	Type DOMBreakpointType `json:"type"`
}

/*
SetEventListenerBreakpointParams represents DOMDebugger.setEventListenerBreakpoint parameters.
*/
type SetEventListenerBreakpointParams struct {
	// DOM Event name to stop on (any DOM event will do).
	EventName string `json:"eventName"`

	// Optional. EventTarget interface name to stop on. If equal to "*" or not provided, will stop
	// on any EventTarget. EXPERIMENTAL
	TargetName string `json:"targetName,omitempty"`
}

/*
SetInstrumentationBreakpointParams represents DOMDebugger.setInstrumentationBreakpoint parameters.
*/
type SetInstrumentationBreakpointParams struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

/*
SetXHRBreakpointParams represents DOMDebugger.setXHRBreakpoint parameters.
*/
type SetXHRBreakpointParams struct {
	// Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
	URL string `json:"url"`
}

/*
DOMBreakpointType is the DOM breakpoint type.
*/
type DOMBreakpointType int

const (
	_subtreeModified DOMBreakpointType = iota
	_attributeModified
	_nodeRemoved
)

func (a DOMBreakpointType) String() string {
	if a == 0 {
		return "subtree-modified"
	}
	if a == 1 {
		return "attribute-modified"
	}
	if a == 2 {
		return "node-removed"
	}
	panic(fmt.Errorf("Invalid DOMBreakpointType %d", a))
}

/*
EventListener is the object event listener.
*/
type EventListener struct {
	// EventListener's type.
	Type string `json:"type"`

	// EventListener's useCapture.
	UseCapture bool `json:"useCapture"`

	// EventListener's passive flag.
	Passive bool `json:"passive"`

	// EventListener's once flag.
	Once bool `json:"once"`

	// Script ID of the handler code.
	ScriptID Runtime.ScriptID `json:"scriptId"`

	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`

	// Column number in the script (0-based).
	ColumnNumber int `json:"columnNumber"`

	// Optional. Event handler function value.
	//
	// This expects an instance of Runtime.RemoteObject, but that doesn't omitempty correctly so it
	// must be added manually.
	Handler interface{} `json:"handler,omitempty"`

	// Optional. Event original handler function value.
	//
	// This expects an instance of Runtime.RemoteObject, but that doesn't omitempty correctly so it
	// must be added manually.
	OriginalHandler interface{} `json:"originalHandler,omitempty"`

	// Optional. Node the listener is added to (if any).
	BackendNodeID DOM.BackendNodeID `json:"backendNodeId,omitempty"`
}