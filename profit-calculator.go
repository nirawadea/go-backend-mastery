package main

import (
	"fmt"
)

func main(){
	/*const inflationRate = 6.5*/
	/*var investmentAmount float64 = 1000*/
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("please input Revenu Amount: ")
	fmt.Scan(&revenue)

	fmt.Print("Your Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Your Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := float64(ebt) * (1-taxRate/100)
	ratio := ebt/profit
	

	fmt.Print("EBT: ")
	fmt.Println(ebt)
	fmt.Print("Profit: ")
	fmt.Println(profit)
	fmt.Print("Ratio: ")
	fmt.Println(ratio)
}
