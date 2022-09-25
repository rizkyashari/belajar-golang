package chef

import (
	"fmt"
	"warmintan-1/waiter"
)

func GetStockFromWaiter(stock string) map[string]string {
	stockfw := waiter.AddStock(stock)
	v := stockfw[""]
	fmt.Println("Second,", v, "is received by Chef")

	// if v != "" {
	// 	warehouse.GetStockFromChef()
	// }

	return stockfw
}

func GetOrderFromWaiter(order string) map[string]string {
	orderfw := waiter.TakeOrder(order)
	v := orderfw[""]
	fmt.Println("Second,", v, "order is received by Chef.\nChef is cooking", v, ".")

	return orderfw
}
