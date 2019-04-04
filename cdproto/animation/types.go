package animation

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"bitbucket.org/ShipwrightTibi/chromecrawlingnew/cdproto/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Animation animation instance.
type Animation struct {
	ID           string  `json:"id"`               // Animation's id.
	Name         string  `json:"name"`             // Animation's name.
	PausedState  bool    `json:"pausedState"`      // Animation's internal paused state.
	PlayState    string  `json:"playState"`        // Animation's play state.
	PlaybackRate float64 `json:"playbackRate"`     // Animation's playback rate.
	StartTime    float64 `json:"startTime"`        // Animation's start time.
	CurrentTime  float64 `json:"currentTime"`      // Animation's current time.
	Type         Type    `json:"type"`             // Animation type of Animation.
	Source       *Effect `json:"source,omitempty"` // Animation's source animation node.
	CSSID        string  `json:"cssId,omitempty"`  // A unique ID for Animation representing the sources that triggered this CSS animation/transition.
}

// Effect animationEffect instance.
type Effect struct {
	Delay          float64           `json:"delay"`                   // AnimationEffect's delay.
	EndDelay       float64           `json:"endDelay"`                // AnimationEffect's end delay.
	IterationStart float64           `json:"iterationStart"`          // AnimationEffect's iteration start.
	Iterations     float64           `json:"iterations"`              // AnimationEffect's iterations.
	Duration       float64           `json:"duration"`                // AnimationEffect's iteration duration.
	Direction      string            `json:"direction"`               // AnimationEffect's playback direction.
	Fill           string            `json:"fill"`                    // AnimationEffect's fill mode.
	BackendNodeID  cdp.BackendNodeID `json:"backendNodeId,omitempty"` // AnimationEffect's target node.
	KeyframesRule  *KeyframesRule    `json:"keyframesRule,omitempty"` // AnimationEffect's keyframes.
	Easing         string            `json:"easing"`                  // AnimationEffect's timing function.
}

// KeyframesRule keyframes Rule.
type KeyframesRule struct {
	Name      string           `json:"name,omitempty"` // CSS keyframed animation's name.
	Keyframes []*KeyframeStyle `json:"keyframes"`      // List of animation keyframes.
}

// KeyframeStyle keyframe Style.
type KeyframeStyle struct {
	Offset string `json:"offset"` // Keyframe's time offset.
	Easing string `json:"easing"` // AnimationEffect's timing function.
}

// Type animation type of Animation.
type Type string

// String returns the Type as string value.
func (t Type) String() string {
	return string(t)
}

// Type values.
const (
	TypeCSSTransition Type = "CSSTransition"
	TypeCSSAnimation  Type = "CSSAnimation"
	TypeWebAnimation  Type = "WebAnimation"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Type) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Type) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Type) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch Type(in.String()) {
	case TypeCSSTransition:
		*t = TypeCSSTransition
	case TypeCSSAnimation:
		*t = TypeCSSAnimation
	case TypeWebAnimation:
		*t = TypeWebAnimation

	default:
		in.AddError(errors.New("unknown Type value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Type) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
