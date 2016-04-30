package idgen

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"sync"
)

var globalIDs = &lockedSource{src: New()}

// Source represents a source of unique id generation.
type Source interface {
	ID() string
}

type lockedSource struct {
	m   sync.Mutex
	src Source
}

func (l *lockedSource) ID() (id string) {
	l.m.Lock()
	id = l.src.ID()
	l.m.Unlock()
	return
}

// ID returns a unique id on every call. The unique id is 40 characters long
// and is hexadecimal encoded. Safe for concurrent access.
func ID() string {
	return globalIDs.ID()
}

// IDSource type holds state for generating new ids.
type IDSource struct {
	seed    []byte
	counter int64
}

// New creates and initializes a new IDSource type.
func New() *IDSource {
	return &IDSource{seed: seed()}
}

// ID returns a unique id on every call. The unique id is 40 characters long
// and is hexadecimal encoded. Not safe for concurrent access, use idgen.ID() or
// pass a separate idgen.New() to each goroutine.
func (id *IDSource) ID() string {
	h := hmac.New(sha1.New, id.seed)
	h.Write([]byte(string(id.counter)))
	id.counter = id.counter + 1
	return hex.EncodeToString(h.Sum(nil))
}

// seed creates a 160 bit seed.
func seed() []byte {
	s := 160 / 8
	b := make([]byte, s)
	n, err := rand.Read(b)
	if err != nil {
		panic(err.Error())
	} else if n != s {
		panic("seed: not enough bytes read from crypto/rand.Read")
	} else if bytes.Equal(b, make([]byte, s)) {
		panic("seed: bytes read from crypto/rand.Read are all zeros")
	}
	return b
}
