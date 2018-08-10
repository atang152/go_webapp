package user

import (
	"errors"
	"github.com/atang152/go_webapp/config"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	Id        string
	Username  string
	Password  []byte
	Firstname string
	Lastname  string
}

// CREATE TABLE users (
//    Id     SERIAL PRIMARY KEY  NOT NULL,
//    username           TEXT    NOT NULL,
//    password           TEXT    NOT NULL,
//    firstname          TEXT    NOT NULL,
//    lastname           TEXT    NOT NULL
// );

func GetUser(r *http.Request) (User, error) {
	u := User{}
	u.Username = r.FormValue("username")
	// Validate Form Value
	if u.Username == "" {
		return u, errors.New("400. Bad request. All fields must be complete.")
	}

	row := config.DB.QueryRow("SELECT * FROM users WHERE username = $1", u.Username)
	err := row.Scan(&u.Id, &u.Username, &u.Password, &u.Firstname, &u.Lastname)

	if err != nil {
		return u, err
	}

	return u, err
}

func userNameTaken(r *http.Request) bool {
	u := User{}
	u.Username = r.FormValue("username")

	// Validate Form Value
	if u.Username == "" {
		errors.New("400. Bad request. Username must be filled.")
	}

	// Check if the user name is taken from Database.
	row := config.DB.QueryRow("SELECT * FROM users WHERE username = $1", u.Username)
	err := row.Scan(&u.Id, &u.Username, &u.Password, &u.Firstname, &u.Lastname)

	if err != nil {
		return false
	}

	return true
}

func InsertUser(r *http.Request) (User, error) {

	// Get form values
	u := User{}
	u.Username = r.FormValue("username")
	p := r.FormValue("password")
	u.Firstname = r.FormValue("firstname")
	u.Lastname = r.FormValue("lastname")

	// Validate Form Value
	if u.Username == "" || u.Firstname == "" || u.Lastname == "" || p == "" {
		return u, errors.New("400. Bad request. All fields must be complete.")
	}

	// Check if user name is taken
	if !userNameTaken(r) {

		// If user name is not taken then encrypt Form Password Value
		encrypt_p, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			return u, errors.New("500. Internal Server Error." + err.Error())
		}
		u.Password = encrypt_p

		// Insert Values into Database
		_, err = config.DB.Exec("INSERT INTO users (username, password, firstname, lastname) VALUES($1, $2, $3, $4)", u.Username, u.Password, u.Firstname, u.Lastname)
		if err != nil {
			return u, errors.New("500. Internal Server Error." + err.Error())
		}
	} else {
		return u, errors.New("403. Status Forbidden. Username already taken.")
	}

	return u, nil
}

// var DbUsers = map[string]User{}
// var DbSessions = map[string]string{}

// type Account struct {
// 	Type     string `json: "type, omitempty"`
// 	Pid      string `json: "pid, omitempty"`
// 	Email    string `json: "email, omitempty"`
// 	Password string `json: "password, omitempty"`
// }

// type Profile struct {
// 	Type      string `json: "type, omitempty"`
// 	Firstname string `json: "firstname, omitempty"`
// 	Lastname  string `json: "lastname, omitempty"`
// }

// type Session struct {
// 	Type string `json: "type, omitempty"`
// 	Pid  string `json: "pid, omitempty"`
// }

// var dbAccount = map[string]Account{} //

// type user struct {
//     UserName string
//     Password []byte
//     First    string
//     Last     string
// }

// var dbUsers = map[string]user{}      // user ID, user
// var dbSessions = map[string]string{} // session ID, user ID
