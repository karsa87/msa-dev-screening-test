package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Hobby string `json:"hobby"`
}

func main() {
	name := []string{"brian", "habib", "malik"}
	age := []int{25, 25, 24}
	hobby := []string{"hiking", "touring", "traveling"}

	students := make([]Student, 0)

	n := 0
	for n < 3 {
		student := Student{}
		student.Name = name[n]
		student.Age = age[n]
		student.Hobby = hobby[n]
		students = append(students, student)
		n++
	}

	studentsJson, err := json.Marshal(students)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Println(string(studentsJson))
}
