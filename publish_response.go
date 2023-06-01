package gotfy

import "time"

type PublishResp struct {
	ID      string    `json:"id"`      // :"bUhbhgmmbeW0"
	Time    time.Time `json:"time"`    // :1685150791
	Expires time.Time `json:"expires"` // :1685193991
	Event   string    `json:"event"`   // :"message"
	Topic   string    `json:"topic"`   // :"Server"
	Message string    `json:"message"` // :"triggered"
}
