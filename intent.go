package alexa

// Intent Object
// https://developer.amazon.com/docs/custom-skills/request-types-reference.html#intent-object
type Intent struct {
	Name               string          `json:"name,omitempty"`
	ConfirmationStatus string          `json:"confirmationStatus,omitempty"`
	Slots              map[string]Slot `json:"slots,omitempty"`
}

// Slot Object
// https://developer.amazon.com/docs/custom-skills/request-types-reference.html#slot-object
type Slot struct {
	Name               string      `json:"name,omitempty"`
	Value              string      `json:"value,omitempty"`
	ConfirmationStatus string      `json:"confirmationStatus,omitempty"`
	Resolutions        Resolutions `json:"resolutions,omitempty"`
}

// Resolutions Object
// https://developer.amazon.com/docs/custom-skills/request-types-reference.html#resolutions-object
type Resolutions struct {
	ResolutionsPerAuthority []Authority `json:"resolutionsPerAuthority,omitempty"`
}

// Authority Object
// See the Resolutions Object
type Authority struct {
	Authority string      `json:"authority,omitempty"`
	Status    Status      `json:"status,omitempty"`
	Values    []SlotValue `json:"values,omitempty"`
}

// Status Object
// See Resolutions
type Status struct {
	Code string `json:"code,omitempty"`
}

// SlotValue Object
// See Resolutions
type SlotValue struct {
	Value ResolvedValue `json:"value,omitempty"`
}

// ResolvedValue Object
// See Resolutions
type ResolvedValue struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
