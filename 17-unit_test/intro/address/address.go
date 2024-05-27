package address

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AddressType verify if an address has a valid type and return
// Function must start with Upper case to be exportable. Comment must start with function name.
func AddressType(address string) string {
	validTypes := []string{"st", "road", "avenue", "square"}
	normalized := strings.ToLower(address)
	addrType := strings.Split(normalized, " ")[1]

	validAddress := false

	for _, validType := range validTypes {
		if validType == addrType {
			validAddress = true
		}
	}

	if validAddress {
		caser := cases.Title(language.AmericanEnglish)
		return caser.String(addrType)
	}

	return "Unrecognized type"

}
