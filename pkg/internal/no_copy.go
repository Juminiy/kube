package internal

// NoCopy
// referred from: sync.noCopy
type NoCopy struct{}

func (*NoCopy) Lock()   {}
func (*NoCopy) Unlock() {}
