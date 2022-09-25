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
