package courses

import "fmt"

// Course is an structure to create new courses.
type Course struct {
	Name    string
	Price   float64
	Free    bool
	UserID  []uint
	Lessons map[uint]string
}

// PrintLessons is an example of Methods
// In GO the Methods are created outsie of the strcture.
func (c Course) PrintLessons() {
	text := "The lessons for "
	courseName := c.Name
	for _, lesson := range c.Lessons {
		text += courseName + " are: " + lesson + ", "
	}
	fmt.Println(text[:len(text)-2])
}

// ChangePrice is a Method with a pointer to change the price os the object
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}
