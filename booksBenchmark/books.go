package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	newBookCost = 6.0
	oldBookCost = 3.5
	couponIndex = 0
)

var (
	coupons = []string{
		"free-stuff",
		"half-off",
	}
	discounts = []float64{
		0.0,
		0.5,
	}
	hasCoupon string
	newBooks  int
	oldBooks  int
)

type CouponInfo struct {
	Coupon   string  `json:"coupon,omitempty"`
	Discount float64 `json:"discount,omitempty"`
}
type CouponDb struct {
	Coupons []CouponInfo
}

func computeCost(new, old float64) float64 {
	return new*newBookCost + old*oldBookCost
}

func index(item string, list []string) int {
	itemIndex := 0
	for i, c := range list {
		if c == item {
			itemIndex = i
		}
	}
	return itemIndex
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isValidCoupon(coupon string) bool {
	return stringInSlice(coupon, coupons)

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

	for {
		fmt.Println("How many new books are you buying today?")
		fmt.Scanln(&newBooks)
		if newBooks < 0 {
			fmt.Println("Please select a value of 0 or greater.")
		} else {
			break
		}
	}
	for {
		fmt.Println("How many old books are you buying today?")
		fmt.Scanln(&oldBooks)
		if oldBooks < 0 {
			fmt.Println("Please select a value of 0 or greater.")
		} else {
			break
		}
	}

	total := computeCost(float64(newBooks), float64(oldBooks))
	for {
		fmt.Println("Do you have a coupon? [Y/N]: ")
		fmt.Scanln(&hasCoupon)
		hasCoupon = strings.ToLower(hasCoupon)
		if hasCoupon != "y" && hasCoupon != "n" {
			fmt.Println("Please provide valid input.")
		} else {
			break
		}
	}

	f, err := os.Open("testingFolder\\coupons.json")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	db := CouponDb{}
	dec.Decode(&db)

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
