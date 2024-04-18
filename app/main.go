package main

import (
	"encoding/json"
	"log"
	"os"

	"17.Working_with_date_and_time/models"
)

func main() {
	file, err := os.Open("employees.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var employees []models.Employee

	err = decoder.Decode(&employees)
	if err != nil {
		log.Fatal(err)
	}
	for i := range employees {
		if employees[i].ID == 3 {
			ChangeEmployee(&employees[i])
			break
		}
	}

	newEmployee := models.Employee {
		ID: 6,
		Name: "Nodirbek Numonov",
		Age: 96,
		Position: "Pensioner",
	}

	employees = append(employees, newEmployee)

	newFile, err := os.Create("employees_new.json")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	encoder := json.NewEncoder(newFile)

	err = encoder.Encode(employees)
	if err != nil {
		log.Fatal(err)
	}
}

func ChangeEmployee(employee *models.Employee) {
	employee.Age = 50
}
