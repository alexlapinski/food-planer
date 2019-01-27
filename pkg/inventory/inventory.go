package inventory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Inventory struct {
	Pantry []InventoryItem `json:"pantry"`
	Fridge []InventoryItem `json:"fridge"`
}

type InventoryItem struct {
	Name  string
	Count int
}

func Read(filepath string) {
	jsonFile, err := os.Open(filepath)

	if err != nil {
		// TODO: Add filename to message
		fmt.Println("Error Reading file", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var inventory Inventory

	json.Unmarshal(byteValue, &inventory)
	fmt.Printf("Successfully Read Inventory from '%v'\n", filepath)

	fmt.Println("Items in the fridge:")
	for i := 0; i < len(inventory.Fridge); i++ {
		item := inventory.Fridge[i]
		fmt.Printf("\t%v (%v)\n", item.Name, item.Count)
	}

	fmt.Println("Items in the pantry:")
	for i := 0; i < len(inventory.Pantry); i++ {
		item := inventory.Pantry[i]
		fmt.Printf("\t%v (%v)\n", item.Name, item.Count)
	}
}
