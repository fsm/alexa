package alexa

const AudioDirectiveType = "AudioPlayer.Play"

type audioDirective struct {
	Type         string    `json:"type"`
	PlayBehavior string    `json:"playBehavior"`
	AudioItem    AudioItem `json:"audioItem"`
}

type AudioItem struct {
	Stream Stream `json:"stream"`
}

type Stream struct {
	URL                  string `json:"url"`
	Token                string `json:"token"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
}

// Make this a directive
func (a audioDirective) DirectiveType() string {
	return a.Type
}
