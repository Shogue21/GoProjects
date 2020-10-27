package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	db []CouponInfo

	// coupons and discount are the original way of getting coupon information that I'm trying to change.
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

//CouponInfo is...
type CouponInfo struct {
	Coupon   string  `json:"coupon,omitempty"`
	Discount float64 `json:"discount,omitempty"`
}

// Adds the cost of new and old books being purchased together.
func computeCost(new, old float64) float64 {
	return new*newBookCost + old*oldBookCost
}

//Indexing tool used in applyCouponDiscount
func index(item string, array []string) int {
	itemIndex := 0
	for i, c := range array {
		if c == item {
			itemIndex = i
		}
	}
	return itemIndex
}

// Checks if string (a) is in provided array.
func stringInSlice(a string, array []string) bool {
	for _, b := range array {
		if b == a {
			return true
		}
	}
	return false
}

// Uses stringInSlice to check if the provided coupon is valid.
func isValidCoupon(coupon string) bool {
	return stringInSlice(coupon, coupons)

}

// Applies the discount associated with the coupon provided, if it is a valid coupon.
func applyCouponDiscount(cost float64, coupon string) float64 {
	couponIndex := index(coupon, coupons)
	if isValidCoupon(coupon) {
		cost = cost * discounts[couponIndex]
	}
	return cost
}

func main() {
	// Importing json file and getting its' contents.
	f, err := os.Open("booksBenchmark\\coupons.json")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	var db []CouponInfo
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(bytes, &db)
	fmt.Println(db)

	// Main Program starts here.
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
