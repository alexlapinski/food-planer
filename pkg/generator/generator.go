package generator

import (
	"fmt"

	"github.com/alexlapinski/menu-planner/pkg/inventory"
)

func execute() {
	fmt.Println("Food-Plan Generator")
	fmt.Println("Reading Inventory")

	inventory.Read("inventory.json")
}
