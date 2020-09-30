package invoice

import (
	"github.com/Christian2310/oop/pkg/pkg/composition/customer"
	"github.com/Christian2310/oop/pkg/pkg/composition/invoiceitem"
)

// Invoice is the struct of a receipt/invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

// New is a constructor function to create New Invoices/Receipts
func New(country string, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is a seter to sum the invoice total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
