package alexa

// responseBody Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#response-format
type responseBody struct {
	Version           string                 `json:"version,omitempty"`
	SessionAttributes map[string]interface{} `json:"sessionAttrributes,omitempty"`
	Response          *response              `json:"response,omitempty"`
}

// response Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#response-parameters
type response struct {
	OutputSpeech     *outputSpeech `json:"outputSpeech"`
	Card             *card         `json:"card"`
	Reprompt         *reprompt     `json:"reprompt"`
	ShouldEndSession bool          `json:"shouldEndSession"`
	Directives       *[]directive  `json:"directives"`
}

// outputSpeech Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#outputspeech-object
type outputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
	SSML string `json:"ssml"`
}

// Card Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#card-object
type card struct {
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Text    string    `json:"text"`
	Image   cardImage `json:"image"`
}

// cardImage is an object within a Card
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#card-object
type cardImage struct {
	SmallImageURL string `json:"smallImageUrl"`
	LargeImageURL string `json:"largeImageUrl"`
}

// reprompt Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#reprompt-object
type reprompt struct {
	OutputSpeech outputSpeech `json:"outputSpeech"`
}

// directive is an object nested within the Response Object
// There are many possible directives, follow the links included in the description:
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#response-object
type directive struct {
	Type string `json:"type"`
	// TODO, implement specific directives
}
