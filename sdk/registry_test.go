package sdk

import "testing"

// reset empties the registry so each test starts clean.
func reset() {
	mu.Lock()
	defer mu.Unlock()
	registered = map[string]Feature{}
}

// fake is the smalles possible feature: its name and value
type fake string

func (f fake) Name() string  { return string(f) }
func (fake) Init(Host) error { return nil }

func TestRegisteredIsSortedByName(t *testing.T) {
	reset()
	Register(fake("b"))
	Register(fake("a"))
	Register(fake("c"))

	got := ""
	for _, f := range Registered() {
		got += f.Name()
	}
	if got != "abc" {
		t.Fatalf("order = %q, wanr %q", got, "abc")
	}
}

func TestRegisterTwicePanics(t *testing.T) {
	reset()
	Register(fake("a"))
	defer func() {
		if recover() == nil {
			t.Fatal("expected a panic for a repeated name")
		}
	}()
	Register(fake("a"))
}

func TestRegisterEmptyNamePanics(t *testing.T) {
	reset()
	defer func() {
		if recover() == nil {
			t.Fatal("expected a panic for an empty name")
		}
	}()
	Register(fake(""))
}
