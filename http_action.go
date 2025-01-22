package gotfy

import (
	"encoding/json"
)

// HttpAction allows attaching an HTTP request action to a notification
type HttpAction[X comparable] struct {
	Label   string            `json:"label"`   //  string  -  Open garage door  Label of the action button in the notification
	URL     string            `json:"url,omitempty"`     //  string  -  https://ntfy.sh/mytopic  URL to which the HTTP request will be sent
	Method  string            `json:"method"`  //  GET/POST/PUT/...  POST GET  HTTP method to use for request, default is POST ⚠️
	Headers map[string]string `json:"headers"` // map of strings  -  see above  HTTP headers to pass in request. When publishing as JSON, headers are passed as a map. When the simple format is used, use headers.<header1>=<value>.
	Body    X                 `json:"body"`    //  string  empty  some body, somebody?  HTTP body
	Clear   bool              `json:"clear"`   //  boolean  false  true  Clear notification after HTTP request succeeds. If the request fails, the notification is not cleared.
}

func (h *HttpAction[X]) actionType() ActionButtonType {
	return HTTP
}

func (h *HttpAction[X]) MarshalJSON() ([]byte, error) {
	m := map[string]any{
		"action": "http",
		"label":  h.Label,
		"url":    h.URL,
	}

	if meth := h.Method; meth != "" {
		m["method"] = meth
	}

	if headers := h.Headers; len(headers) > 0 {
		m["headers"] = headers
	}

	// double marshal body to work for HTTP requests.
	// the api wants body to be a marshaled JSON string
	// so double marshal to escape the string
	var zeroVal X
	if body := h.Body; body != zeroVal {
		buf, err := json.Marshal(h.Body)
		if err != nil {
			return nil, err
		}
		str, err := json.Marshal(string(buf))
		if err != nil {
			return nil, err
		}
		m["body"] = str
	}

	if h.Clear {
		m["clear"] = true
	}

	return json.Marshal(m)
}
