package internal

// NoCmp
// referred from: https://github.com/uber-go/atomic/blob/master/nocmp.go
type NoCmp [0]func()
