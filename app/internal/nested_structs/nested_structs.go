package nested_structs

import "fmt"

var pl = fmt.Println

type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

// attach method to struct
func (b business) info() {
	fmt.Printf("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.lName)
}

func NestedStructsExample() {
	cS := contact{"Mark", "Ledger", "07777777770"}
	bS := business{"BJSS", "Place Street", cS}

	bS.info()
}
