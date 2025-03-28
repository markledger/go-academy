package defined_types

import (
	"errors"
	"fmt"
)

var pl = fmt.Println

type Tsp float64
type Tbsp float64
type ML float64

func (tsp Tsp) TspToML() ML {
	return ML(tsp * 4.92)
}

func TbspToML(tbsp Tbsp) ML {
	return ML(tbsp * 14.79)
}

type Date struct {
	day   int
	month int
	year  int
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}

	d.day = day
	return nil
}

func DefinedTypeExample() {
	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 Tsp = %.2f ML\n", ml1)
	ml2 := ML(Tbsp(3) * 14.79)
	fmt.Printf("3 Tbsp = %.2f ML\n", ml2)
	twoTsp := Tsp(2)
	pl("2tsp > 4tsp", Tsp(2) > Tsp(4))
	pl("2tsp in ml = ", twoTsp.TspToML())

	var d = Date{}
	err := d.SetDay(33)
	if err != nil {
		pl(err)
	}
}
