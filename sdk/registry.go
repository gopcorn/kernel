// sdk/registry.go : how features announce themselves
package sdk

import (
	"fmt"
	"sort"
	"sync"
)

var (
	mu         sync.Mutex
	registered = map[string]Feature{}
)

// Register adds a feature. Call it from the feature package's init function.
// It panics on an empty/ repeated name at startup.
func Register(f Feature) {
	name := f.Name()
	if name == "" {
		panic("sdk: Register called with a feature without a name")
	}

	mu.Lock()
	defer mu.Unlock()
	if _, duplicateCheck := registered[name]; duplicateCheck {
		panic(fmt.Sprintf("sdk: feature %q registered twice", name))
	}

	// feature not yet registered, add it to make.
	registered[name] = f
}

// Registered returns every registered feature, sorted by name.
func Registered() []Feature {
	mu.Lock()
	defer mu.Unlock()

	names := make([]string, 0, len(registered))
	for name := range registered {
		names = append(names, name)
	}
	sort.Strings(names)
	out := make([]Feature, 0, len(names))
	for _, name := range names {
		out = append(out, registered[name])
	}
	return out
}
