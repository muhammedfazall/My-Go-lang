package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	name, _ := getInput("Enter Bill name \n", reader)

	b := newBill(name)
	fmt.Println("Created bill:", name)
	return b
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose options a to add item, s to save, t to add tip : ", reader)

	switch opt {
	case "a":
		itemName, _ := getInput("item name: ", reader)
		price, _ := getInput("price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(itemName, p)

		fmt.Println("Item added!-", itemName, p)
		promptOptions(b)

	case "t":
		tip, _ := getInput("enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)


		fmt.Println("Tip added!-", t)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("Saved the bill", b.name)

	default:
		fmt.Println("Not a valid option...")
		promptOptions(b)
	}

}

func (b *bill) save()  {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt",data,0644)
	if err != nil{
		panic(err)
	}
	fmt.Println("Saved the bill to file-")
}