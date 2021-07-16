package main

import "fmt"

func main() {
	favRecords := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "sunsets"},
	}
	//adding record
	favRecords["James"] = []string{"Jill", "Comics", "Code"}

	//deleting a record
	toBeDeleted := "no_dr"
	if _, ok := favRecords[toBeDeleted]; ok {
		delete(favRecords, toBeDeleted)
		fmt.Println("Deleted from favRecords", toBeDeleted)

	}

	for index, value := range favRecords {
		fmt.Printf("%v %v\n", index, value)
	}

}
