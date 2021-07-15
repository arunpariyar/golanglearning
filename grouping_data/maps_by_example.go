package main

import "fmt"

func main() {
	car := "Ferrari"
	var carCatalog = map[string]string{
		"BMW":   "germany",
		"Tesla": "usa",
	}
	//adding Ferrari into our carCatalog map
	carCatalog["Ferrari"] = "france"

	fmt.Println(car)

	for key := range carCatalog {
		if key == car {
			fmt.Println("Your car is in our catalog")
		}
	}

	//deleting something from the map
	//delete(carCatalog, "Tesla")
	// note that it is also possible to delete something that doesn't exist so better way to delete something from maps is to check that we actually have the key using ,k idiom and then delete it

	if _, ok := carCatalog[car]; ok {
		fmt.Printf("%s is in the the catalog\n", car)
		delete(carCatalog, car)
		fmt.Printf("%s is now deleted\n", car)

	}

	//printing all the elements the carCatalog
	fmt.Println(carCatalog)

}
