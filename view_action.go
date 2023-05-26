package gotfy

import (
	"encoding/json"
	"net/url"
)

type ViewAction struct {
	Label string
	Link  *url.URL
	Clear bool
}

func (v *ViewAction) actionType() ActionButtonType {
	return View
}

func (v *ViewAction) MarshalJSON() ([]byte, error) {
	buf := []byte(`{"action":"view","label":`)

	labelBuf, err := json.Marshal(v.Label)
	if err != nil {
		return nil, err
	}
	buf = append(buf, labelBuf...)

	if v.Link != nil {
		urlBuf, err := json.Marshal(v.Link.String())
		if err != nil {
			return nil, err
		}
		buf = append(buf, `,"url":`...)
		buf = append(buf, urlBuf...)
	}

	if v.Clear {
		buf = append(buf, `,"clear":true`...)
	}

	return append(buf, '}'), nil
}
