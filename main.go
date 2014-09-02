package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/docopt/docopt-go"
	_ "github.com/mattn/go-sqlite3"
)

type MoldyFridge struct {
	DbName string
	Db     *sql.DB
}

func NewFridge(dbname string) *MoldyFridge {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal("Can't create new fridge")
		log.Fatal(err)
	}
	return &MoldyFridge{DbName: dbname, Db: db}
}

func (f *MoldyFridge) init() {
	query := "create table food (id integer not null primary key, name text, amount integer);"
	_, err := f.Db.Exec(query)
	if err != nil {
		log.Print("can't initialize table")
		log.Fatal(err)
	}
}

func (f *MoldyFridge) destroy() {
	f.Db.Close()
	os.Remove(f.DbName)
}

func (f *MoldyFridge) AddFood(name string, amount int) {
	query := fmt.Sprintf("insert into food (name, amount) values ('%s', %d);", name, amount)
	log.Printf(query)
	_, err := f.Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func (f *MoldyFridge) PromptFood() {
	var name string
	var amount int
	fmt.Print("Enter food name to add: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter amount: ")
	fmt.Scanf("%d", &amount)
	f.AddFood(name, amount)
}

func (f *MoldyFridge) GetFood() map[string]int {
	result := make(map[string]int)
	sql := "select name, amount from food;"
	rows, err := f.Db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var amount int
		rows.Scan(&name, &amount)
		result[name] = amount
	}
	return result
}

func (f *MoldyFridge) SearchFood(query string) map[string]int {
	result := make(map[string]int)
	sql := "select name, amount from food where lower(name) like '%" + query + "%';"
	rows, err := f.Db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var amount int
		rows.Scan(&name, &amount)
		result[name] = amount
	}
	return result
}

func main() {
	usage := `moldyfridge.

Usage:
  moldyfridge (add | rm) <food>...
  moldyfridge ls
  moldyfridge search <food>
  moldyfridge reset
  
Options:
  -h --help     Show this screen.
  --version     Show version.`
	arguments, _ := docopt.Parse(usage, nil, true, "moldyfridge 0.1", false)
	fridge := NewFridge("test.db")
	if arguments["add"] == true {
		for _, food := range arguments["<food>"].([]string) {
			fridge.AddFood(food, 1)
		}
	} else if arguments["rm"] == true {
		// TODO Add function to remove food
	} else if arguments["ls"] == true {
		for food, amount := range fridge.GetFood() {
			fmt.Printf("%s: %d\n", food, amount)
		}
	} else if arguments["search"] == true {
		for food, amount := range fridge.SearchFood(arguments["<food>"].([]string)[0]) {
			fmt.Printf("%s: %d\n", food, amount)
		}
	} else if arguments["reset"] == true {
		fmt.Print("You really wan't to reset moldyfridge? (y/n): ")
		var choice string
		fmt.Scanf("%s", &choice)
		switch choice[0] {
		case 'y', 'Y':
			fridge.destroy()
		case 'n', 'N':
			fmt.Println("ok, exiting")
		default:
			fmt.Println("what?")
		}
	}
}
