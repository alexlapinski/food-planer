package main

import "fmt"
import "github.com/alexlapinski/food-planer/inventory"

func main() {
	fmt.Println("Food-Plan Generator")
	fmt.Println("Reading Inventory")

	inventory.ReadFromDisk("inventory.json")
}
