package blake2b

import (
	"v2ray.com/core/app/porter/src/crypto"
)

// Noop is a HashPolicy that returns the original input.
type Noop struct{}

var (
	_ crypto.HashPolicy = (*Noop)(nil)
)

// New returns a hash policy that is a no-op.
func New() *Noop {
	return &Noop{}
}

// HashBytes returns the input bytes.
func (p *Noop) HashBytes(bytes []byte) []byte {
	return bytes
}
