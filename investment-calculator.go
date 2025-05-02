package main

import (
	"fmt"
	"math"
)

func main(){
	const inflationRate = 6.5
	/*var investmentAmount float64 = 1000*/
	var investmentAmount float64
	var years float64 = 10
	expectedReturnRate := 5.5

	fmt.Print("please input investmentAmount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100,years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
