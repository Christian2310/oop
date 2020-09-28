package courses

import "fmt"

// Course is an structure to create new courses.
type course struct {
	Name    string
	Price   float64
	Free    bool
	UserID  []uint
	Lessons map[uint]string
}

// New is a constructor function
func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}

	return &course{
		Name:  name,
		Price: price,
		Free:  isFree,
	}
}

// PrintLessons is an example of Methods
// In GO the Methods are created outsie of the strcture.
func (c course) PrintLessons() {
	text := "The lessons for "
	courseName := c.Name
	for _, lesson := range c.Lessons {
		text += courseName + " are: " + lesson + ", "
	}
	fmt.Println(text[:len(text)-2])
}

// ChangePrice is a Method with a pointer to change the price os the object
func (c *course) ChangePrice(price float64) {
	c.Price = price
}
