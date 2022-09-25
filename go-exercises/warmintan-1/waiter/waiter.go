package waiter

import (
	"fmt"
)

type Waiter struct {
	Stock map[string]string
}

func AddStock(stock string) map[string]string {

	// var stock string
	fmt.Println("Please add stock: ")
	fmt.Scanln(&stock)

	stock1 := make(map[string]string)

	stock1[""] = stock
	v := stock1[""]
	fmt.Println("First,", v, "is received by Waiter")
	// fmt.Print(stock1, "\n")

	// if v != "" {
	// 	chef.GetStockFromWaiter()
	// }
	return stock1
}
