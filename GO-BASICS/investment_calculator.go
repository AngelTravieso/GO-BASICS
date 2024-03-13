package main

import (
	"fmt"
	"math"
)

func main() {

	// investmentAmount = 1000, years = 10
	var investmentAmount float64 = 1000.0
	years := 10.0
	expectedReturnRate := 5.5

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	fmt.Println(futureValue)

}
