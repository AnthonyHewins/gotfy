package gotfy

// Priority is an enum for the message priority
type Priority int8

const (
	UnspecifiedPriority Priority = iota
	Min
	Low
	Default
	High
	Max
)
