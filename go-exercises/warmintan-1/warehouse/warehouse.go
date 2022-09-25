package warehouse

import (
	"fmt"
	"warmintan-1/chef"
)

func GetStockFromChef(stock string) map[string]string {
	stockfc := chef.GetStockFromWaiter(stock)
	v := stockfc[""]
	fmt.Println("Third,", v, "is moved into Warehouse")
	// fmt.Println("Warehouse has received stock from Chef")
	return stockfc
}
