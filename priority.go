package gotfy

// Priority is an enum for the message priority
type Priority int8

//go:generate enumer -type Priority -transform lower
const (
	UnspecifiedPriority Priority = iota
	Min
	Low
	Default
	High
	Max
)
