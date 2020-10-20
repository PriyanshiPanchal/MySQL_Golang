package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:priya@1298@tcp(localhost:3306)/games")
	if err != nil {
		// fmt.Print(err.Error())
		fmt.Println("Error creating DB:", err)
		fmt.Println("To verify, db is:", db)
	}
	defer db.Close()
	fmt.Println("Successfully  Connected to MYSQl")

	// insert, err:=db.Query("INSERT INTO user_games VALUES('Dev')")

	// if err!=nil{
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	// fmt.Println("Successfully inserted into user tables")

	results, err := db.Query("SELECT name FROM user_games")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}
}
