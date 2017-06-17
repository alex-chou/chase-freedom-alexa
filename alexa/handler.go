package alexa

type RequestHandler interface {
	OnSessionStarted(*Request, *Session, *Response) error
	OnLaunch(*Request, *Session, *Response) error
	OnIntent(*Request, *Session, *Response) error
	OnSessionEnded(*Request, *Session, *Response) error
}
