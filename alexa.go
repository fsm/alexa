package alexa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fsm/target-util"

	"github.com/fsm/fsm"
)

// DistillIntent is a function that is responsible for converting
// an intent into an input string
type DistillIntent func(Intent) string

// GetWebhook returns the webhook that Alexa expects to communicate with
func GetWebhook(stateMachine fsm.StateMachine, store fsm.Store, distillIntent DistillIntent) func(http.ResponseWriter, *http.Request) {
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

		// Perform a Step
		input := distillIntent(cb.Request.Intent)
		stateMap := targetutil.GetStateMap(stateMachine)
		emitter := &emitter{
			ResponseWriter: w,
		}
		targetutil.Step(cb.Session.User.UserID, input, store, emitter, stateMap)

		// Write body
		err = emitter.Flush()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
