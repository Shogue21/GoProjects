package main

import (
	"fmt"
	"strings"
)

const (
	newBookCost = 6.0
	oldBookCost = 3.5
)

func computeCost(new, old float64) float64 {
	return new*newBookCost + old*oldBookCost
}

func applyCouponDiscount(cost float64, coupon string) float64 {
	if coupon == "free-stuff" {
		return 0
	} else if coupon == "half-off" {
		return cost * 0.5
	} else {
		return cost
	}
}

func isValidCoupon(coupon string) bool {
	return bool(coupon == "free-stuff" || coupon == "half-off")
}
func main() {

	fmt.Println("Welcome to Bargain Books!")
	fmt.Printf("New books are $%.2f each.\n", newBookCost)
	fmt.Printf("Old books are $%.2f each.\n", oldBookCost)

	fmt.Println("How many new books are you buying today?")
	var newBooks int
	fmt.Scanln(&newBooks)

	fmt.Println("How many old books are you buying today?")
	var oldBooks int
	fmt.Scanln(&oldBooks)
	total := computeCost(float64(newBooks), float64(oldBooks))

	fmt.Println("Do you have a coupon? [Y/N]: ")
	var hasCoupon string
	fmt.Scanln(&hasCoupon)
	hasCoupon = strings.ToLower(hasCoupon)

	if hasCoupon == "y" {
		fmt.Println("What is your coupon? ")
		var coupon string
		fmt.Scanln(&coupon)
		if isValidCoupon(coupon) {
			total = applyCouponDiscount(total, coupon)
			fmt.Println("Coupon successfully applied!")
		} else {
			fmt.Println("I'm sorry, that coupon code is not valid.")
		}
	}

	fmt.Printf("That will be $%.2f", total)
}
