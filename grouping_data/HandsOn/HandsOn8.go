package main

import "fmt"

func main() {
	favRecords := map[string][]string{
		"bond_james": {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr": {"Being evil", "Ice cream", "sunsets"},
	}

	for index, value := range favRecords{
		fmt.Printf("%v %v\n", index, value)
	}
}