package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func view(args ...string) {
	fmt.Print("Printing out database:\n\n")
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
	fmt.Print("\n\n...print complete\nReturning to the main menu...\n\n")
	main()
}

func add(args ...string) {
	var userAmount int
	var contCheck int
	contCheck = 1
	database, _ :=
		sql.Open("sqlite3", "./names.db")
	statement, _ :=
		database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	for contCheck == 1 {
		fmt.Print("How many entries would you like to make?\n")
		fmt.Scanf("%d", &userAmount)
		for counter := 0; counter < userAmount; counter++ {
			fmt.Print("Enter a name in the format: Firstname, Lastname: \n")
			reader := bufio.NewReader(os.Stdin)
			rawInput, _ := reader.ReadString('\n')
			trimmed := strings.Trim(rawInput, "\n")
			slicedInput := strings.Split(trimmed, ", ")
			statement, _ =
				database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
			statement.Exec(slicedInput[0], slicedInput[1])
		}
		fmt.Print("Would you like to add more entries?\n1) Yes\n2) No\n")
		fmt.Scanf("%d", &contCheck)
		fmt.Print("\n")
	}
	fmt.Print("\n\nReturning to main menu...\n\n")
	main()
}

func csvImport(args ...string) {
	var path string
	fmt.Print("Please put the raw path of the CSV\n")
	fmt.Scanf("%s", &path)
	fmt.Print("Now opening: ", path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Print("Starting import...\n\n")
	database, _ :=
		sql.Open("sqlite3", "./names.db")
	statement, _ :=
		database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		statement, _ =
			database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
		statement.Exec(record[0], record[1])
	}
	fmt.Print("\n\nImport complete, returning to main menu...\n\n")
	main()
}

func main() {
	var i int
	fmt.Print("Welcome to the database manager\nPlease select an option:\n\n1) View database\n2) Add entry to database\n3) Import CSV\n4) Exit\n\n")
	fmt.Scanf("%d", &i)
	fmt.Print("\n")
	switch i {
	case 1:
		view()
	case 2:
		add()
	case 3:
		csvImport()
	case 4:
		fmt.Print("Thank you, goodbye!")
		os.Exit(0)
	}
}
