package debug

import (
	"testing"

	"github.com/coredns/coredns/core/dnsserver"

	"github.com/mholt/caddy"
)

func TestDebug(t *testing.T) {
	tests := []struct {
		input         string
		shouldErr     bool
		expectedDebug bool
	}{
		// positive
		{`debug`, false, true},
		{`debug {
			memory
		}`, false, true},
		// negative
		{`debug off`, true, false},
		{`debug {
			memori
		}`, true, true},
	}

	for i, test := range tests {
		c := caddy.NewTestController("dns", test.input)
		err := setup(c)
		cfg := dnsserver.GetConfig(c)

		if test.shouldErr && err == nil {
			t.Fatalf("Test %d: Expected error but found none for input %s", i, test.input)
		}

		if err != nil {
			if !test.shouldErr {
				t.Fatalf("Test %d: Expected no error but found one for input %s. Error was: %v", i, test.input, err)
			}
		}
		if cfg.Debug != test.expectedDebug {
			t.Fatalf("Test %d: Expected debug to be: %t, but got: %t, input: %s", i, test.expectedDebug, cfg.Debug, test.input)
		}
	}
}
