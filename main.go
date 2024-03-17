package main

import "fmt"

func main() {
	var revenue float64
	var expences float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expences: ")
	fmt.Scan(&expences)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expences
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Profit before tax: ", profit)
	fmt.Println("Ratio:  ", ratio)

}
