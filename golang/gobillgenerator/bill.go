package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format
func (b *bill) format() string {

	fs := "Bill breakdown: \n"
	var total float64 = 0

	//items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v....$%v  \n", k+":", v)
		total += v
	}

	//tip

	fs += fmt.Sprintf("%-25v....$%v  \n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("-----------------------------------\n%-25v....$%0.2f", "total:", total+b.tip)

	return fs
}

// add tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item
func (b *bill) addItem(item string, value float64) {
	b.items[item] = value
}

//create bill with user input

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Println("Enter name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Enter Bill name",reader)

	b := newBill(name)
	fmt.Println("Created bill:", name)
	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}