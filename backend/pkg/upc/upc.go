package upc

import (
	"unicode"
)

// UPC Code's component parts
type UPC struct {
	System       string
	Manufacturer string
	Product      string
	Check        string
}

// parse a UPC code into its component parts
func ParseUPC(str string) (*UPC, bool) {
	if !IsUPC(str) {
		return nil, false
	}

	upc := &UPC{
		System:       str[0:1],
		Manufacturer: str[1:6],
		Product:      str[6:11],
		Check:        str[11:12],
	}

	return upc, true
}

// check if a string is a valid UPC-A code
func IsUPC(str string) bool {
	if len(str) != 12 {
		return false
	}
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	// Validate check digit
	sum := 0
	for i := 0; i < 11; i++ {
		digit := int(str[i] - '0')
		if i%2 == 0 {
			sum += digit * 3
		} else {
			sum += digit
		}
	}
	expectedCheck := (10 - (sum % 10)) % 10
	actualCheck := int(str[11] - '0')

	return expectedCheck == actualCheck
}
