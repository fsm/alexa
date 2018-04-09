package alexa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fsm/target-util"

	"github.com/fsm/fsm"
	skillserver "github.com/mikeflynn/go-alexa/skillserver"
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
		// Validate Request
		if !skillserver.IsValidAlexaRequest(w, r) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

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

		// Validate Timestamp
		if !validTimestamp(cb.Request.Timestamp) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Validate App ID
		if cb.Session.Application.ApplicationID != os.Getenv("ALEXA_APP_ID") &&
			cb.Context.System.Application.ApplicationID != os.Getenv("ALEXA_APP_ID") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Prepare the emitter
		emitter := &emitter{
			ResponseWriter:      w,
			supportedInterfaces: cb.Context.System.Device.SupportedInterfaces,
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

func validTimestamp(timestamp string) bool {
	reqTimestamp, _ := time.Parse("2006-01-02T15:04:05Z", timestamp)
	if time.Since(reqTimestamp) < time.Duration(150)*time.Second {
		return true
	}
	return false
}
