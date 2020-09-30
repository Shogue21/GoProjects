package main

import (
	"fmt"
	"strings"
)

const (
	newBookCost = 6.0
	oldBookCost = 3.5
	couponIndex = 0
)

var coupons = []string{
	"free-stuff",
	"half-off",
}
var discounts = []float64{
	0.0,
	0.5,
}

func computeCost(new, old float64) float64 {
	return new*newBookCost + old*oldBookCost
}

func index(item string, list []string) int {
	itemIndex := 0
	for i, c := range list {
		if c == item {
			itemIndex = i
			fmt.Println("made it!")
		}
	}
	return itemIndex
}
func isValidCoupon(coupon string) bool {
	return bool(coupon == "free-stuff" || coupon == "half-off")
}

func applyCouponDiscount(cost float64, coupon string) float64 {
	couponIndex := index(coupon, coupons)
	if isValidCoupon(coupon) {
		cost = cost * discounts[couponIndex]
	}
	return cost
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
