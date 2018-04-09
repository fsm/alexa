package alexa

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/fsm/emitable"
)

// emitter is an implementation of an FSM emitter for Amazon Alexa
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
type emitter struct {
	ResponseWriter      io.Writer
	supportedInterfaces supportedInterfaces
	hasSpeech           bool
	speechBuffer        bytes.Buffer
	shouldEndSession    bool
	directives          []directive
}

// Emit prepares the data to be output at the end of the request.
func (e *emitter) Emit(input interface{}) error {
	switch v := input.(type) {

	case string:
		e.speechBuffer.WriteString(copyToSSML(v))
		e.hasSpeech = true
		return nil

	case emitable.Sleep:
		e.speechBuffer.WriteString("<break time=\"")
		e.speechBuffer.WriteString(strconv.Itoa(v.LengthMillis))
		e.speechBuffer.WriteString("ms\"/>")
		return nil

	case emitable.QuickReply:
		e.hasSpeech = true

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
		format := "You can say %v"
		if v.RepliesFormat != "" {
			format = v.RepliesFormat
		}

		// Write out options
		e.speechBuffer.WriteString(copyToSSML(fmt.Sprintf(format, optionsBuffer.String())))

		// Write message
		e.speechBuffer.WriteString(copyToSSML(v.Message))
		return nil

	case emitable.Typing:
		// Intentionally do nothing
		return nil

	case emitable.Audio:
		e.speechBuffer.WriteString("<audio src=\"")
		e.speechBuffer.WriteString(v.URL)
		e.speechBuffer.WriteString("\"/>")
		return nil

	case emitable.Video:
		// TODO
		return nil

	case emitable.File:
		// TODO
		return nil

	case emitable.Image:
		if e.supportedInterfaces.Display != nil {
			bodyTemplate := bodyTemplate7{
				BackgroundImageURL: v.URL,
			}
			e.directives = append(e.directives, bodyTemplate.asDirective())
		}
		return nil

	case EndSession:
		e.shouldEndSession = true
		return nil
	}
	return errors.New("AlexaEmitter cannot handle " + reflect.TypeOf(input).String())
}

// Converts copy to an appropriate SSML paragraph that will be read out
// by Alexa as naturally as possible.
//
// This function adds appropriate pauses to punctuation in the middle
// of a paragraph tag, as Alexa doesn't seem to do this normally for
// some strange reason.
func copyToSSML(copy string) string {
	// Trim last punctuation, as we don't want to add a pause,
	// as the pause will be handled by the end of the <p></p> SSML tag.
	trimmed := ""
	ssml := copy
	if strings.HasSuffix(copy, ".") {
		trimmed = "."
		ssml = strings.TrimSuffix(ssml, trimmed)
	} else if strings.HasSuffix(copy, "!") {
		trimmed = "!"
		ssml = strings.TrimSuffix(ssml, trimmed)
	} else if strings.HasSuffix(copy, "?") {
		trimmed = "?"
		ssml = strings.TrimSuffix(ssml, trimmed)
	}

	// Add pauses after punctuation
	ssml = strings.Replace(ssml, ".", ".<break time=\"150ms\"/>", -1)
	ssml = strings.Replace(ssml, "?", "?<break time=\"150ms\"/>", -1)
	ssml = strings.Replace(ssml, "!", "!<break time=\"150ms\"/>", -1)
	ssml = strings.Replace(ssml, ",", ",<break time=\"50ms\"/>", -1)

	// Join it all together and return
	return "<s>" + ssml + trimmed + "</s>"
}

// Flush writes the expected Alexa response to the a.ResponseWriter.
func (e *emitter) Flush() error {
	// Prepare response body
	response := &responseBody{
		Version: "1.0",
		Response: &response{
			ShouldEndSession: e.shouldEndSession,
			Directives:       &e.directives,
		},
	}

	// Handle speech
	if e.hasSpeech {
		ssml := "<speak>" + e.speechBuffer.String() + "</speak>"
		response.Response.OutputSpeech = &outputSpeech{
			Type: "SSML",
			SSML: ssml,
		}
	}

	// Output response
	b, err := json.Marshal(response)
	if err != nil {
		return err
	}
	fmt.Fprint(e.ResponseWriter, string(b))
	return nil
}
