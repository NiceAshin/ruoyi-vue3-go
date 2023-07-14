package main

import (
	"database/sql"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	db, err := sql.Open("mysql", "root:rootroot@tcp(192.168.16.81:3306)/gy_istsp")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to database!")
}
