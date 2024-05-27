package address_test

import (
	. "intro/address" // dot is an alias for the default package imported.
	"testing"
)

type scenario struct {
	input    string
	expected string
}

func TestAddressType(t *testing.T) {
	t.Parallel() // Execute function in parallel with other tests
	scenarios := []scenario{
		{"10 St", "St"},
		{"Kirkman Road", "Road"},
		{"Main Square", "Square"},
		{" ", "Unrecognized type"},
	}

	for _, scen := range scenarios {
		actual := AddressType(scen.input)

		if actual != scen.expected {
			t.Errorf("Expected %s, actual %s", scen.expected, actual)
		}

	}
}

func TestAnotherFunction(t *testing.T) {
	t.Parallel() // Execute function in parallel with other tests
	if 1 > 2 {
		t.Error("Impossible")
	}
}
