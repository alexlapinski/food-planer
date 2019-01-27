package generator

import (
	"fmt"

	"github.com/alexlapinski/menu-planner/pkg/inventory"
)

func Execute() {
	fmt.Println("Food-Plan Generator")
	fmt.Println("Reading Inventory")

	inventory.Read("./assets/inventory.json")
}
