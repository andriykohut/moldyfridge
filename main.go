package main

import (
	"fmt"
	"strconv"

	"github.com/andriykohut/moldyfridge/fridgedb"
	"github.com/docopt/docopt-go"
)

func main() {
	usage := `moldyfridge.

Usage:
  moldyfridge (add | rm) <food>...
  moldyfridge (add | rm) <food> [--amount <amount>]
  moldyfridge ls
  moldyfridge search <food>
  moldyfridge reset
  
Options:
  -h --help     Show this screen.
  --version     Show version.`
	arguments, _ := docopt.Parse(usage, nil, true, "moldyfridge 0.1", false)
	fridge := fridgedb.NewFridge("test.db")
	if fridge.CheckDb() {
		fridge.Init()
	}
	var amount int
	if arguments["<amount>"] != nil {
		amount, _ = strconv.Atoi(arguments["<amount>"].(string))
	} else {
		amount = 1
	}
	if arguments["add"] == true {
		for _, food := range arguments["<food>"].([]string) {
			fridge.AddFood(food, amount)
		}
	} else if arguments["rm"] == true {
		// TODO: Need to fix this shit
		if amount != 1 {
			fridge.RemoveFood(arguments["<food>"].([]string)[0], amount)
		} else {
			fridge.RemoveFood(arguments["<food>"].([]string)[0])
		}
	} else if arguments["ls"] == true {
		for _, food := range fridge.GetFood() {
			fmt.Println(food.ToString())
		}
	} else if arguments["search"] == true {
		for _, food := range fridge.SearchFood(arguments["<food>"].([]string)[0]) {
			fmt.Println(food.ToString())
		}
	} else if arguments["reset"] == true {
		fmt.Print("You really wan't to reset moldyfridge? (y/n): ")
		var choice string
		fmt.Scanf("%s", &choice)
		switch choice[0] {
		case 'y', 'Y':
			fridge.Destroy()
		case 'n', 'N':
			fmt.Println("ok, exiting")
		default:
			fmt.Println("what?")
		}
	}
}
