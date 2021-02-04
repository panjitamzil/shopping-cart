package main

import (
  	"fmt"
	"log"
)

func main() {
	var (
		user, item, redisKey string
		menu, price int
		exit = true
	)

	fmt.Println("Please login!")
	fmt.Println("Username : ")
	fmt.Scan(&user)
	cart := newCart()

	for exit {

		fmt.Println("Menu")
		fmt.Println("1. Add Item")
		fmt.Println("2. Remove Item")
		fmt.Println("3. List Cart")
		fmt.Println("4. Delete Cart")
		fmt.Println("5. Exit")
		fmt.Println("Input your number of menu : ")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			fmt.Println("Name of item : ")
			fmt.Scan(&item)
			fmt.Println("Price : ")
			fmt.Scan(&price)
			addItem(item,price,cart)
		case 2:
			fmt.Println("Name of item (delete) : ")
			fmt.Scan(&item)
			removeItem(cart, item)		
		case 3:
			listCart(cart)
		case 4:
			deleteCart(cart)
		case 5:
			exit = false
		default:
			fmt.Println("Invalid menu! Please re-run the program")
			exit = false
		}
	}

}

func newCart() map[string]int {
	return map[string]int{}
}

func addItem(item string, price int, cart map[string]int) {
	cart[item] = price
}

func listCart(cart map[string]int) {
	log.Println("List of Cart")
	for item, price := range cart {
		log.Printf("Item : %s, Price : Rp %d", item, price)
	}
}

func removeItem(cart map[string]int, item string) {
	delete(cart, item)
}

func deleteCart(cart map[string]int) {
	for k := range cart {
		delete(cart, k)
	}
}