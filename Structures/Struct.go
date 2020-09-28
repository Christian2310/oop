// package main

// import "fmt"

// // Course is an structure to create new courses.
// type Course struct {
// 	Name    string
// 	Price   float64
// 	Free    bool
// 	UserID  []uint
// 	Lessons map[uint]string
// }

// // PrintLessons is an example of Methods
// // In GO the Methods are created outsie of the strcture.
// func (c Course) PrintLessons() {
// 	text := "The lessons for "
// 	courseName := c.Name
// 	for _, lesson := range c.Lessons {
// 		text += courseName + " are: " + lesson + ", "
// 	}
// 	fmt.Println(text[:len(text)-2])
// }

// func main() {
// 	// Go is one of the different ways to create an instance of the "Course Struct".
// 	GO := Course{
// 		Name:   "Go from scratch",
// 		Price:  29.99,
// 		Free:   false,
// 		UserID: []uint{12, 56, 90},
// 		Lessons: map[uint]string{
// 			1: "Variables",
// 			2: "Data types",
// 			3: "Slices",
// 			4: "Functions",
// 			5: "Methods",
// 		},
// 	}
// 	fmt.Println(GO.Name, GO.Lessons)

// 	// JS is an object to show thta if you will use all field in the struc you can omit the name of the fields.
// 	JS := Course{
// 		"Javascript",
// 		29.99,
// 		false,
// 		[]uint{12, 56, 90},
// 		map[uint]string{
// 			1: "Variables",
// 			2: "Data types",
// 			3: "Slices",
// 			4: "Functions",
// 			5: "Methods",
// 		},
// 	}
// 	fmt.Println(JS.Name, JS.Lessons)

// 	// CSS is an explample of "Course Struct" instance with just some fields.
// 	CSS := Course{Name: "CSS from scratch", Price: 29.99}
// 	// to this instance GO will assing to the not used fields the 0 values.
// 	fmt.Printf("%v\n", CSS)

// 	// SASS example way to create
// 	SASS := Course{}
// 	SASS.Name = "SASS"
// 	SASS.Price = 14.99
// 	fmt.Printf("%v\n", SASS)

// 	GO.PrintLessons()
// 	JS.PrintLessons()

// 	// Is not necesari to pass the "&" to complete the pointer, GO reads the "*" in the functions and knows that is a pointer
// 	(GO).ChangePrice(67.12)
// 	fmt.Println(GO.Price)

// }

// // ChangePrice is a Method with a pointer to change the price os the object
// func (c *Course) ChangePrice(price float64) {
// 	c.Price = price
// }
