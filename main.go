package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ :=
		sql.Open("sqlite3", "./names.db")
	statement, _ :=
		database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()

	fmt.Print("Enter a name in the format: Firstname, Lastname: \n")
	reader := bufio.NewReader(os.Stdin)
	rawInput, _ := reader.ReadString('\n')
	slicedInput := strings.Split(rawInput, ", ")

	statement, _ =
		database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec(slicedInput[0], slicedInput[1])
	rows, _ :=
		database.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}
