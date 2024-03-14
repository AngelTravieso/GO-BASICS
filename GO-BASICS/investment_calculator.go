package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5

	// investmentAmount = 1000, years = 10
	var investmentAmount float64 // if no have value there a null value (0.00 for float64)
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")

	// fmt.Scan to take user input
	fmt.Scan(&investmentAmount) // & -> pointers

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
