package main

import (
	"fmt"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

/**
* @param name string
* @return b bill
 */
func newUserBill(name string) bill {
	data := bill{
		name:  "John Doe",
		items: map[string]float64{"Sneakers": 26.34, "Jeans": 82.24, "T-Shirt": 42.90},
		tip:   0,
	}
	return data
}

//Format bill
func (data bill) formatter() string {
	fs := data.name + "'s Bill Breakdown \n"
	var total float64 = 0
	// list each item
	for k, v := range data.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total: ", total)
	return fs
}
