package warehouse

import (
	"fmt"
	"warmintan-1/chef"
)

func GetStockFromChef() {
	stockfc := chef.GetStockFromWaiter()
	v := stockfc[""]
	fmt.Println(v, "is received by Warehouse")
	// fmt.Println("Warehouse has received stock from Chef")
}
