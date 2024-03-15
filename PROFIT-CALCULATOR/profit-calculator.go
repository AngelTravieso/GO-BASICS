package main

import (
	"fmt"
)

func main() {

	var renevue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Renevue: ")
	fmt.Scan(&renevue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("TaxRate: ")
	fmt.Scan(&taxRate)

	ebt := renevue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("Ebt:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)

}
