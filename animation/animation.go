package Animation

import (
	"github.com/mkenney/go-chrome/runtime"
)

/*
GetCurrentTimeParams represents Animation.getCurrentTime parameters.
*/
type GetCurrentTimeParams struct {
	// ID of animation.
	ID string `json:"id"`
}

/*
GetCurrentTimeResult represents the result of calls to Animation.getCurrentTime.
*/
type GetCurrentTimeResult struct {
	// ID of animation.
	ID string `json:"id"`
}

/*
GetPlaybackRateResult represents the result of calls to Animation.getPlaybackRate.
*/
type GetPlaybackRateResult struct {
	// Playback rate for animations on page.
	PlaybackRate float64 `json:"playbackRate"`
}

/*
ReleaseAnimationsParams represents Animation.releaseAnimations parameters.
*/
type ReleaseAnimationsParams struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
}

/*
ResolveAnimationParams represents Animation.resolveAnimation parameters.
*/
type ResolveAnimationParams struct {
	// Animation ID.
	AnimationID string `json:"animationId"`
}

/*
ResolveAnimationResult represents the result of calls to Animation.resolveAnimation.
*/
type ResolveAnimationResult struct {
	// Corresponding remote object.
	RemoteObject Runtime.RemoteObject `json:"remoteObject"`
}

/*
SeekAnimationsParams represents Animation.seekAnimations parameters.
*/
type SeekAnimationsParams struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`

	// Set the current time of each animation.
	CurrentTime float64 `json:"currentTime"`
}

/*
SetPausedParams represents Animation.setPaused parameters.int16
*/
type SetPausedParams struct {
	// Animations to set the pause state of.
	Animations []string `json:"animations"`

	// Paused state to set to.
	Paused bool `json:"paused"`
}

/*
SetPlaybackRateParams represents Animation.setPlaybackRate parameters.
*/
type SetPlaybackRateParams struct {
	// Playback rate for animations on page.
	PlaybackRate float64 `json:"playbackRate"`
}

/*
SetTimingParams represents Animation.setTiming parameters.
*/
type SetTimingParams struct {
	// Animation ID.
	AnimationID string `json:"animationId"`

	// Duration of the animation.
	Duration float64 `json:"duration"`

	// Delay of the animation.
	Delay float64 `json:"delay"`
}

/*
CanceledEvent represents Animation.animationCanceled event data.
*/
type CanceledEvent struct {
	// ID of the animation that was cancelled.
	ID string `json:"id"`
}

/*
CreatedEvent represents Animation.animationCreated event data.
*/
type CreatedEvent struct {
	// ID of the animation that was created.
	ID string `json:"id"`
}

/*
StartedEvent represents Animation.animationStarted event data.
*/
type StartedEvent struct {
	// Animation that was started.
	Animation Animation `json:"animation"`
}

/*
Animation instance.
*/
type Animation struct {
	// Animation's id.
	ID string `json:"id"`

	// Animation's name.
	Name string `json:"name"`

	// Animation's internal paused state.
	PausedState bool `json:"pausedState"`

	// Animation's play state.
	PlayState string `json:"playState"`

	// Animation's playback rate.
	PlaybackRate float64 `json:"playbackRate"`

	// Animation's start time.
	StartTime float64 `json:"startTime"`

	// Animation's current time.
	CurrentTime float64 `json:"currentTime"`

	// Animation type of Animation. Allowed values: CSSTransition, CSSAnimation, WebAnimation.
	Type string `json:"type"`

	// Optional. Animation's source animation node.
	//
	// This expects an instance of AnimationEffect, but that doesn't omitempty correctly so it must
	// be added manually.
	Source interface{} `json:"source,omitempty"`

	// Optional. A unique ID for Animation representing the sources that triggered this CSS
	// animation/transition.
	CSSID string `json:"cssId,omitempty"`
}

/*
Effect instance
*/
type Effect struct {
	// AnimationEffect's delay.
	Delay float64 `json:"delay"`

	// AnimationEffect's end delay.
	EndDelay float64 `json:"endDelay"`

	// AnimationEffect's iteration start.
	IterationStart float64 `json:"iterationStart"`

	// AnimationEffect's iterations.
	Iterations float64 `json:"iterations"`

	// AnimationEffect's iteration duration.
	Duration float64 `json:"duration"`

	// AnimationEffect's playback direction.
	Direction string `json:"direction"`

	// AnimationEffect's fill mode.
	Fill string `json:"fill"`

	// Optional. AnimationEffect's target node.
	//
	// This expects an instance of DOM.BackendNodeID, but that doesn't omitempty correctly so it
	// must be added manually.
	BackendNodeID interface{} `json:"backendNodeId,omitempty"`

	// Optional. AnimationEffect's keyframes.
	//
	// This expects an instance of KeyframesRule, but that doesn't omitempty correctly so it must be
	// added manually.
	KeyframesRule interface{} `json:"keyframesRule,omitempty"`

	// AnimationEffect's timing function.
	Easing string `json:"easing"`
}

/*
KeyframesRule is the keyframes rule
*/
type KeyframesRule struct {
	// Optional. CSS keyframed animation's name.
	Name string `json:"name,omitempty"`

	// List of animation keyframes.
	Keyframes []*KeyframeStyle `json:"keyframes"`
}

/*
KeyframeStyle is the keyframe Style
*/
type KeyframeStyle struct {
	// Keyframe's time offset.
	Offset string `json:"offset"`

	// AnimationEffect's timing function.
	Easing string `json:"easing"`
}