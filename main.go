//gorm

package main

import (
	"connectingDBtest/database"
	"connectingDBtest/models"
	"connectingDBtest/settings"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"log"
)


//var dsn = flag.String("dsn", "postgres://postgres:123456789@qw@localhost:5432/postgres", "Postgres DSN")

func main() {
	fmt.Println("Start server...")
	//Unmarshal / marshal --> Parse from settings.json to model SettingDB
	bytes, err := ioutil.ReadFile("settings/settings.json")
	if err != nil {
		log.Fatalf("Error can't read from file %e",err)
	}
	var DataObject settings.SettingsDB
	err = json.Unmarshal(bytes,&DataObject)
	if err != nil {
		log.Fatalf("Error can't parse bytes %e", err)
	}
	fmt.Println(string(bytes))
	// urlExample := "postgres://username:password@localhost:5432/database_name
	urlDatabase := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", DataObject.Username, DataObject.Password, DataObject.DatabaseName)
	connect, err := pgx.Connect(context.Background(), urlDatabase)
	if err != nil {
		log.Fatalf("can't connect to db %e", err)
	}
	database.DbInitialization(connect)
	user, err := models.GetUsers(connect)
	if err != nil {
		log.Println("Bad")
	} else {
		fmt.Println(user)
	}
/*	_, err = connect.Exec(context.Background(), `create table if not exists juniors (
id serial primary key,
name varchar(30) not null);`)
	if err != nil {
		fmt.Println("can't create a table", err)
	}*/
//QueryRow
/*	type juniors struct {
		id int `json:"id"`
		name string `json:"name"`
	}

	//name := "RustamR"
	var myJun juniors
	err = connect.QueryRow(context.Background(), `select * from juniors where name = 'RustamR'`).Scan(&myJun.id, &myJun.name)
	if err != nil {
		fmt.Println("ploxo", err)
	} else {
		fmt.Println("good")
	}
	fmt.Println(myJun)*/
//Query
/*	results, err := connect.Query(context.Background(),"SELECT id, name FROM juniors")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&myJun.id, &myJun.name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(myJun.name)
	}*/
}
