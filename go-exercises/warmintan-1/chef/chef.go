package chef

import (
	"fmt"
	"warmintan-1/waiter"
)

func GetStockFromWaiter() map[string]string {
	stockfw := waiter.AddStock()
	v := stockfw[""]
	fmt.Println(v, "is received by Chef")

	// if v != "" {
	// 	warehouse.GetStockFromChef()
	// }

	return stockfw
}
