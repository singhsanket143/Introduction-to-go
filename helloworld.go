package main

import "fmt"

func main() {

	const availableStock = 50

	var productName string         // Declare a variable to hold the product name but no value is assigned
	var productPrice int = 10000   // Declare a variable to hold the product price and assign a value
	var companyName = "Apple Inc." // Declare a variable to hold the company name and assign a value but type is inferred

	category := "Electronics" // Declare a variable to hold the category and assign a value but type is inferred

	fmt.Println("Product name is", productName, "and price is", productPrice, "and company name is", companyName)
	fmt.Println("Category is", category)

	productName = "iPhone 15 Pro" // Assign a value to the product name variable

	loops_demo()
	arrays_demo()
	maps_demo()

	x, y := check_odd_even(11)

	fmt.Println(x, y)

	demo_pointers()
}

func loops_demo() {

	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	for i := range 3 {
		fmt.Println("Range loop iteration:", i)
	}

	for i, char := range "Sanket" {
		fmt.Println("String range loop iteration:", i, char)
	}

	// Below is a while loop using a for loop

	j := 10

	for j > 0 {
		if j == 3 {
			j--
			continue
		} else if j == 5 {
			fmt.Println("Skipping 5")
			j--
			continue
		} else {
			fmt.Println("While loop iteration:", j)
			j--
			continue
		}

	}

	for {
		fmt.Println("Infinite loop iteration")
		break
	}
}

func arrays_demo() {

	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	var arr1 []int // slice of empty integers

	arr1 = append(arr1, 1, 2, 3, 4, 5)

	fmt.Println("Array 1:", arr1)

	for i, value := range arr1 {
		fmt.Println("Index:", i, "Value:", value)
	}

}

func maps_demo() {
	productprices := map[string]int{
		"iPhone 15 Pro": 1000,
		"iPhone 15":     800,
		"iPhone 14":     600,
	}

	fmt.Println("Product Prices:", productprices)

	customMap := map[string]string{}

	fmt.Println("Custom Map:", customMap)

	emptyMap := make(map[string]int)

	fmt.Println("Empty Map:", emptyMap)

	emptyMap["key1"] = 100
	emptyMap["key2"] = 200
	emptyMap["key3"] = 300
	fmt.Println("Populated Empty Map:", emptyMap)

	for key, value := range emptyMap {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func check_odd_even(num int) (string, int) {
	if num%2 == 0 {
		return "Even", 0
	} else {
		return "Odd", 1
	}
}

func demo_pointers() {

	i := 120

	var ptr *int = &i // this pointer can store the address of an integer variable

	ptr1 := &i

	fmt.Println("Value of i:", i, "Pointer to i:", ptr, "Pointer1 to i:", ptr1)

	fmt.Println("value present at pointer", *ptr)
}
