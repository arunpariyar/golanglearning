// maps are essentially key value pair that is unordered and its best for situations when we need to store and find things with unique keys

package main

import "fmt"

func main() {
	phonebook := map[string]int{
		"jack":  1234,
		"james": 2326,
		"jenny": 3456,
	}
	fmt.Println(phonebook)

	//to display if the phone number of jack

	fmt.Println(phonebook["jack"])

	fmt.Println(phonebook["jeffery"]) // this results in a 0 even though the key is not present
	//however it is possibel to search for something that not in the map and we get a 0 return, this can be particularly confusing because that doesnt give us the sureity of weather or not the key is present in the map to check this the , ok idiom can be used

	value, ok := phonebook["jeffery"]

	fmt.Println(value) // prints the value
	fmt.Println(ok)    // return a boolean, true if the key is present and false if not

	// a comman pattern thats see to work with the above seems to be

	if _, ok := phonebook["jack"]; ok {
		fmt.Println("This line runs because the key is present")
	}

}
