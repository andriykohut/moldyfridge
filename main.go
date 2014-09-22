package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strconv"

	"github.com/andriykohut/gotable"
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
	usr, _ := user.Current()
	home := usr.HomeDir
	confdir := path.Join(home, ".moldyfridge")
	dbpath := path.Join(confdir, "fridge.db")
	fridge := fridgedb.NewFridge(dbpath)
	if fridge.CheckDb() {
		os.Mkdir(confdir, 0755)
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
		var food_rows []map[string]string
		for _, food := range fridge.GetFood() {
			food_rows = append(food_rows, map[string]string{
				"name": food.Name, "age": food.StringAge(),
				"amount": strconv.Itoa(food.Amount),
			})
		}
		table := gotable.NewTable(food_rows, true, []string{"name", "age", "amount"})
		fmt.Println(table.GetTable())
	} else if arguments["search"] == true {
		var food_rows []map[string]string
		for _, food := range fridge.SearchFood(arguments["<food>"].([]string)[0]) {
			food_rows = append(food_rows, map[string]string{
				"name": food.Name, "age": food.StringAge(),
				"amount": strconv.Itoa(food.Amount),
			})
		}
		table := gotable.NewTable(food_rows, true, []string{"name", "age", "amount"})
		fmt.Println(table.GetTable())
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
