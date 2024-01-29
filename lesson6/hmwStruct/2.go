package main

import "fmt"

type Student struct {
 Name        string
 Scholarship int
 Course      int
}

func (s *Student) ShortName(students []Student) []Student {
	var filteredStudents []Student
	for _, student := range students {
		if len(student.Name) < 5 {
			filteredStudents = append(filteredStudents, student)
  		}
 	}
return filteredStudents
}

func main() {
	students := []Student{
	{
	Name:        "Abdulloh",
	Scholarship: 500,
	Course:      2,
	},
	{
	Name:        "Vali",
	Scholarship: 400,
	Course:      4,
	},
	{
	Name:        "Rustam",
	Scholarship: 200,
	Course:      2,
	},
	{
	Name:        "Aziz",
	Scholarship: 350,
	Course:      1,
	},
	{
	Name:        "Nurbek",
	Scholarship: 240,
	Course:      2,
	},
	{
	Name:        "Abbos",
	Scholarship: 500,
	Course:      2,
	},
	{
	Name:        "Sherali",
	Scholarship: 400,
	Course:      4,
	},
	{
	Name:        "Ruslon",
	Scholarship: 200,
	Course:      2,
	},
	{
	Name:        "Aziza",
	Scholarship: 350,
	Course:      1,
	},
	{
	Name:        "Nurali",
	Scholarship: 240,
	Course:      2,
	},
	}

 	var studentObj Student
 	filteredStudents := studentObj.ShortName(students)

 	fmt.Println("Filtered:")
 	for _, student := range filteredStudents {
  		fmt.Println(student)
 	}
}