package main

import (
	"fmt"
	"log"
	"math"

	"github.com/shopspring/decimal"
)

var mealCost float64
var numGuests int64

// TipInterface creates the prompt for the Tip function
func TipInterface() {
	fmt.Println("")
	fmt.Println("*** SPLIT the TIP ***")
	fmt.Println("Enter the Total and Number of Guests")
	fmt.Println("Enter the Total Meal Amount")
	fmt.Scanf("%f", &mealCost)
	fmt.Println("Enter the Number of Guests")
	fmt.Scanf("%d", &numGuests)

	tips, err := SplitTip(mealCost, numGuests)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Tips: for each guest", tips)
}

// SplitTip returns a list of values for each tip
func SplitTip(mealCost float64, numGuests int64) (splits []float64, err error) {
	if ok := CheckDecimal(mealCost); !ok {
		return nil, fmt.Errorf("Invalid number of decimal places")
	}
	tip := decimal.NewFromFloat(1.1500)
	mealCostDec := decimal.NewFromFloat(mealCost)
	decimalCost := mealCostDec.Mul(tip).Round(2)
	decNumGuests := decimal.NewFromFloat(float64(numGuests))

	tempDec := decimalCost.DivRound(decNumGuests, 2)
	if tempDec.Mul(decimal.New(numGuests, 0)).Equal(decimalCost) {
		for i := 0; i < int(numGuests); i++ {
			floatVal, _ := tempDec.Float64()
			splits = append(splits, floatVal)
		}

		return splits, err
	}
	for {
		amt := decimalCost.DivRound(decimal.NewFromFloat(float64(numGuests)), 2)
		fv, _ := amt.Float64()
		splits = append(splits, fv)
		decimalCost = decimalCost.Sub(amt)
		numGuests--
		if numGuests == 0 {
			break
		}
	}
	return splits, err
}

// CheckDecimal returns an error if float receieved is greater than 2 decimal
// spaces
func CheckDecimal(num float64) bool {
	val := num * math.Pow(10.0, 2)
	check := val - float64(int(val))
	return check == 0
}
