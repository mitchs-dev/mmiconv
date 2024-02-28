/*
Package mmiconv is a simple package to convert millimeters to inches and vice versa.

By default, the package will convert from millimeters to inches.
To convert from inches to millimeters, use the -i flag.
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var ConvertFromInch bool

func main() {
	getFlags()
	checkIfEmpty()
	if ConvertFromInch {
		ConvertToMM()
	} else {
		ConvertToInch()
	}
}

func checkIfEmpty() {
	var expectedArgs int
	if ConvertFromInch {
		// executable path + -i flag + value
		expectedArgs = 3
	} else {
		// executable path + value
		expectedArgs = 2
	}
	if len(os.Args) < expectedArgs {
		fmt.Println("Error: Please provide a value to convert")
		fmt.Println("Usage: mmiconv [-i] [value]")
		os.Exit(1)
	}
}

func getFlags() {
	flag.BoolVar(&ConvertFromInch, "i", false, "Convert to inch")
	flag.Parse()
}

func ConvertToMM() {
	value := os.Args[2]
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		fmt.Println("Error: Conversion value must be a number")
		return
	}
	valueToFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	mm := valueToFloat * 25.4
	mmToString := strconv.FormatFloat(mm, 'f', 2, 64)
	fmt.Println(mmToString + "mm")
}

func ConvertToInch() {
	value := os.Args[1]
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		fmt.Println("Error: Conversion value must be a number")
		return
	}
	valueToFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	inch := valueToFloat * 0.0393701
	inchToString := strconv.FormatFloat(inch, 'f', 2, 64)
	fmt.Println(inchToString + "in")
}
