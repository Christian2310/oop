package main

import (
	"fmt"

	"github.com/Christian2310/oop/courses"
)

func main() {
	// 	Go is one of the different ways to create an instance of the "Course Struct".
	GO := courses.Course{
		Name:   "Go from scratch",
		Price:  29.99,
		Free:   false,
		UserID: []uint{12, 56, 90},
		Lessons: map[uint]string{
			1: "Variables",
			2: "Data types",
			3: "Slices",
			4: "Functions",
			5: "Methods",
		},
	}
	fmt.Println(GO)
}
