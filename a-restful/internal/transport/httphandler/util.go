package httphandler

func ResponseHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}
