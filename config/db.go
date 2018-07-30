package config

import (
	"database/sql"
	"fmt"
	"github.com/go_webapp/helper"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {

	var err error

	c, err := LoadConfiguration("config/config.xxjson")
	if err != nil {
		panic(err)
	}

	ls := []string{"postgres://", c.Database.Host, ":", c.Database.Password, "@localhost/collections?sslmode=disable"}
	login := helper.Concat(ls)

	DB, err = sql.Open("postgres", login)

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You are connected to your database.")

}
