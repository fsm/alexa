package alexa

// ResponseBody Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#response-format
type ResponseBody struct {
	Version           string                 `json:"version,omitempty"`
	SessionAttributes map[string]interface{} `json:"sessionAttrributes,omitempty"`
	Response          *Response              `json:"response,omitempty"`
}

// Response Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#response-parameters
type Response struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech"`
	Card             *Card         `json:"card"`
	Reprompt         *Reprompt     `json:"reprompt"`
	ShouldEndSession bool          `json:"shouldEndSession"`
	Directives       *[]Directive  `json:"directives"`
}

// OutputSpeech Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#outputspeech-object
type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
	SSML string `json:"ssml"`
}

// Card Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#card-object
type Card struct {
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Text    string    `json:"text"`
	Image   CardImage `json:"image"`
}

// CardImage is an object within a Card
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#card-object
type CardImage struct {
	SmallImageURL string `json:"smallImageUrl"`
	LargeImageURL string `json:"largeImageUrl"`
}

// Reprompt Object
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#reprompt-object
type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech"`
}

// Directive is an object nested within the Response Object
// There are many possible directives, follow the links included in the description:
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#response-object
type Directive struct {
	Type string `json:"type"`
	// TODO, implement specific directives
}
