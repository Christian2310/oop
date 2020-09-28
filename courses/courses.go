package courses

import "fmt"

// Course is an structure to create new courses.
type course struct {
	name    string
	price   float64
	free    bool
	userID  []uint
	lessons map[uint]string
}

func (c *course) SetName(name string)               { c.name = name }
func (c *course) Name() string                      { return c.name }
func (c *course) SetPrice(price float64)            { c.price = price }
func (c *course) Price() float64                    { return c.price }
func (c *course) SetUserID(userid []uint)           { c.userID = userid }
func (c *course) UserID() []uint                    { return c.userID }
func (c *course) SetLessons(lesson map[uint]string) { c.lessons = lesson }
func (c *course) Lessons() map[uint]string          { return c.lessons }

// New is a constructor function
func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}

	return &course{
		name:  name,
		price: price,
		free:  isFree,
	}
}

// PrintLessons is an example of Methods
// In GO the Methods are created outsie of the strcture.
func (c course) PrintLessons() {
	text := "The lessons for "
	courseName := c.name
	for _, lesson := range c.lessons {
		text += courseName + " are: " + lesson + ", "
	}
	fmt.Println(text[:len(text)-2])
}
