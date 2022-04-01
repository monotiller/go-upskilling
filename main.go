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
	var i int
	var j int
	j = 1
	database, _ :=
		sql.Open("sqlite3", "./names.db")
	statement, _ :=
		database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	for j == 1 {
		fmt.Print("How many entries would you like to make?\n")
		fmt.Scanf("%d", &i)
		for counter := 0; counter < i; counter++ {
			fmt.Print("Enter a name in the format: Firstname, Lastname: \n")
			reader := bufio.NewReader(os.Stdin)
			rawInput, _ := reader.ReadString('\n')
			slicedInput := strings.Split(rawInput, ", ")
			statement, _ =
				database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
			statement.Exec(slicedInput[0], slicedInput[1])
		}
		fmt.Print("Would you like to add more entries?\n1) Yes\n2) No\n")
		fmt.Scanf("%d", &j)
		fmt.Printf("\n")
	}
	main()
}

func main() {
	var i int
	fmt.Print("Welcome to the database manager\nPlease select an option:\n\n1) View database\n2) Add entry to database\n3) Exit\n\n")
	fmt.Scanf("%d", &i)
	fmt.Print("\n")
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
