package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
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
	http.HandleFunc("/student", handleStudentRequest)
	http.HandleFunc("/professor", handleProfessorRequest)
	http.ListenAndServe(":8050", nil)
}

func handleStudentRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var student student
		json.NewDecoder(r.Body).Decode(&student)

		db, err := sql.Open("postgres", "user=postgres password=password dbname=uns port=5432 sslmode=disable")
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
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func handleProfessorRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var professor professor
		json.NewDecoder(r.Body).Decode(&professor)

		db, err := sql.Open("postgres", "user=postgres password=password dbname=uns port=5432 sslmode=disable")
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
