package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/docopt/docopt-go"
	_ "github.com/mattn/go-sqlite3"
)

const SqliteTimeLayout = "2006-01-02 15:04:05"

type MoldyFridge struct {
	DbName string
	Db     *sql.DB
}

type Food struct {
	Name   string
	Amount int
	Added  *time.Time
}

func (f *Food) toString() string {
	return fmt.Sprintf("%s: %d, added on %s", f.Name, f.Amount, f.FormatDateTime())
}

func (f *Food) FormatDateTime() string {
	return f.Added.Format(SqliteTimeLayout)
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
	query := "create table food (id integer not null primary key, name text, amount integer, added datetime);"
	_, err := f.Db.Exec(query)
	if err != nil {
		log.Print("can't initialize table")
		log.Fatal(err)
	}
}

func (f *MoldyFridge) checkDb() bool {
	_, err := f.Db.Exec("select id from food limit 1;")
	return err != nil
}

func (f *MoldyFridge) destroy() {
	f.Db.Close()
	os.Remove(f.DbName)
}

func (f *MoldyFridge) AddFood(name string, amount int) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now)
	query := fmt.Sprintf("insert into food (name, amount, added) values ('%s', %d, '%s');", name, amount, now)
	_, err := f.Db.Exec(query)
	if err != nil {
		log.Printf("Can't add %s\n", name)
		log.Fatal(err)
	}
}

func (f *MoldyFridge) RemoveFood(name string) {
	query := "delete from food where name = '" + name + "';"
	_, err := f.Db.Exec(query)
	if err != nil {
		log.Printf("Can't remove %s\n", name)
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

func (f *MoldyFridge) GetFood() []Food {
	var result []Food
	sql := "select name, amount, added from food;"
	rows, err := f.Db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var amount int
		var added time.Time
		rows.Scan(&name, &amount, &added)
		result = append(result, Food{name, amount, &added})
	}
	return result
}

func (f *MoldyFridge) SearchFood(query string) []Food {
	var result []Food
	sql := "select name, amount, added from food where lower(name) like '%" + query + "%';"
	rows, err := f.Db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var amount int
		var added_str string
		rows.Scan(&name, &amount, &added_str)
		added, _ := time.Parse(SqliteTimeLayout, added_str)
		result = append(result, Food{name, amount, &added})
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
	if fridge.checkDb() {
		fridge.init()
	}
	if arguments["add"] == true {
		for _, food := range arguments["<food>"].([]string) {
			fridge.AddFood(food, 1)
		}
	} else if arguments["rm"] == true {
		fridge.RemoveFood(arguments["<food>"].([]string)[0])
	} else if arguments["ls"] == true {
		for _, food := range fridge.GetFood() {
			fmt.Println(food.toString())
		}
	} else if arguments["search"] == true {
		for _, food := range fridge.SearchFood(arguments["<food>"].([]string)[0]) {
			fmt.Println(food.toString())
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
