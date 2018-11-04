package identifier

import (
	"fmt"
	"sync"

	"github.com/opsidian/basil/util"
)

// Registry stores and generates identifiers
type Registry struct {
	ids       map[string]struct{}
	minLength int
	maxLength int
	lock      sync.RWMutex
}

// NewRegistry creates a new id registry
func NewRegistry(minLength int, maxLength int) *Registry {
	return &Registry{
		ids:       map[string]struct{}{},
		minLength: minLength,
		maxLength: maxLength,
	}
}

// IDExists returns true if the identifier already exists
func (r *Registry) IDExists(id string) bool {
	r.lock.RLock()
	defer r.lock.RUnlock()

	_, exists := r.ids[id]
	return exists
}

// GenerateID generates and registers a new block id
// If there are many block ids generated with the current minimum length and it's getting hard to generate unique ones
// thenb the min length will be increased by one (up to the maximum length)
func (r *Registry) GenerateID() string {
	util.SeedMathRand()

	tries := 0
	for {
		id := util.RandHexString(r.minLength, true)
		err := r.RegisterID(id)
		if err == nil {
			return id
		}
		tries++
		if tries == 3 && r.minLength < r.maxLength {
			r.minLength++
			tries = 0
		}
	}
}

// RegisterID registers a new block id and returns an error if it is already taken
func (r *Registry) RegisterID(id string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.ids == nil {
		r.ids = map[string]struct{}{}
	}

	_, exists := r.ids[id]
	if exists {
		return fmt.Errorf("%q identifier already exists", id)
	}

	r.ids[id] = struct{}{}

	return nil
}
