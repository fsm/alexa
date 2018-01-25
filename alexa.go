package alexa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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
		cb := &RequestBody{}
		err := json.Unmarshal([]byte(body), cb)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Get traverser
		newTraverser := false
		traverser, err := store.FetchTraverser(cb.Session.User.UserID)
		if err != nil {
			traverser, _ = store.CreateTraverser(cb.Session.User.UserID)
			traverser.SetCurrentState("start")
			newTraverser = true
		}

		// Create emitter
		emitter := &AlexaEmitter{
			ResponseWriter: w,
		}

		// Get current state
		currentState := stateMachine[traverser.CurrentState()](emitter, traverser)
		if newTraverser {
			currentState.EntryAction()
		}

		// Transition
		distilledValue := distillIntent(cb.Request.Intent)
		newState := currentState.Transition(distilledValue)
		err = newState.EntryAction()
		if err == nil {
			traverser.SetCurrentState(newState.Slug)
		}

		// Write body
		err = emitter.Flush()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
