package main

import "fmt"

type Studentt struct{
	Name 		string
	Scholarship int
	Course 		int
}

func GetScholarship(students []Studentt) {
	allScholarship := 0
	for _, student := range students {
		if student.Course == 2 {
	  		allScholarship += student.Scholarship
	 	}
	}
fmt.Println(allScholarship)
}
   

func main()  {
	student := []Studentt{
		{
			Name: "Abdulloh",
			Scholarship: 500,
			Course: 2,
		},
		{
			Name: "Vali",
			Scholarship: 400,
			Course: 4,
		},
		{
			Name: "Rustam",
			Scholarship: 200,
			Course: 2,
		},
		{
			Name: "Aziz",
			Scholarship: 350,
			Course: 1,
	    },
		{
			Name: "Nurbek",
			Scholarship: 240,
			Course: 2,
	    },
		{
			Name: "Abbos",
			Scholarship: 500,
			Course: 2,
		},
		{
			Name: "Sherali",
			Scholarship: 400,
			Course: 4,
		},
		{
			Name: "Ruslon",
			Scholarship: 200,
			Course: 2,
	    },
		{
			Name: "Aziza",
			Scholarship: 350,
			Course: 1,
	    },
		{
			Name: "Nurali",
			Scholarship: 240,
			Course: 2,
	    },
	}
GetScholarship(student)
}