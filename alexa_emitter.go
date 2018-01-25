package alexa

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"

	"github.com/fsm/emitable"
)

// AlexaEmitter is an implementation of an FSM emitter for Amazon Alexa
//
// Because Amazon Alexa expects all outgoing messages / data to be in the form
// of a response to the inbound request (as compared to pushing messages), there
// is a speechBuffer that is generated within this struct as Emit is called
// throughout the lifecycle of a state.
//
// When Flush() is called on this struct, the SpeechBuffer is converted into the
// expected Alexa response, and written to the ResponseWriter.
//
// https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html#ssml-supported
type AlexaEmitter struct {
	ResponseWriter io.Writer
	hasSpeech      bool
	speechBuffer   bytes.Buffer
}

// Emit prepares the data to be output at the end of the request.
func (a *AlexaEmitter) Emit(input interface{}) error {
	switch v := input.(type) {

	case string:
		a.speechBuffer.WriteString("<s>")
		a.speechBuffer.WriteString(v)
		a.speechBuffer.WriteString("</s>")
		a.hasSpeech = true
		return nil

	case emitable.Sleep:
		a.speechBuffer.WriteString("<break time=\"")
		a.speechBuffer.WriteString(strconv.Itoa(v.LengthMillis))
		a.speechBuffer.WriteString("ms\"/>")
		return nil

	case emitable.QuickReply:
		// Write message
		a.speechBuffer.WriteString("<p>")
		a.speechBuffer.WriteString(v.Message)
		a.speechBuffer.WriteString("</p>")

		// Options
		optionsBuffer := new(bytes.Buffer)
		for i, reply := range v.Replies {
			optionsBuffer.WriteString(reply)
			if i+2 < len(v.Replies) && len(v.Replies) > 2 {
				optionsBuffer.WriteString(", ")
			} else if i+1 < len(v.Replies) {
				if len(v.Replies) > 2 {
					optionsBuffer.WriteString(", or ")
				} else {
					optionsBuffer.WriteString(" or ")
				}
			}
		}

		// Determine format
		format := "You can %v"
		if v.RepliesFormat != "" {
			format = v.RepliesFormat
		}

		// Write out options
		a.speechBuffer.WriteString("<p>")
		a.speechBuffer.WriteString(fmt.Sprintf(format, optionsBuffer.String()))
		a.speechBuffer.WriteString("</p>")
		return nil

	case emitable.Typing:
		// Intentionally do nothing
		return nil

	case emitable.Audio:
		// TODO
		return nil

	case emitable.Video:
		// TODO
		return nil

	case emitable.File:
		// TODO
		return nil

	case emitable.Image:
		// TODO
		return nil
	}
	return errors.New("AlexaEmitter cannot handle " + reflect.TypeOf(input).String())
}

// Flush writes the expected Alexa response to the a.ResponseWriter.
func (a *AlexaEmitter) Flush() error {
	// Prepare response body
	response := &ResponseBody{
		Version: "1.0",
		Response: &Response{
			ShouldEndSession: true,
		},
	}

	// Handle speech
	if a.hasSpeech {
		ssml := "<speak>" + a.speechBuffer.String() + "</speak>"
		response.Response.OutputSpeech = &OutputSpeech{
			Type: "SSML",
			SSML: ssml,
		}
	}

	// Output response
	b, err := json.Marshal(response)
	if err != nil {
		return err
	}
	fmt.Fprint(a.ResponseWriter, string(b))
	return nil
}
