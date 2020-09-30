package main

import (
	"fmt"

	"github.com/Christian2310/oop/pkg/pkg/composition/customer"
	"github.com/Christian2310/oop/pkg/pkg/composition/invoice"
	"github.com/Christian2310/oop/pkg/pkg/composition/invoiceitem"
)

func main() {
	receipt := invoice.New(
		"USA",
		"Rancho Cucamonga",
		customer.New("Christian Colorado", "9445 Charles Smith Ave", "9096843224"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "BF", 250.25),
			invoiceitem.New(1, "NB", 130.5),
			invoiceitem.New(1, "EFT", 140.15),
		},
	)
	receipt.SetTotal()
	fmt.Printf("%v", receipt)

}
