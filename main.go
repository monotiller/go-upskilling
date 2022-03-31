package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dixonwille/wmenu"
	_ "github.com/mattn/go-sqlite3"
)

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

	return
}

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
	return
}

type userInput struct {
	option wmenu.Opt
}

func (u *userInput) optFunc(option wmenu.Opt) error {
	u.option = option
	return nil
}

func createMenu(p string, m []string, u *userInput) {
	menu := wmenu.NewMenu(p)
	menu.ChangeReaderWriter(os.Stdin, os.Stdout, os.Stderr)
	for i, m := range m {
		menu.Option(m, i, false, u.optFunc)

	}
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	prompt := "Please select an option"
	menuitems := []string{"View entries in database", "Add entry to database"}
	u := &userInput{}
	createMenu(prompt, menuitems, u)
	switch u.option.ID {
	case 0:
		view()
	case 1:
		add()
	}
}
