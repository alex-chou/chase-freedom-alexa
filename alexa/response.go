package alexa

const (
	PlainTextSpeech SpeechType = "PlainText"
	SSMLSpeech      SpeechType = "SSML"

	SimpleCard      CardType = "Simple"
	StandardCard    CardType = "Standard"
	LinkAccountCard CardType = "LinkAccount"
)

type SpeechType string
type CardType string

// ResponseBody
type ResponseBody struct {
	Version           string             `json:"version"`
	SessionAttributes map[string]*Object `json:"sessionAttributes,omitempty"`
	Response          *Object            `json:"response"`
}

// Object
type Object struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	Directives       []*Directive  `json:"directives,omitempty"`
	ShouldEndSession bool          `json:"shouldEndSession"`
}

// OutputSpeech
type OutputSpeech struct {
	Type SpeechType `json:"type"`
	Text string     `json:"text,omitempty"`
	SSML string     `json:"ssml,omitempty"`
}

// Card
type Card struct {
	Type    CardType `json:"type"`
	Title   string   `json:"title,omitempty"`
	Content string   `json:"content,omitempty"`
	Text    string   `json:"text,omitempty"`
	Image   struct {
		SmallURL string `json:"smallImageUrl"`
		LargeURL string `json:"largeImageUrl"`
	} `json:"image,omitempty"`
}
