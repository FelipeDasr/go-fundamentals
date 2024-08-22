package addresses

import "testing"

type TestPayload struct {
	value string
	expectedValue string
}

func TestGetAddressType(t *testing.T) {
	// "rua", "avenida", "estrada", "rodovia"

	tests := []TestPayload{
		{"Rua dos Bobos", "rua"},
		{"Avenida Paulista", "avenida"},
		{"Estrada dos Alvarengas", "estrada"},
		{"Rodovia dos Imigrantes", "rodovia"},
		{"RuA dos Bobos", "rua"},
		{"AVENIDA PAULISTA", "avenida"},
		{"Praça DOS ALVARENGAS", "invalid_type"},
	}

	for _, test := range tests {
		result := GetAddressType(test.value)
		if result != test.expectedValue {
			t.Errorf("Expected %s, got %s", test.expectedValue, result)
		}
	}
}