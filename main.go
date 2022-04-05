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

	"github.com/joho/sqltocsv"
	_ "github.com/mattn/go-sqlite3"
)

func view(args ...string) {
	fmt.Println("Printing out database:")
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
	fmt.Println("...print complete\nReturning to the main menu...")
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
		fmt.Println("How many entries would you like to make?")
		fmt.Scanf("%d", &userAmount)
		for counter := 0; counter < userAmount; counter++ {
			fmt.Println("Enter a name in the format: Firstname, Lastname: ")
			reader := bufio.NewReader(os.Stdin)
			rawInput, _ := reader.ReadString('\n')
			trimmed := strings.Trim(rawInput, "\n")
			slicedInput := strings.Split(trimmed, ", ")
			statement, _ =
				database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
			statement.Exec(slicedInput[0], slicedInput[1])
		}
		fmt.Println("Would you like to add more entries?\n1) Yes\n2) No")
		fmt.Scanf("%d", &contCheck)
	}
	fmt.Println("Returning to main menu...")
	main()
}

func csvImport(args ...string) {
	var path string
	fmt.Println("Please put the raw path of the CSV")
	fmt.Scanf("%s", &path)
	fmt.Println("Now opening: ", path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Starting import...")
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
	fmt.Println("Import complete, returning to main menu...")
	main()
}

func csvExport(args ...string) {
	var path string
	fmt.Println("Please specify the raw path of where you would like to export the database to including file name.\nExample: /Documents/Export.csv")
	fmt.Scanf("%s", &path)
	fmt.Printf("Now exporting database to the file: %s\n\n", path)
	database, _ :=
		sql.Open("sqlite3", "./names.db")
	rows, _ :=
		database.Query("SELECT firstname, lastname FROM people")
	err := sqltocsv.WriteFile(path, rows)
	if err != nil {
		panic(err)
	}
	fmt.Println("Export complete, returning you to the main menu")
	main()
}

func main() {
	var i int
	fmt.Println("Welcome to the database manager\nPlease select an option:\n\n1) View database\n2) Add entry to database\n3) Import CSV\n4) Export database to a CSV\n5) Exit")
	fmt.Scanf("%d", &i)
	switch i {
	case 1:
		view()
	case 2:
		add()
	case 3:
		csvImport()
	case 4:
		csvExport()
	case 5:
		fmt.Println("Thank you, goodbye!")
		os.Exit(0)
	}
}
