package structs

import "fmt"

var pl = fmt.Println

type customer struct {
	name    string
	address string
	bal     float64
}

func (c customer) PrintNameAndAddress() {
	pl(fmt.Sprintf("%s currently lives at %s \n", c.name, c.address))
}

func StructsExample() {
	var tS customer
	tS.name = "Jude"
	tS.address = "High Hunting Road"
	tS.bal = 23.45

	getCustInfo(tS)
	newCustomerAdd(&tS, "Heeley Road")
	pl("Address: ", tS.address)

	sS := customer{"Sally Smith", "123 Main Road", 0.0}
	pl("Name: ", sS.name)

	sS.PrintNameAndAddress()
}

func newCustomerAdd(c *customer, address string) {
	c.address = address

}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f \n", c.name, c.bal)
}
