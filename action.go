package gotfy

// ActionButtonType specifies all the currently supported action buttons
// that you can use for a notification
type ActionButtonType byte

const (
	UnspecifiedAction ActionButtonType = iota
	View
	HTTP
	Broadcast
)

type ActionButton interface {
	actionType() ActionButtonType
}
