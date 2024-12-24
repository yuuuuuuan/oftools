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

	// Return the Conversion struct with all the systems
	return Conversion{
		Hex: hexValue,
		Dec: decimalValue,
		Bin: binaryValue,
		Oct: octalValue,
	}, nil
}

// ConvertDecToOthers converts a decimal value to hex, binary, and octal
func ConvertDecToOthers(decimalValue int64) Conversion {
	// Convert decimal to hex, binary, and octal
	hexValue := fmt.Sprintf("0x%X", decimalValue)
	binaryValue := fmt.Sprintf("%b", decimalValue)
	octalValue := fmt.Sprintf("%o", decimalValue)

	return Conversion{
		Hex: hexValue,
		Dec: decimalValue,
		Bin: binaryValue,
		Oct: octalValue,
	}
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

// Display prints all the converted values
// func (c Conversion) Display() {
// 	fmt.Printf("Hex: %s\n", c.Hex)
// 	fmt.Printf("Decimal: %d\n", c.Dec)
// 	fmt.Printf("Binary: %s\n", c.Bin)
// 	fmt.Printf("Octal: %s\n", c.Oct)
// }

// func main() {
// 	// Example: Convert a hex value to all other systems
// 	hexValue := "0x1a"
// 	conversion, err := ConvertHexToOthers(hexValue)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	conversion.Display()

// 	// Example: Convert a decimal value to all other systems
// 	decValue := int64(26)
// 	conversion = ConvertDecToOthers(decValue)
// 	conversion.Display()

// 	// Example: Convert a binary value to all other systems
// 	binValue := "11010"
// 	conversion, err = ConvertBinToOthers(binValue)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	conversion.Display()

// 	// Example: Convert an octal value to all other systems
// 	octalValue := "32"
// 	conversion, err = ConvertOctToOthers(octalValue)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	conversion.Display()
// }
