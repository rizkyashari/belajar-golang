package main

import (
	"fmt"
	"warmintan-1/chef"
	"warmintan-1/warehouse"
)

func main() {

	var init int

	fmt.Println("Welcome, Waiter!\nPlease choose menu below:")
	fmt.Println("1. Add stock\n2. Take order")
	fmt.Scanln(&init)

	if init == 1 {
		// waiter.AddStock()
		// chef.GetStockFromWaiter()
		warehouse.GetStockFromChef("")

	} else if init == 2 {
		// waiter.TakeOrder("")
		chef.GetOrderFromWaiter("")
	}

}
