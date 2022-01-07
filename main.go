package main

import (
	"flag"
	"fmt"
	"time"
)

type InputError struct{ Errstr string }

func (i *InputError) Error() string {
	return i.Errstr
}

func rule110() []int {
	return []int{0, 1, 1, 1, 0, 1, 1, 0}
}

func flgoutCheck(flgout *string) {
	//check submitted flags
	if *flgout != "b" && *flgout != "s" && *flgout != "d" && *flgout != "a" {
		panic(InputError{"-out flag input malformed."})
	}
}

func flglatCheck(flgout *string) {
	//check submitted flags
	if *flgout != "b" && *flgout != "s" && *flgout != "d" && *flgout != "a" {
		panic(InputError{"-out flag input malformed."})
	}
}

func icCheck(ic string) {
	for i := 0; i < len(ic); i++ {
		cell := ic[i]
		if cell != '0' && cell != '1' {
			panic(InputError{"-out flag input malformed."})
		}
	}
}

func main() {
	//graph output - binary/state/domain/all
	flgout := flag.String("out", "b", "sets output parameters -(b)inary -(s)tate -(d)omain -(a)ll")

	//lattice - toroidal, zero edged, 1 edged, unbounded
	flglat := flag.String("lat", "t", "sets lattice properties -(t)oroidal -(z)ero-edged -(o)ne-edged -(u)nbounded -(a)ll")
	flag.Parse()

	//initial conditions are a string of 1's and 0's
	ic := flag.Args()[0]

	flgoutCheck(flgout)
	flglatCheck(flglat)
	icCheck(ic)

	//generate state ca and work off of that

	fmt.Println("Program Arguments")
	fmt.Println("-out: " + *flgout)
	fmt.Println("-lat: " + *flglat)
	fmt.Println("IC: " + flag.Args()[0])
	fmt.Println(time.Now())
}
