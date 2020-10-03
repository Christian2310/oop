// package main

// import "fmt"

// type course struct {
// 	name string
// }

// func (c course) Print() {
// 	fmt.Printf("%+v\n", c)
// }

// // Alias declarations
// type myAlias = course

// type newBool bool

// func (b newBool) String() string {
// 	if b {
// 		return "VERDADERO"
// 	}
// 	return "FALSO"
// }

// // Type definitions
// //////////////////////////////////////////////////////
// // 													//
// //	When you use a type definitions you can access  //
// //	the fields of the original structure but you    //
// //	can't access to the methods for the original    //
// //	struct										    //
// //												    //
// //////////////////////////////////////////////////////
// type newCourse course

// func main() {
// 	c := newCourse{name: "Go"}
// 	fmt.Printf("El tipo de dato es: %T\n", c)

// 	var b newBool = true

// 	fmt.Println(b.String())

// }
