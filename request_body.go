package alexa

const (
	RequestTypeLaunch       = "LaunchRequest"
	RequestTypeIntent       = "IntentRequest"
	RequestTypeSessionEnded = "SessionEndedRequest"
)

// requestBody Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#request-body-parameters
type requestBody struct {
	Version string  `json:"version,omitempty"`
	Session session `json:"session,omitempty"`
	Context context `json:"context,omitempty"`
	Request request `json:"request,omitempty"`
}

// session Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#session-object
type session struct {
	New         bool                   `json:"new,omitempty"`
	SessionID   string                 `json:"sessionId,omitempty"`
	Application application            `json:"application,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	User        user                   `json:"user,omitempty"`
}

type application struct {
	ApplicationID string `json:"applicationId,omitempty"`
}

type user struct {
	UserID      string      `json:"userId,omitempty"`
	AccessToken string      `json:"accessToken,omitempty"`
	Permissions permissions `json:"permissions,omitempty"`
}

type permissions struct {
	ConsentToken string `json:"consentToken,omitempty"`
}

// context Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#context-object
type context struct {
	System      system      `json:"System,omitempty"`
	AudioPlayer audioPlayer `json:"AudioPlayer,omitempty"`
}

// system Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#system-object
type system struct {
	Device         device      `json:"device,omitempty"`
	Application    application `json:"application,omitempty"`
	User           user        `json:"user,omitempty"`
	APIEndpoint    string      `json:"apiEndpoint,omitempty"`
	APIAccessToken string      `json:"apiAccessToken,omitempty"`
}

type device struct {
	DeviceID            string              `json:"deviceId,omitempty"`
	SupportedInterfaces supportedInterfaces `json:"SupportedInterfaces,omitempty"`
}

// audioPlayer Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#audioplayer-object
type audioPlayer struct {
	PlayerActivity       string `json:"playerActivity,omitempty"`
	Token                string `json:"token,omitempty"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
}

type supportedInterfaces struct {
	AudioPlayer interface{} `json:"AudioPlayer,omitempty"`
}

// A request object that provides the details of the user’s request. There are several different request types avilable, see:
// Standard Requests: https://developer.amazon.com/docs/custom-skills/request-types-reference.html
// AudioPlayer Requests: https://developer.amazon.com/docs/custom-skills/audioplayer-interface-reference.html#requests
// PlaybackController Requests: https://developer.amazon.com/docs/custom-skills/playback-controller-interface-reference.html#requests
type request struct {
	Type        string `json:"type,omitempty"`
	RequestID   string `json:"requestId,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	Reason      string `json:"reason,omitempty"`
	Error       Error  `json:"error,omitempty"`
	DialogState string `json:"dialogState,omitempty"`
	Locale      string `json:"locale,omitempty"`
	Intent      Intent `json:"intent,omitempty"`
}

type Error struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}
