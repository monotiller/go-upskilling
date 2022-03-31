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

func view(args ...string) {
	database, _ :=
		sql.Open("sqlite3", "./names.db")
	rows, _ :=
		database.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
	main()
}

func add(args ...string) {
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

	main()
}

func main() {
	fmt.Print("Welcome to the database manager\nPlease select an option:\n\n1) View database\n2) Add entry to database\n3) Exit\n\n")
	var i int
	fmt.Scanf("%d", &i)
	switch i {
	case 1:
		view()
	case 2:
		add()
	case 3:
		fmt.Print("Thank you, goodbye!")
		os.Exit(0)
	}
}
