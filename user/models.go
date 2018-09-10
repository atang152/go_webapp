package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/atang152/go_webapp/config"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// To do. Query Database and review whether session expired
// Done: Added creation time of cookie

type User struct {
	Id        string
	Username  string
	Password  []byte
	Firstname string
	Lastname  string
	Cookie    *http.Cookie // Cookie is not stored in "users" database but inside "sessions" database
}

type Session struct {
	Id          string
	SessionUser string
	UserCookie  string
	timeCreated time.Time
}

// CREATE TABLE users (
//    Id     SERIAL PRIMARY KEY  NOT NULL,
//    username           TEXT    NOT NULL,
//    password           TEXT    NOT NULL,
//    firstname          TEXT    NOT NULL,
//    lastname           TEXT    NOT NULL
// );

// CREATE TABLE sessions(
//   id SERIAL PRIMARY KEY NOT NULL,
//   username TEXT NOT NULL,
// 	 cookie TEXT NOT NULL,
//   timeCreated TIME NOT NULL,
// );

func createSession() (*http.Cookie, time.Time) {

	const sessionLength int = 60
	var timeCreated time.Time

	// Create Cookie used in web session
	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	c.MaxAge = sessionLength

	// Store creation time of Cookie into Postgres Database
	// Note that time.Now() uses current local time configured to your computer
	timeCreated = time.Now()
	fmt.Println(timeCreated)

	return c, timeCreated
}

func CleanSessionDB() {
	// To Do: Use Arithmetic operation on time values in PostgresDB
	// Change minutes to a sessionLength
	_, err := config.DB.Exec("DELETE FROM sessions WHERE timeCreated < $1 - time '00:05';", time.Now())

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No results found")
		// return errors.New("No results found" + err.Error())
	case err != nil:
		fmt.Println(err.Error())
		// return errors.New("400. Bad request" + err.Error())

	}

	fmt.Println("Cookie sessions in DB cleaned")
	// return nil

	// rows, err := config.DB.Query("SELECT * FROM sessions")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// defer rows.Close()

	// sessions := make([]Session, 0)

	// for rows.Next() {
	// 	s := Session{}
	// 	err := rows.Scan(&s.Id, &s.SessionUser, &s.UserCookie, &s.timeCreated)
	// }
}

func DeleteSession(r *http.Request) error {

	// Retrieve session information
	c, err := r.Cookie("session")

	if c == nil {
		fmt.Println("No cookies found in session")
		return errors.New("No cookies found in session" + err.Error())
	}

	if err != nil {
		return errors.New("500. Internal Server Error." + err.Error())
	}

	fmt.Println(c.Value)

	_, err = config.DB.Exec("DELETE FROM sessions WHERE cookie=$1;", c.Value)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("Cookie not found")
		return errors.New("No results found" + err.Error())
	case err != nil:
		fmt.Println(err.Error())
		return errors.New("400. Bad request" + err.Error())

	}

	// Clean All Historical sessions
	CleanSessionDB()

	fmt.Println("Cookie deleted")
	return nil
}

func AlreadyLoggedIn(r *http.Request) bool {

	// Retrieve session information
	c, err := r.Cookie("session")

	if c == nil {
		errors.New("No cookies found in session" + err.Error())
		fmt.Println("No cookies found in session")
		return false
	}

	if err != nil {
		errors.New("500. Internal Server Error." + err.Error())
		return false
	}

	fmt.Println(c.Value)

	// Query Database for session information
	row := config.DB.QueryRow("SELECT id, username, cookie FROM sessions where cookie = $1", c.Value)

	s := Session{}
	err = row.Scan(&s.Id, &s.SessionUser, &s.UserCookie)

	switch {
	case err == sql.ErrNoRows:
		errors.New("Cookie not found" + err.Error())
		fmt.Println("No cookie monster")
		return false
	case err != nil:
		errors.New("Cookie not found" + err.Error())
		fmt.Println(err.Error())
		return false
	}

	// Query user database wtih cookie information
	u := User{}
	row = config.DB.QueryRow("SELECT * FROM users WHERE username = $1", s.SessionUser)
	err = row.Scan(&u.Id, &u.Username, &u.Password, &u.Firstname, &u.Lastname)

	switch {
	case err == sql.ErrNoRows:
		errors.New("Username associated with cookie is not found in user database" + err.Error())
		fmt.Println("Username associated with cookie is not found in user database")
		return false
	case err != nil:
		errors.New("Username associated with cookie is not found in user database" + err.Error())
		fmt.Println(err.Error())
		return false
	}

	fmt.Println("User already logged in")
	return true
}

func GetUser(r *http.Request) (User, error) {

	var timeCreated time.Time

	// Process form submission
	u := User{}
	u.Username = r.FormValue("username")
	p := r.FormValue("password")

	// Validate Form Value
	if u.Username == "" {
		return u, errors.New("400. Bad request. All fields must be completed.")
	}

	// Query Database for user information
	row := config.DB.QueryRow("SELECT * FROM users WHERE username = $1", u.Username)
	err := row.Scan(&u.Id, &u.Username, &u.Password, &u.Firstname, &u.Lastname)

	// Username input is not found in database
	if err != nil {
		return u, errors.New("Username and/or password do not match.")
	}

	// Verify whether password match the stored password
	err = bcrypt.CompareHashAndPassword(u.Password, []byte(p))
	if err != nil {
		return u, errors.New("Username and/or password do not match.")
	}

	// Create session
	u.Cookie, timeCreated = createSession()
	// Insert username-cookie session into Database
	_, err = config.DB.Exec("INSERT INTO sessions (username, cookie, timecreated) VALUES($1, $2, $3)", u.Username, u.Cookie.Value, timeCreated)
	if err != nil {
		return u, errors.New("500. Internal Server Error." + err.Error())
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

	// Check if the user name is already taken from Database.
	row := config.DB.QueryRow("SELECT * FROM users WHERE username = $1", u.Username)
	err := row.Scan(&u.Id, &u.Username, &u.Password, &u.Firstname, &u.Lastname)

	if err != nil {
		errors.New("500. Internal Server Error." + err.Error())
		return false
	}

	return true
}

func InsertUser(r *http.Request) (User, error) {

	var timeCreated time.Time

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

		u.Cookie, timeCreated = createSession()

		// Insert username-cookie session into Database
		_, err := config.DB.Exec("INSERT INTO sessions (username, cookie, timecreated) VALUES($1, $2, $3)", u.Username, u.Cookie.Value, timeCreated)
		if err != nil {
			return u, errors.New("500. Internal Server Error." + err.Error())
		}

		// If user name is not taken then encrypt Form Password Value
		encrypt_p, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			return u, errors.New("500. Internal Server Error." + err.Error())
		}
		u.Password = encrypt_p

		// Insert user info into Database
		_, err = config.DB.Exec("INSERT INTO users (username, password, firstname, lastname) VALUES($1, $2, $3, $4)", u.Username, u.Password, u.Firstname, u.Lastname)
		if err != nil {
			return u, errors.New("500. Internal Server Error." + err.Error())
		}

	} else {
		return u, errors.New("403. Status Forbidden. Username already taken.")
	}

	return u, nil
}

// To do should seperate Account and Profile

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
