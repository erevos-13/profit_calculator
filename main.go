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

	profit := revenue - expences
	tax := profit * taxRate / 100
	profitAfterTax := profit - tax
	profitBeforeTax := profit + tax
	ratio := profitAfterTax / revenue
	fmt.Println("Profit after tax: ", profitAfterTax)
	fmt.Println("Profit before tax: ", profitBeforeTax)
	fmt.Println("Ratio:  ", ratio)

}
