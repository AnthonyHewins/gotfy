package gotfy

import (
	"encoding/json"
)

type ViewAction struct {
	Label string
	Link  string
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

	if v.Link != "" {
		urlBuf, err := json.Marshal(v.Link)
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
