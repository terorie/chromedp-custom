package overlay

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"bitbucket.org/ShipwrightTibi/chromecrawlingnew/cdproto/cdp"
	"bitbucket.org/ShipwrightTibi/chromecrawlingnew/cdproto/page"
)

// EventInspectNodeRequested fired when the node should be inspected. This
// happens after call to setInspectMode or when user manually inspects an
// element.
type EventInspectNodeRequested struct {
	BackendNodeID cdp.BackendNodeID `json:"backendNodeId"` // Id of the node to inspect.
}

// EventNodeHighlightRequested fired when the node should be highlighted.
// This happens after call to setInspectMode.
type EventNodeHighlightRequested struct {
	NodeID cdp.NodeID `json:"nodeId"`
}

// EventScreenshotRequested fired when user asks to capture screenshot of
// some area on the page.
type EventScreenshotRequested struct {
	Viewport *page.Viewport `json:"viewport"` // Viewport to capture, in device independent pixels (dip).
}

// EventInspectModeCanceled fired when user cancels the inspect mode.
type EventInspectModeCanceled struct{}
