// Package fridgedb provides toplevel operations on database
package fridgedb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/andriykohut/moldyfridge/models"
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

func (f *MoldyFridge) Init() {
	query := "create table food (id integer not null primary key, name text, amount integer, added integer);"
	_, err := f.Db.Exec(query)
	if err != nil {
		log.Print("can't initialize table")
		log.Fatal(err)
	}
}

func (f *MoldyFridge) CheckDb() bool {
	_, err := f.Db.Exec("select id from food limit 1;")
	return err != nil
}

func (f *MoldyFridge) Destroy() {
	f.Db.Close()
	os.Remove(f.DbName)
}

func (f *MoldyFridge) AddFood(name string, amount int) {
	now := time.Now().Unix()
	query := fmt.Sprintf("insert into food (name, amount, added) values ('%s', %d, %d);", name, amount, now)
	_, err := f.Db.Exec(query)
	if err != nil {
		log.Printf("Can't add %s\n", name)
		log.Fatal(err)
	}
}

func (f *MoldyFridge) RemoveFood(args ...interface{}) {
	name := args[0].(string)
	amount := 0
	var query string
	if len(args) == 2 {
		amount = args[1].(int)
	}
	if amount == 0 {
		query = "delete from food where name = '" + name + "';"
	} else {
		query = fmt.Sprintf("update food set amount = amount - %d where name = '%s';", amount, name)
	}
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

func (f *MoldyFridge) GetFood() []models.Food {
	var result []models.Food
	sql := "select name, amount, added from food order by added;"
	rows, err := f.Db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var amount int
		var added int64
		rows.Scan(&name, &amount, &added)
		result = append(result, models.Food{name, amount, added})
	}
	return result
}

func (f *MoldyFridge) SearchFood(query string) []models.Food {
	var result []models.Food
	sql := "select name, amount, added from food where lower(name) like '%" + query + "%' order by added;"
	rows, err := f.Db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var amount int
		var added int64
		rows.Scan(&name, &amount, &added)
		result = append(result, models.Food{name, amount, added})
	}
	return result
}
