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
	if ConvertFromInch {
		ConvertToMM()
	} else {
		ConvertToInch()
	}
}

func getFlags() {
	flag.BoolVar(&ConvertFromInch, "i", false, "Convert to inch")
	flag.Parse()
}

func ConvertToMM() {
	value := os.Args[2]
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
	valueToFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	inch := valueToFloat * 0.0393701
	inchToString := strconv.FormatFloat(inch, 'f', 2, 64)
	fmt.Println(inchToString + "in")
}
