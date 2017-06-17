package alexa

type Alexa struct {
	ApplicationID   string
	RequestHandler  RequestHandler
	IgnoreTimestamp bool
}

func (a *Alexa) ProcessRequest(body *RequestBody) (*ResponseBody, error) {
	response := &ResponseBody{
		Version: body.Version,
		Response: &Response{
			ShouldSessionEnd: true,
		},
	}
	if body.Session.New {
		if err := a.RequestHandler.OnSessionStarted(body.Request, body.Session, response.Response); err != nil {
			return nil, err
		}
	}

	switch body.Request.Type {
	case LaunchRequest:
		if err := a.RequestHandler.OnLaunch(body.Request, body.Session, response.Response); err != nil {
			return nil, err
		}
	case IntentRequest:
		if err := a.RequestHandler.OnIntent(body.Request, body.Session, response.Response); err != nil {
			return nil, err
		}
	case SessionEndedRequest:
		if err := a.RequestHandler.OnSessionEnded(body.Request, body.Session, response.Response); err != nil {
			return nil, err
		}
	}
	return response, nil
}
