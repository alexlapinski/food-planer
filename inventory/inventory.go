package inventory

import (
	"fmt"
	"os"
)

type InventoryItem struct {
	Name  string
	Count int
}

func ReadFromDisk(filepath string) {
	jsonFile, err := os.Open(filepath)

	if err != nil {
		// TODO: Add filename to message
		fmt.Println("Error Reading file", err)
	}

	// TODO: Add filename to message
	fmt.Println("Successfully Opened file")

	defer jsonFile.Close()
}
