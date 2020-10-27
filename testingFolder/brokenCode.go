package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type CouponInfo struct {
	Coupon   string  `json:"coupon,omitempty"`
	Discount float64 `json:"discount,omitempty"`
}
type CouponDb struct {
	Coupons []CouponInfo
}

func main() {
	// coupons := []CouponInfo{
	// 	{Coupon: "free-stuff", Discount: 0.0},
	// 	{Coupon: "half-off", Discount: 0.5}}

	// db := CouponDb{Coupons: coupons}

	// var buf = new(bytes.Buffer)
	// enc := json.NewEncoder(buf)
	// enc.Encode(db)
	// f, err := os.Create("testingFolder\\coupons.json")
	// if nil != err {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()
	// io.Copy(f, buf)
	f, err := os.Open("testingFolder\\coupons.json")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	db := CouponDb{}
	dec.Decode(&db)
	fmt.Println(db)
}
