package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

type InputError struct{ Errstr string }

func (i *InputError) Error() string {
	return i.Errstr
}

func rule110() []int {
	return []int{0, 1, 1, 1, 0, 1, 1, 0}
}

//flag can only consist of several distinct options
func flagOutCheck(flgout *string) {
	//check submitted flags
	if *flgout != "b" && *flgout != "s" && *flgout != "d" && *flgout != "a" {
		panic(InputError{"-out flag input malformed."})
	}
}

//flag can only consist of several distinct options
func flagLatCheck(flgout *string) {
	//check submitted flags
	if *flgout != "b" && *flgout != "s" && *flgout != "d" && *flgout != "a" {
		panic(InputError{"-out flag input malformed."})
	}
}

//an ic can only consist of <0>'s and <1>'s
func icCheck(ic string) {
	for i := 0; i < len(ic); i++ {
		cell := ic[i]
		if cell != '0' && cell != '1' {
			panic(InputError{"-out flag input malformed."})
		}
	}
}

//gens must be greater than 0 and a valid integer
func flagGenCheck(gens *int) {
	if *gens <= 0 {
		panic(InputError{"gen must be greater than 0 and a valid integer"})
	}
}

func InitLattice(ic string) *[]string {
	lattice := []string{}

	//apply initial condition
	lattice = append(lattice, ic)

	return &lattice
}

func main() {
	//graph output - binary/state/domain/all
	flagOut := flag.String("out", "b", "sets output parameters -(b)inary -(s)tate -(d)omain -(a)ll")

	//lattice - toroidal, zero edged, 1 edged, unbounded
	flagLat := flag.String("lat", "t", "sets lattice properties -(t)oroidal -(z)ero-edged -(o)ne-edged -(u)nbounded -(a)ll")

	//number of generations to compute
	flagGen := flag.Int("gens", 10, "sets number of generations to compute before stopping")
	flag.Parse()

	if len(flag.Args()) <= 0 {
		err := InputError{"Must include an initial condition string of 1's and 0's"}
		fmt.Println(err)
		panic(err)
	}

	//initial conditions are a string of 1's and 0's
	ic := flag.Args()[0]

	flagOutCheck(flagOut)
	flagLatCheck(flagLat)
	flagGenCheck(flagGen)
	icCheck(ic)

	//generate state ca and work off of that
	lattice := InitLattice(ic)
	//output

	fmt.Println("Program Arguments")
	fmt.Println("-out: " + *flagOut)
	fmt.Println("-lat: " + *flagLat)
	fmt.Println("-gens" + strconv.Itoa(*flagGen))
	fmt.Println("IC: " + flag.Args()[0])
	fmt.Println(time.Now())
	fmt.Println(lattice)

}
