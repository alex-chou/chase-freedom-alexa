package alexa

const (
	LaunchRequest       RequestType = "LaunchRequest"
	IntentRequest       RequestType = "IntentRequest"
	SessionEndedRequest RequestType = "SessionEndedRequest"

	ConfirmationNone      ConfirmationStatus = "NONE"
	ConfirmationConfirmed ConfirmationStatus = "CONFIRMED"
	ConfirmationDenied    ConfirmationStatus = "DENIED"
)

type RequestType string
type ConfirmationStatus string

// RequestBody is the structure of an Amazon Alexa request body.
type RequestBody struct {
	Version string   `json:"version"`
	Session *Session `json:"session"`
	Request *Request `json:"request"`
}

// Session
type Session struct {
	New         bool   `json:"new"`
	SessionID   string `json:"sessionId"`
	Application struct {
		ID string `json:"applicationId"`
	} `json:"application"`
	Attributes map[string]interface{} `json:"attributes"`
	User       struct {
		ID          string `json:"userId"`
		Permissions struct {
			ConsentToken string `json:"consentToken"`
		} `json:"permissions"`
	} `json:"user"`
}

// Request
type Request struct {
	Type        RequestType `json:"type"`
	RequestID   string      `json:"requestId"`
	Timestamp   string      `json:"timestamp"`
	DialogState string      `json:"dialogState"`
	Locale      string      `json:"locale"`
	Intent      *Intent     `json:"intent"`
	Error       struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}

// Intent
type Intent struct {
	Name         string             `json:"name"`
	Confirmation ConfirmationStatus `json:"confirmationStatus"`
	Slots        map[string]*Slot   `json:"slots"`
}

// Slot
type Slot struct {
	Name         string             `json:"name"`
	Value        string             `json:"value"`
	Confirmation ConfirmationStatus `json:"confirmationStatus"`
}
