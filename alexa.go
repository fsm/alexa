package alexa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fsm/target-util"

	"github.com/fsm/fsm"
)

const platform = "amazon-alexa"

// DistillIntent is a function that is responsible for converting
// an intent into an input string
type DistillIntent func(Intent) string

// GetWebhook returns the webhook that Alexa expects to communicate with
func GetWebhook(stateMachine fsm.StateMachine, store fsm.Store, distillIntent DistillIntent) func(http.ResponseWriter, *http.Request) {
	// Get StateMap
	stateMap := targetutil.GetStateMap(stateMachine)

	// Return Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Get body
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()

		// Parse body into struct
		cb := &requestBody{}
		err := json.Unmarshal([]byte(body), cb)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Prepare the emitter
		emitter := &emitter{
			ResponseWriter: w,
		}

		// Handle request type
		var input string
		switch cb.Request.Type {
		case RequestTypeLaunch:
			input = RequestTypeLaunch
			break
		case RequestTypeIntent:
			input = distillIntent(cb.Request.Intent)
			break
		case RequestTypeSessionEnded:
			emitter.Flush()
			return
		}

		// Perform a Step
		targetutil.Step(platform, cb.Session.User.UserID, input, store, emitter, stateMap)

		// Write body
		err = emitter.Flush()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
