package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DataPerson struct {
	PersonName string
	CarName    string
}

type Person struct {
	ID   int
	Name string
}

type Car struct {
	ID       int
	PersonID int
	Name     string
}

func main() {
	getPeople()
	getCars()
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/trydatabase")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("Success connect to database")
	return db, nil
}

func getPeople() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var people []Person

	for rows.Next() {
		person := Person{}
		err = rows.Scan(&person.ID, &person.Name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	
		people = append(people, person)
	}
	for _, v := range people {
		fmt.Println(v)
	}
}

func getCars() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT people.name AS person_name, cars.name AS car_name FROM people INNER JOIN cars ON people.id = cars.person_id")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	dataPeople := []DataPerson{}
	for rows.Next() {
		dataPerson := DataPerson{}
		err = rows.Scan(&dataPerson.PersonName, &dataPerson.CarName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dataPeople = append(dataPeople, dataPerson)
	}

	for _, v := range dataPeople {
		fmt.Println(v)
	}
}