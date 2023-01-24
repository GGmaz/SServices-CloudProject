package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type student struct {
	Jmbg      string `json:"jmbg"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Index     string `json:"index"`
}

type professor struct {
	Jmbg      string `json:"jmbg"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	createTables()
	http.HandleFunc("/student", handleStudentRequest)
	http.HandleFunc("/professor", handleProfessorRequest)
	http.ListenAndServe(":8050", nil)
}

func handleStudentRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var student student
		json.NewDecoder(r.Body).Decode(&student)

		db, err := sql.Open("postgres", "host=db user=postgres password=ftn dbname=uns port=5432 sslmode=disable")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		row := db.QueryRow("SELECT 1 FROM students WHERE jmbg=$1", student.Jmbg)
		var exists int
		err = row.Scan(&exists)
		if err == sql.ErrNoRows {
			// Insert student data into database
			_, err := db.Exec("INSERT INTO students (jmbg, first_name, last_name, index) VALUES ($1, $2, $3, $4)", student.Jmbg, student.FirstName, student.LastName, student.Index)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func handleProfessorRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var professor professor
		json.NewDecoder(r.Body).Decode(&professor)

		db, err := sql.Open("postgres", "host=db user=postgres password=ftn dbname=uns port=5432 sslmode=disable")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		row := db.QueryRow("SELECT 1 FROM professors WHERE jmbg=$1", professor.Jmbg)
		var exists int
		err = row.Scan(&exists)
		if err == sql.ErrNoRows {
			// Insert professor data into database
			_, err := db.Exec("INSERT INTO professors (jmbg, first_name, last_name) VALUES ($1, $2, $3)", professor.Jmbg, professor.FirstName, professor.LastName)
			if err != nil {
				fmt.Println(err)
				return
			}
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func createTables() {
	db, err := sql.Open("postgres", "host=db user=postgres password=ftn dbname=uns port=5432 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	var exists int
	err = db.QueryRow("SELECT COUNT(*) FROM pg_tables WHERE tablename = $1", "students").Scan(&exists)
	if err != nil {
		log.Fatalf("Error while checking table students: %v", err)
	}
	if exists == 0 {
		_, err = db.Exec(`CREATE TABLE students (
			id SERIAL PRIMARY KEY,
			jmbg VARCHAR(13) NOT NULL,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			index VARCHAR(10) NOT NULL
		);`)
		if err != nil {
			log.Fatalf("Error while creating the students table: %v", err)
		}
	}

	err = db.QueryRow("SELECT COUNT(*) FROM pg_tables WHERE tablename = $1", "professors").Scan(&exists)
	if err != nil {
		log.Fatalf("Error while checking table professors: %v", err)
	}
	if exists == 0 {
		_, err = db.Exec(`CREATE TABLE professors (
			id SERIAL PRIMARY KEY,
			jmbg VARCHAR(13) NOT NULL,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL
		);`)
		if err != nil {
			log.Fatalf("Error while creating the professors table: %v", err)
		}
	}

}
