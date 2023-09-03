package main

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Animal struct {
	ID         int
	AnimalName string
}

func init() {
	// Open a database connection
	var err error
	db, err = sql.Open("mysql", "Mark123:123456@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err.Error())
	}

	// Check if the database connection is successful
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

func GetAnimals() ([]*Animal, error) {
	// Query the database to retrieve data
	rows, err := db.Query("SELECT ID, AnimalName FROM AnimalTable")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice of pointers to Animal
	animals := make([]*Animal, 0)

	// Iterate through the result set and scan data into pointers
	for rows.Next() {
		animal := new(Animal)
		err := rows.Scan(&animal.ID, &animal.AnimalName)
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}

	return animals, err
}

func GetAnimal(id int) (*Animal, error) {

	// mysql -- version
	animal := Animal{}
	query := "SELECT ID, AnimalName FROM AnimalTable WHERE ID=? "
	err := db.QueryRow(query, id).Scan(&animal.ID, &animal.AnimalName)

	if err != nil {
		return nil, err
	}
	return &animal, nil

	/* sql server -- version
	query := "select id, name from AnimalTable where id=@id"
	row := db.QueryRow(query, sql.Named("ID", id))
	*/
}

func main() {
	/*
		animals, err := GetAnimals()
		if err != nil {
			panic(err)
		}
	*/
	// Now, you have a slice of pointers to Animal objects that represent the rows in the table
	// You can work with these objects as needed
	/*
		for _, a := range animals {
			fmt.Printf("ID %d : %s\n", a.ID, a.AnimalName)
		}
	*/

	/* Find Animal Name by ID
	animal, err := GetAnimal(5)
	if err != nil {
		panic(err)
	}
	fmt.Println(*animal)
	*/
	/*
		animal := Animal{12, "Giraffe"}
		err := AddAnimal(animal)
		if err != nil {
			panic(err)
		}
	*/

	/*
		animal := Animal{12, "Monkey king"}
		err := Animal(animal)
		if err != nil {
			panic(err)
		}
	*/
	/*
		err := DeleteAnimal(12)
		if err != nil {
			panic(err)
		}
	*/

}

func AddAnimal(animal Animal) error {
	tx, err := db.Begin()

	query := "INSERT INTO AnimalTable (ID, AnimalName) VALUES (?, ?)"
	result, err := tx.Exec(query, animal.ID, animal.AnimalName)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affected <= 0 {
		return errors.New("cannot insert")
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func UpdateAnimal(animal Animal) error {
	query := "UPDATE AnimalTable SET AnimalName = ? WHERE ID = ?"
	result, err := db.Exec(query, animal.AnimalName, animal.ID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("cannot insert")
	}

	return nil
}

func DeleteAnimal(id int) error {
	query := "DELETE FROM AnimalTable WHERE ID = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("cannot delete")
	}

	return nil
}
