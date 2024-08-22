package addresses

import "strings"

// GetAddressType returns the type of the address
func GetAddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	addressInLowerCase := strings.ToLower(address)
	firstWord := strings.Split(addressInLowerCase, " ")[0]

	addressIsValid := false;
	for _, value := range validTypes {
		if value == firstWord {
			addressIsValid = true;
			break
		}
	}

	if addressIsValid {
		return firstWord
	}

	return "invalid_type"
}