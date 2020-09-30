package invoiceitem

// Item contains information about invoice items
type Item struct {
	id      uint
	product string
	value   float64
}

// New is a constructor to create invoice items
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Value is a functions to return a value
func (i Item) Value() float64 {
	return i.value
}
