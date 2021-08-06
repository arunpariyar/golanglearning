// TipCalculator
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	fmt.Println("************** Tip Calculator v1.0 *****************")

	//implementing the reader interface to get console input
	reader := bufio.NewReader(os.Stdin)
	
	//   Prompt for billAmount with "What is the bill amount?"
	fmt.Println("What is the bill amount ?")

	billAmountString, _ := reader.ReadString('\n')
	billAmountString = strings.Replace(billAmountString, "\n", "", -1)
	
	//   Prompt for tipRate with "What is the tip rate?"
	fmt.Println("What is the tip rate ?")
	
	tipRateString, _ := reader.ReadString('\n')
	tipRateString = strings.Replace(tipRateString, "\n", "", -1)

	// converting billAmount and TipRateString to float64
	billAmount, _:= strconv.ParseFloat(billAmountString, 64)
	tipRate,_ := strconv.ParseFloat(tipRateString,64 )

	//calculating the tip
	tip := billAmount * (tipRate / 100)

	//   Display "Tip: $" + tip
	fmt.Println("Tip:$", tip)
	//   Display "Total: $" + total
	fmt.Println("Total : $", billAmount+tip)
	// End
}
