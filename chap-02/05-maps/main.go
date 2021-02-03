package main

import "fmt"

func main() {
	languages := map[int]string{
		1: "English",
		2: "French",
		3: "Spanish",
		4: "Hindi",
	}

	fmt.Println("Languages in map:", languages)

	fmt.Println("Range over Languages")
	for key, value := range languages {
		fmt.Println("Language", key, "has value of", value)
	}

	// using make func
	products := make(map[int]string)
	products[1] = "chair"
	products[2] = "mobile phone"
	products[3] = "laptop"
	products[4] = "tv"

	fmt.Println("Product with id 1:", products[1])

	for key, val := range products {
		fmt.Println("Product with key", key, "has value of", val)
	}

	fmt.Println("Deleting single product")
	delete(products, 1)
	for key, val := range products {
		fmt.Println("Product with key", key, "has value of", val)
	}
}
