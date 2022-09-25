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
	fmt.Println("First,", v, "is added by Waiter")
	// fmt.Print(stock1, "\n")

	// if v != "" {
	// 	chef.GetStockFromWaiter()
	// }
	return stock1
}

func TakeOrder(order string) map[string]string {
	// var stock string
	fmt.Println("Please take order: ")
	fmt.Scanln(&order)

	order1 := make(map[string]string)

	order1[""] = order
	v := order1[""]
	fmt.Println("First,", v, "order taken by Waiter")
	// fmt.Print(stock1, "\n")

	return order1
}
