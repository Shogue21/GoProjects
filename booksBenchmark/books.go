package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	newBookCost = 6.0
	oldBookCost = 3.5
	couponIndex = 0
)

var (
	db        []CouponInfo
	hasCoupon string
	newBooks  int
	oldBooks  int
	coupon    string
	total     float64
	purchase  PurchaseInfo
)

//CouponInfo is...
type CouponInfo struct {
	Coupon   string  `json:"coupon,omitempty"`
	Discount float64 `json:"discount,omitempty"`
}

//PurchaseInfo is...
type PurchaseInfo struct {
	NewBooks int     `json:"New"`
	OldBooks int     `json:"Old"`
	Coupon   string  `json:"Coupon"`
	Total    float64 `json:"Total"`
}

//EncodeAsStrings is...
func (user PurchaseInfo) EncodeAsStrings() (ss []string) {
	ss = make([]string, 4)
	ss[0] = strconv.Itoa(user.NewBooks)
	ss[1] = strconv.Itoa(user.OldBooks)
	ss[2] = user.Coupon
	ss[3] = strconv.FormatFloat(user.Total, 'f', 2, 64)
	return
}

// Adds the cost of new and old books being purchased together.
func computeCost(new, old float64) float64 {
	return new*newBookCost + old*oldBookCost
}

// Checks if the coupon is valid by looping over the coupon info in db.
func isValidCoupon(coupon string) bool {
	for _, ci := range db {
		if ci.Coupon == coupon {
			return true
		}
	}
	return false
}

// Applies the discount associated with the coupon provided, if it is a valid coupon.
func applyCouponDiscount(cost float64, coupon string) float64 {
	if isValidCoupon(coupon) {
		for _, ci := range db {
			if ci.Coupon == coupon {
				cost = cost * ci.Discount
			}
		}
	}
	return cost
}

func main() {

	// Importing json file and getting its' contents.
	f, err := os.Open("booksBenchmark/coupons.json")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(bytes, &db)
	// Main Program starts here.
	fmt.Println("Welcome to Bargain Books!")
	fmt.Printf("New books are $%.2f each.\n", newBookCost)
	fmt.Printf("Old books are $%.2f each.\n", oldBookCost)
	for {
		fmt.Println("How many new books are you buying today?")
		_, err := fmt.Scanln(&newBooks)
		if err != nil || newBooks < 0 {
			fmt.Println("Please select a value of 0 or greater.")
		} else {
			break
		}
	}
	for {
		fmt.Println("How many old books are you buying today?")
		_, err := fmt.Scanln(&oldBooks)
		if err != nil || oldBooks < 0 {
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
		fmt.Println("What is your coupon?")
		fmt.Scanln(&coupon)
		if isValidCoupon(coupon) {
			total = applyCouponDiscount(total, coupon)
			fmt.Println("Coupon successfully applied!")
		} else {
			fmt.Println("I'm sorry, that coupon code is not valid.")
		}
	}
	fmt.Printf("That will be $%.2f", total)

	purchase := []PurchaseInfo{
		{newBooks,
			oldBooks,
			coupon,
			total},
	}

	purchases, err := os.OpenFile("booksBenchmark/purchases.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer purchases.Close()

	writer := csv.NewWriter(purchases)
	for _, info := range purchase {
		ss := info.EncodeAsStrings()
		writer.Write(ss)
	}
	writer.Flush()

	err = writer.Error()
	if err != nil {
		log.Fatalln(err)
	}
}
