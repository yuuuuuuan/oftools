package algorithm

import (
	"fmt"
	"strconv"
	"strings"
)

// Conversion struct to store values in all numeral systems
type Conversion struct {
	Hex string
	Dec int64
	Bin string
	Oct string
}

// ConvertHexToOthers converts a hexadecimal value to decimal, binary, and octal
func ConvertHexToOthers(hexValue string) (Conversion, error) {
	// Remove the '0x' prefix if it exists
	hexValue = strings.TrimPrefix(hexValue, "0x")

	// Convert hex to decimal
	decimalValue, err := strconv.ParseInt(hexValue, 16, 64)
	if err != nil {
		return Conversion{}, fmt.Errorf("invalid hex value: %v", err)
	}

	// Convert decimal to binary and octal
	binaryValue := fmt.Sprintf("%b", decimalValue)
	octalValue := fmt.Sprintf("%o", decimalValue)
	hexValue = fmt.Sprintf("0x%X", decimalValue)

	// Return the Conversion struct with all the systems
	return Conversion{
		Hex: hexValue,
		Dec: decimalValue,
		Bin: binaryValue,
		Oct: octalValue,
	}, nil
}

// ConvertDecToOthers converts a decimal value to hex, binary, and octal
func ConvertDecToOthers(decimalValue string) (Conversion, error) {
	value, err := strconv.ParseInt(decimalValue, 10, 64)
	if err != nil {
		return Conversion{}, fmt.Errorf("Error:%v", err)

	}
	// Convert decimal to hex, binary, and octal
	hexValue := fmt.Sprintf("0x%X", value)
	binaryValue := fmt.Sprintf("%b", value)
	octalValue := fmt.Sprintf("%o", value)

	return Conversion{
		Hex: hexValue,
		Dec: value,
		Bin: binaryValue,
		Oct: octalValue,
	}, nil
}

// ConvertBinToOthers converts a binary value to hex, decimal, and octal
func ConvertBinToOthers(binaryValue string) (Conversion, error) {
	// Convert binary to decimal
	decimalValue, err := strconv.ParseInt(binaryValue, 2, 64)
	if err != nil {
		return Conversion{}, fmt.Errorf("invalid binary value: %v", err)
	}

	// Convert to hex and octal
	hexValue := fmt.Sprintf("0x%X", decimalValue)
	octalValue := fmt.Sprintf("%o", decimalValue)

	return Conversion{
		Hex: hexValue,
		Dec: decimalValue,
		Bin: binaryValue,
		Oct: octalValue,
	}, nil
}

// ConvertOctToOthers converts an octal value to hex, decimal, and binary
func ConvertOctToOthers(octalValue string) (Conversion, error) {
	// Convert octal to decimal
	decimalValue, err := strconv.ParseInt(octalValue, 8, 64)
	if err != nil {
		return Conversion{}, fmt.Errorf("invalid octal value: %v", err)
	}

	// Convert to hex and binary
	hexValue := fmt.Sprintf("0x%X", decimalValue)
	binaryValue := fmt.Sprintf("%b", decimalValue)

	return Conversion{
		Hex: hexValue,
		Dec: decimalValue,
		Bin: binaryValue,
		Oct: octalValue,
	}, nil
}
