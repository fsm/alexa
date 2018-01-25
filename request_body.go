package alexa

// RequestBody Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#request-body-parameters
type RequestBody struct {
	Version string  `json:"version,omitempty"`
	Session Session `json:"session,omitempty"`
	Context Context `json:"context,omitempty"`
	Request Request `json:"request,omitempty"`
}

// Session Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#session-object
type Session struct {
	New         bool                   `json:"new,omitempty"`
	SessionID   string                 `json:"sessionId,omitempty"`
	Application Application            `json:"application,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	User        User                   `json:"user,omitempty"`
}

// Application Object
// See Session
type Application struct {
	ApplicationID string `json:"applicationId,omitempty"`
}

// User Object
// See Session
type User struct {
	UserID      string      `json:"userId,omitempty"`
	AccessToken string      `json:"accessToken,omitempty"`
	Permissions Permissions `json:"permissions,omitempty"`
}

// Permissions Object
// See User
type Permissions struct {
	ConsentToken string `json:"consentToken,omitempty"`
}

// Context Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#context-object
type Context struct {
	System      System      `json:"System,omitempty"`
	AudioPlayer AudioPlayer `json:"AudioPlayer,omitempty"`
}

// System Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#system-object
type System struct {
	Device         Device      `json:"device,omitempty"`
	Application    Application `json:"application,omitempty"`
	User           User        `json:"user,omitempty"`
	APIEndpoint    string      `json:"apiEndpoint,omitempty"`
	APIAccessToken string      `json:"apiAccessToken,omitempty"`
}

// Device Object
// See System
type Device struct {
	DeviceID            string              `json:"deviceId,omitempty"`
	SupportedInterfaces SupportedInterfaces `json:"SupportedInterfaces,omitempty"`
}

// AudioPlayer Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#audioplayer-object
type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity,omitempty"`
	Token                string `json:"token,omitempty"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
}

// SupportedInterfaces Object
// See AudioPlayer
type SupportedInterfaces struct {
	AudioPlayer interface{} `json:"AudioPlayer,omitempty"`
}

// A Request object that provides the details of the user’s request. There are several different request types avilable, see:
// Standard Requests: https://developer.amazon.com/docs/custom-skills/request-types-reference.html
// AudioPlayer Requests: https://developer.amazon.com/docs/custom-skills/audioplayer-interface-reference.html#requests
// PlaybackController Requests: https://developer.amazon.com/docs/custom-skills/playback-controller-interface-reference.html#requests
type Request struct {
	Type        string `json:"type,omitempty"`
	RequestID   string `json:"requestId,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	Reason      string `json:"reason,omitempty"`
	Error       Error  `json:"error,omitempty"`
	DialogState string `json:"dialogState,omitempty"`
	Locale      string `json:"locale,omitempty"`
	Intent      Intent `json:"intent,omitempty"`
}

// Error Object
// See Request
type Error struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}
