package main

import (
	tempconv "Calc/tempconv"
	"fmt"
	"strings"
)

type Calculator struct {
	functions []tempconv.MathFunction
}

func (c *Calculator) addMathFunction(m tempconv.MathFunction) {
	c.functions = append(c.functions, m)
}

func (c *Calculator) doCalculation(name string, arg float64) float64 {
	var result float64
	for _, f := range c.functions {
		if strings.EqualFold(name, f.GetName()) {
			result = f.Calculate(arg)
			return result //, nil
		}
	}
	return 0
}

func main() {
	myCalculator := Calculator{}

	var degree float64
	var t int

	fmt.Println("Which conversion dou you want ?")
	fmt.Println("Enter 1 for Celcius to Fahrenheit conversion")
	fmt.Println("Enter 2 for Fahrenheit to Celcius conversion")
	fmt.Println("Enter 3 for Celcius to Kelvin conversion")
	fmt.Println("Enter 4 for Kelvin to Celcius conversion")
	fmt.Println("Enter 5 for Fahrenheit to Kelvin conversion")
	fmt.Println("Enter 6 for Kelvin to Fahrenheit conversion")

	fmt.Scanln(&t)
	fmt.Println("Enter Degree")
	fmt.Scanln(&degree)

	switch t {
	case 1:
		myCalculator.addMathFunction(tempconv.CToF{Degree: degree, Name: "CelcToFahr"})
		fmt.Printf("Conversion Result : %f", myCalculator.doCalculation("CelcToFahr", degree))
	case 2:
		myCalculator.addMathFunction(tempconv.FToC{Degree: degree, Name: "FahrToCelc"})
		fmt.Printf("Conversion Result : %f", myCalculator.doCalculation("FahrToCelc", degree))
	case 3:
		myCalculator.addMathFunction(tempconv.CToK{Degree: degree, Name: "CelcToKel"})
		fmt.Printf("Conversion Result : %f", myCalculator.doCalculation("CelcToKel", degree))
	case 4:
		myCalculator.addMathFunction(tempconv.KToC{Degree: degree, Name: "KelToCelc"})
		fmt.Printf("Conversion Result : %f", myCalculator.doCalculation("KelToCelc", degree))
	case 5:
		myCalculator.addMathFunction(tempconv.FToK{Degree: degree, Name: "FahrToKel"})
		fmt.Printf("Conversion Result : %f", myCalculator.doCalculation("FahrToKel", degree))
	case 6:
		myCalculator.addMathFunction(tempconv.KToF{Degree: degree, Name: "KelToFahr"})
		fmt.Printf("Conversion Result : %f", myCalculator.doCalculation("KelToFahr", degree))
	}

}
