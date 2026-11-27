package sdk

import "testing"

func TestAccessString(t *testing.T) {
	tests := []struct {
		in   Access
		want string
	}{
		{Public, "public"},
		{User, "user"},
		{Admin, "admin"},
		{Internal, "internal"},
		{Access(0), "access(0)"},
		{Access(9), "access(9)"},
	}
	for _, tc := range tests {
		if got := tc.in.String(); got != tc.want {
			t.Errorf("Access(%d).String() = %q, want %q", int(tc.in), got, tc.want)
		}
	}
}
